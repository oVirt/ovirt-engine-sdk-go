/*
Copyright (c) 2016 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package org.ovirt.sdk.go;

import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toCollection;
import static java.util.stream.Collectors.toList;

import java.io.File;
import java.io.IOException;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Deque;
import java.util.List;
import java.util.Optional;
import java.util.function.Predicate;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.Concept;
import org.ovirt.api.metamodel.concepts.ListType;
import org.ovirt.api.metamodel.concepts.Locator;
import org.ovirt.api.metamodel.concepts.Method;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.NameParser;
import org.ovirt.api.metamodel.concepts.Parameter;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.Service;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.SchemaNames;

/**
 * This class is responsible for generating the classes that represent the services of the model.
 */
public class ServicesGenerator implements GoGenerator {
    // Well known method names:
    private static final Name ADD = NameParser.parseUsingCase("Add");
    private static final Name GET = NameParser.parseUsingCase("Get");
    private static final Name LIST = NameParser.parseUsingCase("List");
    private static final Name REMOVE = NameParser.parseUsingCase("Remove");
    private static final Name UPDATE = NameParser.parseUsingCase("Update");

    // The directory were the output will be generated:
    protected File out;

    // Reference to the objects used to generate the code:
    @Inject private GoNames goNames;
    @Inject private SchemaNames schemaNames;

    // The buffer used to generate the code:
    private GoBuffer buffer;

    /**
     * Set the directory were the output will be generated.
     */
    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Prepare the buffer:
        buffer = new GoBuffer();
        buffer.setPackageName(goNames.getServicesPackageName());

