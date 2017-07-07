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
        // buffer.addLine(baseName.getClassName());
        buffer.addLine("BaseService");
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

        // Get the primary param, Add function: the first In is just the Out param
        Parameter primaryParameter = getFirstParameter(method);
        // Get other in params
        List<Parameter> secondaryParameters = getSecondaryParameters(method);

        // Begin method:
        Name methodName = method.getName();
        //      get primary parameter name
        String primaryArg = goNames.getParameterStyleName(primaryParameter.getName());
        //      get primary parameter type name
        GoTypeReference primaryParameterGoTypeReference = goNames.getTypeReference(primaryParameter.getType());
        buffer.addImports(primaryParameterGoTypeReference.getImports());

        //      Generate function doc
        generateActionDoc(method, (Parameter p) -> p.isIn() && p.isOut());
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s(",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        //      Generate func-codes definition
        buffer.startBlock();
        buffer.addLine("%1$s %2$s,", primaryArg, primaryParameterGoTypeReference.getText());
        secondaryParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) (");
        //      Generate the output parameters
        buffer.startBlock();
        this.generateOutputParameter(primaryParameter);
        buffer.addLine("error) {");
        buffer.endBlock();
        //      Generate function ending
        buffer.endBlock();
 
        // Start body:
        buffer.startBlock();

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
        buffer.addLine("// Send the request and get the response");
        buffer.addLine("ovResp, err := op.internalAdd(%1$s, headers, query, wait)", primaryArg);
        //      Parse the result
        this.generateOvResponseParsing(primaryParameter);
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
        //      Get function output parameters
        Parameter outParam = getOutParameter(method);
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s(",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) (");
        //      Generate the output parameters
        buffer.startBlock();
        if (outParam != null) {
            this.generateOutputParameter(outParam);
        }
        buffer.addLine("error) {");
        buffer.endBlock();
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();

        //      Generate the code to populate the action:
        buffer.addLine("// Populate the action:");
        buffer.addLine("action := &Action{");
        buffer.startBlock();
        inParameters.forEach(this::generateSetActionAttribute);
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();

        buffer.addLine("// Send the request and wait for the response:");
        if (outParam == null) {
            buffer.addLine(
                "_, err := op.internalAction(action, \"%1$s\", headers, query, wait)",
                getPath(methodName)
            );
            buffer.addLine("return err");
        } else {
            buffer.addLine(
                "ovResp, err := op.internalAction(action, \"%1$s\", headers, query, wait)",
                getPath(methodName)
            );
            // Generate the ovResponse parsing
            this.generateOvResponseParsing(outParam);
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
        //      Get function output parameters
        Parameter outParameter = getOutParameter(method);
        GoTypeReference outParamGoTypeReference = goNames.getTypeReference(outParameter.getType());
        //      Generate function body
        buffer.addLine(
            "func (op *%1$s) %2$s(",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) ( ");
        //      Generate the output parameters
        buffer.startBlock();
        this.generateOutputParameter(outParameter);
        buffer.addLine("error) {");
        buffer.endBlock();
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();

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
        buffer.addLine("ovResp, err := op.internalGet(headers, query, wait)");

        // Generate ovResponse parsing
        this.generateOvResponseParsing(outParameter);

        // End body:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateHttpPut(Method method, Service service) {
        // Classify the parameters: it's also the Output param
        Parameter primaryParameter = getFirstParameter(method);
        List<Parameter> secondaryParameters = getSecondaryParameters(method);
        // Get the service class name
        GoClassName serviceName = goNames.getServiceName(service);

        // Begin method:
        Name methodName = method.getName();
        //      get primary parameter name
        String primaryArg = goNames.getParameterStyleName(primaryParameter.getName());
        //      get primary parameter type name
        GoTypeReference primaryParameterGoTypeReference = goNames.getTypeReference(primaryParameter.getType());
        buffer.addImports(primaryParameterGoTypeReference.getImports());
        //      Generate function doc
        generateActionDoc(method, (Parameter p) -> p.isIn() && p.isOut());
        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s(",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        //      Generate function parameters definition
        buffer.startBlock();
        buffer.addLine("%1$s %2$s,", primaryArg, primaryParameterGoTypeReference.getText());
        secondaryParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) (");
        //      Generate the output parameters
        buffer.startBlock();
        this.generateOutputParameter(primaryParameter);
        buffer.addLine("error) {");
        buffer.endBlock();
        //      Generate function ending
        buffer.endBlock();

        // Start body:
        buffer.startBlock();

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
        buffer.addLine("ovResp, err := op.internalUpdate(%1$s, headers, query, wait)", primaryArg);

        //      Generate ovResponse parsing
        this.generateOvResponseParsing(primaryParameter);

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
        //      Get function output parameters
        //      remove() function has no out parameters

        //      Generate function definition
        buffer.addLine(
            "func (op *%1$s) %2$s(",
            serviceName.getClassName(),
            goNames.getMethodStyleName(methodName));
        buffer.startBlock();
        inParameters.forEach(this::generateFormalParameter);
        buffer.addLine("headers map[string]string,");
        buffer.addLine("query map[string]string,");
        buffer.addLine("wait bool) (");
        //      Generate the output parameters
        buffer.startBlock();
        buffer.addLine("error) {");
        buffer.endBlock();
        //      Generate function ending
        buffer.endBlock();

        // Begin body:
        buffer.startBlock();

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
        buffer.addLine("_, err := op.internalRemove(headers, query, wait)");
        buffer.addLine("return err");

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

    private void generateOutputParameter(Parameter parameter) {
        // Get parameter type
        GoTypeReference goTypeReference = goNames.getTypeReference(parameter.getType());
        buffer.addImports(goTypeReference.getImports());

        // Get parameter type name
        buffer.addLine(
            "%1$s,", goTypeReference.getText());
    }

    private void generateUrlParameter(Parameter parameter) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        String arg = goNames.getParameterStyleName(name);
        String tag = schemaNames.getSchemaTagName(name);
        buffer.addLine("query[\"%1$s\"] = fmt.Sprintf(\"%%v\", %2$s)", tag, arg);
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
        GoClassName serviceName = goNames.getServiceName(locator.getService());
        String memberName = goNames.getMemberStyleName(locator.getName());
        buffer.addLine("%1$sServ  *%2$s", memberName, serviceName.getClassName());
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
        buffer.addLine("func (op *%1$s) Service(path string) (interface{}, error) {", serviceName.getClassName());
        buffer.startBlock();
        buffer.addLine("if path == \"\" {");
        buffer.startBlock();
        buffer.addLine("return op, nil");
        buffer.endBlock();
        buffer.addLine("}");

        // Generate the code that checks if the path corresponds to any of the locators without parameters:
        service.locators().filter(x -> x.getParameters().isEmpty()).sorted().forEach(locator -> {
            Name name = locator.getName();
            String segment = getPath(name);
            buffer.addLine("if path == \"%1$s\" {", segment);
            buffer.startBlock();
            buffer.addLine(  "return op.%1$sService(), nil", goNames.getMethodStyleName(name));
            buffer.endBlock();
            buffer.addLine("}");
            buffer.addLine("if strings.HasPrefix(path, \"%1$s/\") {", segment);
            buffer.addImport("strings");
            buffer.startBlock();
            buffer.addLine(
                "return op.%1$sService().Service(path[%2$d:])",
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
            buffer.addLine("index := strings.Index(path, \"/\")");
            buffer.addLine("if index == -1 {");
            buffer.startBlock();
            buffer.addLine("return *(op.%1$sService(path)), nil", goNames.getMemberStyleName(name));
            buffer.endBlock();
            buffer.addLine("}");
            buffer.addLine(
                "return op.%1$sService(path[:index]).Service(path[index + 1:])",
                goNames.getMemberStyleName(name)
            );
        }
        else {
            buffer.addLine("return nil, fmt.Errorf(\"The path <%%s> doesn't correspond to any service\", path)");
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

    private void generateXmlUnmarshal(Parameter outParameter) {
        buffer.addImport("encoding/xml");
        buffer.addLine("xml.Unmarshal([]byte(ovResp.Body), &%1$s)",
            goNames.getVariableStyleName(outParameter.getName()));
    }

    private void generateOvResponseParsing(Parameter outParameter) {
        GoTypeReference outParamTypeReference = goNames.getTypeReference(outParameter.getType());
        String outParamNameAsVar = goNames.getVariableStyleName(outParameter.getName());
        Model model = outParameter.getType().getModel();

        Type outParameterType = outParameter.getType();
        if (outParameterType instanceof PrimitiveType) {
            if (outParameterType == model.getBooleanType()) {
                buffer.addLine("if err != nil {");
                buffer.startBlock();
                buffer.addLine("return false, err");
                buffer.endBlock();
                buffer.addLine("}");
                buffer.addImport("strconv");
                buffer.addLine("return strconv.ParseBool(ovResp.Body)");
            } else if (outParameterType == model.getIntegerType()) {
                buffer.addLine("if err != nil {");
                buffer.startBlock();
                buffer.addLine("return 0, err");
                buffer.endBlock();
                buffer.addLine("}");
                buffer.addImport("strconv");
                buffer.addLine("return strconv.ParseInt(ovResp.Body, 10, 64)");
            } else if (outParameterType == model.getDecimalType()) {
                buffer.addLine("if err != nil {");
                buffer.startBlock();
                buffer.addLine("return 0, err");
                buffer.endBlock();
                buffer.addLine("}");
                buffer.addImport("strconv");
                buffer.addLine("return srconv.ParseFloat(ovResp.Body, 10, 64");
            } else if (outParameterType == model.getStringType()) {
                buffer.addLine("if err != nil {");
                buffer.startBlock();
                buffer.addLine("return \"\", err");
                buffer.endBlock();
                buffer.addLine("}");
                buffer.addLine("return ovResp.Body, nil");
            } else {
                throw new IllegalArgumentException(
                    "XMLParsing: Don't know how to build reference for primitive type \"" + outParameterType + "\""
                );
            }
        } else if (outParameterType instanceof StructType) {
            buffer.addLine("if err != nil {");
            buffer.startBlock();
            buffer.addLine("return nil, err");
            buffer.endBlock();
            buffer.addLine("}");
            buffer.addLine("var %1$s %2$s", outParamNameAsVar, goNames.getTypeName(outParameter.getType()).getClassName());
            this.generateXmlUnmarshal(outParameter);
            buffer.addLine("return &%1$s, nil", outParamNameAsVar);
        } else if (outParameterType instanceof ListType) {
            buffer.addLine("if err != nil {");
            buffer.startBlock();
            buffer.addLine("return nil, err");
            buffer.endBlock();
            buffer.addLine("}");

            ListType outParamListType = (ListType) outParameterType;
            if (outParamListType.getElementType() instanceof StructType) {
                String outParamListTypeStr = outParamTypeReference.getText().replace("[]", "");
                String outParamListAttriTypeStr = String.format("%ss", outParamListTypeStr);
                outParamListTypeStr = String.format("%ss", outParamListTypeStr);

                buffer.addLine("var %1$s %2$s", outParamNameAsVar, outParamListTypeStr);
                generateXmlUnmarshal(outParameter);
                buffer.addLine("return %1$s, nil",
                    String.format("%s.%s", outParamNameAsVar, outParamListAttriTypeStr));
            } else if (outParamListType.getElementType() == model.getStringType()) {
                buffer.addLine("return []string{ovResp.Body}, nil");
            } else {
                throw new IllegalArgumentException(
                    "XMLParsing: Don't know how to build reference for type: List-Of- \"" + outParameterType + "\""
                );
            }
        }
    }

    /**
     * Return the ONLY (Java) one output parameter of method
     */
    private Parameter getOutParameter(Method method) {
        return method.parameters()
            .filter(x -> x.isOut())
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
        String varTypeSuffix = "";
        if (GoTypes.isGoPrimitiveType(parameter.getType())) {
            varTypeSuffix = "&";
        }
        buffer.addLine("%1$s: %2$s%3$s,", memberName, varTypeSuffix, parameterName);
    }

}