        // Generate the code:
        generateServices(model);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing services module", exception);
        }
    }

    private void generateServices(Model model) {
        // The declarations of the services need to appear in inheritance order, otherwise some symbols won't be
        // defined and that will produce errors. To order them correctly we need first to sort them by name, and
        // then sort again so that bases are before extensions.
        Deque<Service> pending = model.services()
            .sorted()
            .collect(toCollection(ArrayDeque::new));
        Deque<Service> sorted = new ArrayDeque<>(pending.size());
        while (!pending.isEmpty()) {
            Service current = pending.removeFirst();
            Service base = current.getBase();
            if (base == null || sorted.contains(base)) {
                sorted.addLast(current);
            }
            else {
                pending.addLast(current);
            }
        }
        sorted.forEach(this::generateService);
    }

    private void generateService(Service service) {
        // Begin class:
        GoClassName serviceName = goNames.getServiceName(service);
        Service base = service.getBase();
        GoClassName baseName = base != null? goNames.getServiceName(base): goNames.getBaseServiceName();
        
        // Generate struct definition
        generateDoc(service);
        buffer.addLine("type %1$s struct {", serviceName.getClassName());

        // Generate struct members definition
        buffer.startBlock();
        //      with Service struct mixin
        buffer.addLine(baseName.getClassName());
        buffer.addLine();
        //      members
        service.locators().sorted().forEach(this::generateLocatorMember);
        buffer.endBlock();
        // Generate struct ending
        buffer.addLine("}");
        buffer.addLine();

        // Generate the constructor by New:
        buffer.addLine(
            "func New%1$s(connection *Connection, path string) *%2$s {",
            serviceName.getClassName(), serviceName.getClassName());
        buffer.startBlock();
        //      inititalize struct
        buffer.addLine("var result %1$s", serviceName.getClassName());
        buffer.addLine("result.Connection = connection");
        buffer.addLine("result.Path = path");
        buffer.addLine("return &result");
        buffer.endBlock();

        // Generate constructor ending
        buffer.addLine("}");
        buffer.addLine();

        // Generate the methods
        List<Method>methods = service.methods().sorted().collect(toCollection(ArrayList::new));
        for (Method method : methods) {
            this.generateMethod(method, service);
        }
        
        // Generate locator methods
        List<Locator> locators = service.locators().sorted().collect(toCollection(ArrayList::new));
        for (Locator locator : locators) {
            this.generateLocatorMethod(locator, service);
        }

        // Generate locators by path
        generatePathLocator(service);

        // Generate other methods that don't correspond to model methods or locators:
        generateStr(service);

        buffer.addLine();
    }

    private void generateMethod(Method method, Service service) {
        Name name = method.getName();
        if (ADD.equals(name)) {
            generateAddHttpPost(method, service);
        }
        else if (GET.equals(name) || LIST.equals(name)) {
            generateHttpGet(method, service);
        }
        else if (REMOVE.equals(name)) {
            generateHttpDelete(method, service);
        }
        else if (UPDATE.equals(name)) {
            generateHttpPut(method, service);
        }
        else {
            generateActionHttpPost(method, service);
        }
    }

    private void generateAddHttpPost(Method method, Service service) {
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);

        // Get the parameters:
        Parameter primaryParameter = getFirstParameter(method);
        List<Parameter> secondaryParameters = getSecondaryParameters(method);

        // Begin method:
        Name methodName = method.getName();
        //      get primary parameter name
        Name primaryParameterName = primaryParameter.getName();
        String primaryArg = goNames.getParameterStyleName(primaryParameterName);
        //      get primary parameter type name
        GoTypeReference goTypeReference = goNames.getTypeReference(primaryParameter.getType());
        buffer.addImports(goTypeReference.getImports());
        String primaryArgTypeName = goTypeReference.getText();

        //      Generate function doc
        generateActionDoc(method, (Parameter p) -> p.isIn() && p.isOut());
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s (",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        //      Generate func-codes definition
        buffer.startBlock();
        buffer.addLine("%1$s %2$s,", primaryArg, primaryArgTypeName);
        secondaryParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) {");
        //      Generate function ending
        buffer.endBlock();
 
        // Start body:
        buffer.startBlock();
        //      Generate the wait default is True
        buffer.addLine("if wait == nil {");
        buffer.startBlock();
        buffer.addLine("wait = true");
        buffer.endBlock();
        buffer.addLine("}");
        //      Generate the code to build the URL query:
        buffer.addLine("// Build the URL:");
        buffer.addLine("if query == nil {");
        buffer.startBlock();
        buffer.addLine("query = make(map[string]string)");
        buffer.endBlock();
        buffer.addLine("}");
        secondaryParameters.forEach(this::generateUrlParameter);
        buffer.addLine();
        //      Generate the code to send the request
        buffer.addLine("// Send the request");
        buffer.addLine("return op.internalAdd(%1$s, headers, query, wait)", primaryArg);

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateActionHttpPost(Method method, Service service) {
        // Get the input parameters:
        List<Parameter> inParameters = method.parameters()
            .filter(Parameter::isIn)
            .sorted()
            .collect(toList());
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);
        // Begin method:
        Name methodName = method.getName();
        //      Generate function doc
        generateActionDoc(method, Parameter::isIn);
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s (",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) {");
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();
        //      Generate the wait default is True
        buffer.addLine("if wait == nil {");
        buffer.startBlock();
        buffer.addLine("wait = true");
        buffer.endBlock();
        buffer.addLine("}");
        //      Generate the code to populate the action:
        buffer.addLine("// Populate the action:");
        buffer.addLine("action = Action{");
        buffer.startBlock();
        inParameters.forEach(this::generateSetActionAttribute);
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();

        // Generate the code to send the request and wait for the response:
        Parameter parameter = method.parameters()
            .filter(Parameter::isOut)
            .findFirst()
            .orElse(null);
        String member = parameter == null ? null : goNames.getParameterStyleName(parameter.getName());

        buffer.addLine("// Send the request and wait for the response:");
        if (member == null) {
            buffer.addLine("return internalAction(action, \"%1$s\", nil, headers, query, wait)",  getPath(methodName));
        }
        else {
            buffer.addLine(
                "return internalAction(action, \"%1$s\", \"%2$s\", headers, query, wait)",
                getPath(methodName),
                member
            );
        }

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateHttpGet(Method method, Service service) {
        // Get the input parameters:
        List<Parameter> inParameters = method.parameters()
            .filter(Parameter::isIn)
            .sorted()
            .collect(toList());
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);
        // Begin method:
        Name methodName = method.getName();
        //      Generate function doc
        generateActionDoc(method, Parameter::isIn);
        //      Generate function definition
        //          Get returned type
        String returnedType = "*" +  serviceName.getClassName().substring(
            0, serviceName.getClassName().indexOf("Service"));
        //          Generate function body
        buffer.addLine(
            "func (op *%1$s) %2$s (",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) %1$s {", returnedType);
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();
        //      Generate the wait default is True
        buffer.addLine("if wait == nil {");
        buffer.startBlock();
        buffer.addLine("wait = true");
        buffer.endBlock();
        buffer.addLine("}");
        //      Generate the code to build the URL query:
        buffer.addLine("// Build the URL:");
        buffer.addLine("if query == nil {");
        buffer.startBlock();
        buffer.addLine("query = make(map[string]string)");
        buffer.endBlock();
        buffer.addLine("}");
        inParameters.forEach(this::generateUrlParameter);
        buffer.addLine();

        // Generate the code to send the request and wait for the response:
        buffer.addLine("// Send the request and wait for the response:");
        buffer.addLine("return op.internalGet(headers, query, wait)");

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateHttpPut(Method method, Service service) {
        // Classify the parameters:
        Parameter primaryParameter = getFirstParameter(method);
        List<Parameter> secondaryParameters = getSecondaryParameters(method);
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);

        // Begin method:
        Name methodName = method.getName();
        //      get primary parameter name
        Name primaryParameterName = primaryParameter.getName();
        String primaryArg = goNames.getParameterStyleName(primaryParameterName);
        //      get primary parameter type name
        GoTypeReference goTypeReference = goNames.getTypeReference(primaryParameter.getType());
        buffer.addImports(goTypeReference.getImports());
        String primaryArgTypeName = goTypeReference.getText();
        //      Generate function doc
        generateActionDoc(method, (Parameter p) -> p.isIn() && p.isOut());
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s (",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        //      Generate function parameters definition
        buffer.startBlock();
        buffer.addLine("%1$s %2$s,", primaryArg, primaryArgTypeName);
        secondaryParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) {");
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();
        //      Generate the wait default is True
        buffer.addLine("if wait == nil {");
        buffer.startBlock();
        buffer.addLine("wait = true");
        buffer.endBlock();
        buffer.addLine("}");
        //      Generate the code to build the URL query:
        buffer.addLine("// Build the URL:");
        buffer.addLine("if query == nil {");
        buffer.startBlock();
        buffer.addLine("query = make(map[string]string)");
        buffer.endBlock();
        buffer.addLine("}");
        secondaryParameters.forEach(this::generateUrlParameter);
        buffer.addLine();
        //      Generate the code to send the request
        buffer.addLine("// Send the request");
        buffer.addLine("return op.internalUpdate(%1$s, headers, query, wait)", primaryArg);

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateHttpDelete(Method method, Service service) {
        // Get the parameters:
        List<Parameter> inParameters = method.parameters()
            .filter(Parameter::isIn)
            .collect(toList());
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);
        // Begin method:
        Name methodName = method.getName();
        //      Generate function doc
        generateActionDoc(method, Parameter::isIn);
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s (",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) {");
        //      Generate function ending
        buffer.endBlock();

        // Begin body:
        buffer.startBlock();
        //      Generate the wait default is True
        buffer.addLine("if wait == nil {");
        buffer.startBlock();
        buffer.addLine("wait = true");
        buffer.endBlock();
        buffer.addLine("}");
        //      Generate the code to build the URL query:
        buffer.addLine("// Build the URL:");
        buffer.addLine("if query == nil {");
        buffer.startBlock();
        buffer.addLine("query = make(map[string]string)");
        buffer.endBlock();
        buffer.addLine("}");
        inParameters.forEach(this::generateUrlParameter);
        buffer.addLine();

        // Generate the code to send the request and wait for the response:
        buffer.addLine("// Send the request and wait for the response:");
        buffer.addLine("op.internalRemove(headers, query, wait)");

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateFormalParameter(Parameter parameter) {
        // Get parameter name
        Name parameterName = parameter.getName();
        GoTypeReference goTypeReference = goNames.getTypeReference(parameter.getType());
        buffer.addImports(goTypeReference.getImports());

        String arg = goNames.getParameterStyleName(parameterName);
        // Get parameter type name
        buffer.addLine(
            "%1$s %2$s,", arg, goTypeReference.getText());
    }

    private void generateUrlParameter(Parameter parameter) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        String arg = goNames.getParameterStyleName(name);
        String tag = schemaNames.getSchemaTagName(name);
        buffer.addLine("if %1$s != nil {", arg);
        buffer.startBlock();
        buffer.addLine("query[\"%1$s\"] = %2$s", tag, arg);
        buffer.endBlock();
        buffer.addLine("}");
    }

    private void generateStr(Service service) {
        GoClassName serviceName = goNames.getServiceName(service);
        buffer.addLine("func (op *%1$s) String() string {", serviceName.getClassName());
        buffer.startBlock();
        buffer.addImport("fmt");
        buffer.addLine("return fmt.Sprintf(\"%1$s:%%s\", op.Path)", serviceName.getClassName());
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateLocatorMember(Locator locator) {
        String memberName = goNames.getMemberStyleName(locator.getName());
        buffer.addLine("%1$sService  *%1$sService", memberName, memberName);
    }

    private void generateLocatorMethod(Locator locator, Service service) {
        Parameter parameter = locator.getParameters().stream().findFirst().orElse(null);
        if (parameter != null) {
            generateLocatorWithParameters(locator, service);
        }
        else {
            generateLocatorWithoutParameters(locator, service);
        }
    }

    private void generateLocatorWithParameters(Locator locator, Service service) {
        Parameter parameter = locator.parameters().findFirst().get();
        String methodName = goNames.getMemberStyleName(locator.getName());
        String argName = goNames.getParameterStyleName(parameter.getName());
        GoTypeReference parameterTypeReference = goNames.getTypeReference(parameter.getType());
        buffer.addImports(parameterTypeReference.getImports());
        GoClassName locatorServiceName = goNames.getServiceName(locator.getService());
        generateDoc(locator);

        // Get receiver class
        GoClassName receiverClassName = goNames.getServiceName(service);

        buffer.addLine(
            "func (op *%1$s) %2$sService(%3$s %4$s) *%5$s {",
            receiverClassName.getClassName(),
            methodName, argName, parameterTypeReference.getText(),
            locatorServiceName.getClassName());

        buffer.startBlock();
        buffer.addImport("fmt");
        buffer.addLine(
            "return New%1$s(op.Connection, fmt.Sprintf(\"%%s/%%s\", op.Path, %2$s))",
            locatorServiceName.getClassName(),
            argName);
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateLocatorWithoutParameters(Locator locator, Service service) {
        String methodName = goNames.getMethodStyleName(locator.getName());
        String urlSegment = getPath(locator.getName());
        GoClassName locatorServiceName = goNames.getServiceName(locator.getService());
        generateDoc(locator);

        // Get receiver class
        GoClassName receiverClassName = goNames.getServiceName(service);
        // Generate *Service function
        buffer.addLine("func (op *%1$s) %2$sService() *%3$s {",
            receiverClassName.getClassName(), methodName, locatorServiceName.getClassName());

        buffer.startBlock();
        buffer.addImport("fmt");
        buffer.addLine(
            "return New%1$s(op.Connection, fmt.Sprintf(\"%%s/%2$s\", op.Path))",
            locatorServiceName.getClassName(),
            urlSegment);
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generatePathLocator(Service service) {
        GoClassName serviceName = goNames.getServiceName(service);
        // Generate comment
        buffer.startComment();
        buffer.addLine("// Service locator method, returns individual service on which the URI is dispatched.");
        buffer.endComment();

        // Begin method:
        buffer.addLine("func (op *%1$s) Service(path string) (IService, error) {", serviceName.getClassName());
        buffer.startBlock();
        buffer.addLine("if path == nil {");
        buffer.startBlock();
        buffer.addLine("return *op, nil");
        buffer.endBlock();
        buffer.addLine("}");

        // Generate the code that checks if the path corresponds to any of the locators without parameters:
        service.locators().filter(x -> x.getParameters().isEmpty()).sorted().forEach(locator -> {
            Name name = locator.getName();
            String segment = getPath(name);
            buffer.addLine("if path == \"%1$s\" {", segment);
            buffer.startBlock();
            buffer.addLine(  "return *(op.%1$sService()), nil", goNames.getMethodStyleName(name));
            buffer.endBlock();
            buffer.addLine("}");
            buffer.addLine("if strings.HasPrefix(\"%1$s/\") {", segment);
            buffer.addImport("strings");
            buffer.startBlock();
            buffer.addLine(
                "return op.%1$sService().Service(path[%2$d:]), nil",
                goNames.getMemberStyleName(name),
                segment.length() + 1
            );
            buffer.endBlock();
            buffer.addLine("}");
        });

        // If the path doesn't correspond to a locator without parameters, then it will correspond to the locator
        // with parameters, otherwise it is an error:
        Optional<Locator> optional = service.locators().filter(x -> !x.getParameters().isEmpty()).findAny();
        if (optional.isPresent()) {
            Locator locator = optional.get();
            Name name = locator.getName();
            buffer.addImport("strings");
            buffer.addLine("index = strings.Index(path, \"/\")");
            buffer.addLine("if index == -1:");
            buffer.startBlock();
            buffer.addLine("return *(op.%1$sService(path)), nil", goNames.getMemberStyleName(name));
            buffer.endBlock();
            buffer.addLine(
                "return op.%1$sService(path[:index]).Service(path[index + 1:]), nil",
                goNames.getMemberStyleName(name)
            );
        }
        else {
            buffer.addImport("errors");
            buffer.addLine("return nil, errors.New(fmt.Sprintf(\"The path <%%s> doesn\\'t correspond to any service\", path))");
        }

        // End method:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private String getPath(Name name) {
        return name.words().map(String::toLowerCase).collect(joining());
    }

    private void generateActionDoc(Method method, Predicate<Parameter> predicate) {
        buffer.startComment();
        if (method.getDoc() != null) {
            generateDocText(method);
        }
        if (method.parameters().filter(predicate).findFirst().orElse(null) != null) {
            List<String> lines = method.parameters()
                .filter(predicate)
                .filter(p -> p.getDoc() != null)
                .map(p -> String.format("`%s`:: %s", goNames.getMemberStyleName(p.getName()), p.getDoc()))
                .collect(toList());

            if (!lines.isEmpty()) {
                buffer.addLine("// This method supports the following parameters:");
                lines.forEach(this::generateDocText);
                buffer.addLine("// `headers`:: Additional HTTP headers.");
                buffer.addLine("// `query`:: Additional URL query parameters.");
                buffer.addLine("// `wait`:: If `True` wait for the response.");
            }
        }
        buffer.endComment();
    }

    protected void generateDoc(Concept concept) {
        buffer.startComment();
        generateDocText(concept);
        buffer.endComment();
    }

    private void generateDocText(Concept concept) {
        generateDocText(concept.getDoc());
    }

    private void generateDocText(String doc) {
        List<String> lines = new ArrayList<>();
        if (doc != null) {
            Collections.addAll(lines, doc.split("\n"));
        }
        if (!lines.isEmpty()) {
            lines.stream().filter(l -> !l.isEmpty()).forEach(buffer::addCommentLine);
        }
    }

    private Parameter getFirstParameter(Method method) {
        return method.parameters()
            .filter(x -> x.isIn() && x.isOut())
            .findFirst()
            .orElse(null);
    }

    private List<Parameter> getSecondaryParameters(Method method) {
        return method.parameters()
            .filter(x -> x.isIn() && !x.isOut())
            .sorted()
            .collect(toList());
    }

    private void generateSetActionAttribute(Parameter parameter) {
        String memberName = goNames.getMemberStyleName(parameter.getName());
        String parameterName = goNames.getParameterStyleName(parameter.getName());

        buffer.addLine("%1$s: %2$s,", memberName, parameterName);
    }

}
