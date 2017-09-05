/*
Copyright (c) 2017 Joey <majunjiev@gmail.com>.

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
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.Concept;
import org.ovirt.api.metamodel.concepts.EnumType;
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
import org.ovirt.api.metamodel.tool.Names;
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

    // The buffer used to generate the code:
    private GoBuffer buffer;

    // Reference to the object used to calculate Go types:
    @Inject
    private GoTypes goTypes;
    
    // Reference to the objects used to generate the code:
    @Inject
    private Names names;

    @Inject private SchemaNames schemaNames;

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
        model.services().forEach(this::generateService);
    }

    private void generateService(Service service) {
        // Begin class:
        GoClassName serviceName = goNames.getServiceName(service);
        
        // Generate struct definition
        generateDoc(service);
        buffer.addLine("type %1$s struct {", serviceName.getClassName());

        // Generate struct members definition
        //      with Service struct mixin
        buffer.addLine("baseService");

        // Generate struct ending
        buffer.addLine("}");
        buffer.addLine();

        // Generate the service struct constructor by Newer function
        this.generateConstructor(serviceName);

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

    private void generateConstructor(GoClassName serviceName) {
        buffer.addLine(
            "func %1$s(connection *Connection, path string) *%2$s {",
            goTypes.getServiceConstructorFuncName(serviceName),
            serviceName.getClassName());

        // Inititalize struct
        buffer.addLine("var result %1$s", serviceName.getClassName());
        buffer.addLine("result.Connection = connection");
        buffer.addLine("result.Path = path");
        buffer.addLine("return &result");

        // Generate constructor ending
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateMethod(Method method, Service service) {
        // Generate the request and response struct for method
        generateRequest(method, service);
        generateResponse(method, service);

        // Generate the method using Request/Response
        Name methodName = method.getName();
        String request = getRequestClassName(method, service);
        String methodNameString = goNames.getPublicMethodStyleName(methodName);
        GoClassName serviceClassName = goNames.getServiceName(service);
        buffer.addLine("func (p *%1$s) %2$s() *%3$s {",
            serviceClassName.getClassName(), methodNameString, request);
        buffer.addLine("return &%1$s{%2$s: p}",
            request,
            goNames.getPrivateMemberStyleName(serviceClassName.getClassName()));
        buffer.addLine("}");
    }

    private void generateRequest(Method method, Service service) {
        // Begin class
        Name methodName = method.getName();
        String request = getRequestClassName(method, service);

        buffer.addLine("type %1$s struct {", request);

        // Service itself
        buffer.addLine("%1$s *%2$s", 
            goNames.getPrivateMemberStyleName(goNames.getServiceName(service).getClassName()),
            goNames.getServiceName(service).getClassName());

        //      Generate common parameters
        generateRequestCommonParameter();
        
        //      Generate the input parameters
        method.parameters()
            .filter(Parameter::isIn)
            .sorted().forEach(this::generateRequestParameter);
        
        buffer.addLine("}");
        // End class

        // Generate methods to set common parameters
        generateRequestCommonParameterMethod(request);

        // Generate methods to set input parameters
        List<Parameter> parameterList = method.parameters()
            .filter(Parameter::isIn)
            .sorted().collect(toCollection(ArrayList::new));
        for (Parameter para : parameterList) {
            generateRequestParameterMethod(para, request);
        }

        // Generate send method:
        buffer.addLine("func (p *%1$s) Send() (*%2$s, error) {",
            request, getResponseClassName(method, service));

        // Generate method code based on response type:
        if (ADD.equals(methodName)) {
            generateAddRequestImplementation(method, service);
        }
        else if (GET.equals(methodName) || LIST.equals(methodName)) {
            generateListRequestImplementation(method, service);
        }
        else if (REMOVE.equals(methodName)) {
            generateRemoveRequestImplementation(method, service);
        }
        else if (UPDATE.equals(methodName)) {
            generateUpdateRequestImplementation(method, service);
        }
        else {
            generateActionRequestImplementation(method, service);
        }

        // End send method:
        buffer.addLine("}");
        buffer.addLine();

    }

    private void generateRequestCommonParameter() {
        // Add common header and query parameters
        buffer.addLine("header map[string]string");
        buffer.addLine("query map[string]string");
    }

    private void generateRequestCommonParameterMethod(String requestClassName) {
        String[] commonParameters = new String[]{"header", "query"};

        for (String para : commonParameters) {
            buffer.addLine("func (p *%1$s) %2$s(key, value string) *%1$s {",
                requestClassName, GoNames.capitalize(para));
            buffer.addLine(  "if p.%1$s == nil {", para);
            buffer.addLine(    "p.%1$s = make(map[string]string)", para);
            buffer.addLine(  "}");
            buffer.addLine(  "p.%1$s[key] = value", para);
            buffer.addLine(  "return p");
            buffer.addLine("}");
            buffer.addLine();
        }
    }

    private void generateRequestParameter(Parameter parameter) {
        // Get parameter name
        Name parameterName = parameter.getName();
        GoTypeReference goTypeReference = goNames.getRefTypeReference(parameter.getType());
        buffer.addImports(goTypeReference.getImports());

        String arg = goNames.getParameterStyleName(parameterName);
        // Get parameter type name
        buffer.addLine(
            "%1$s %2$s", arg, goTypeReference.getText());
    }

    private void generateRequestParameterMethod(Parameter parameter, String requestClassName) {
        Type paraType = parameter.getType();
        GoTypeReference paraTypeReference = goNames.getTypeReference(paraType);
        buffer.addImports(paraTypeReference.getImports());
        String paraName = goNames.getParameterStyleName(parameter.getName());
        String paraMethodName = goNames.getPublicMethodStyleName(parameter.getName());

        buffer.addLine("func (p *%1$s) %2$s(%3$s %4$s) *%1$s{",
            requestClassName, paraMethodName, paraName, paraTypeReference.getText());
        if (goTypes.isGoPrimitiveType(paraType)) {
            buffer.addLine("p.%1$s = &%1$s", paraName);
        } else {
            buffer.addLine("p.%1$s = %1$s", paraName);
        }
        buffer.addLine("return p");
        
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateAddRequestImplementation(Method method, Service service) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        buffer.addLine("rawURL := fmt.Sprintf(\"%%s%%s\", p.%1$s.Connection.URL(), p.%1$s.Path)",
            goNames.getPrivateMemberStyleName(serviceClassName));
        buffer.addImport("net/url");
        buffer.addLine("values := make(url.Values)");

        getSecondaryParameters(method)
            .forEach(this::generateRequestParameterQueryBuilder);
        generateAdditionalQueryParameters();
        
        // Generate the final URL
        buffer.addLine("if len(values) > 0 {");
        buffer.addLine(  "rawURL = fmt.Sprintf(\"%%s?%%s\", rawURL, values.Encode())");
        buffer.addLine("}");

        // Generate the net/http request.Body (via bytes.Buffer)
        buffer.addImport("bytes");
        buffer.addLine("var body bytes.Buffer");
        generateWriteRequestBody(getFirstParameter(method));

        // Construct the net/http request
        buffer.addImport("net/http");
        buffer.addLine("req, err := http.NewRequest(\"POST\", rawURL, &body)");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");

        generateCommonRequestImplementation(method, service, new String[]{"200", "201", "202"});
        generateResponseParseImplementation(method, service);
	}

    private void generateListRequestImplementation(Method method, Service service) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        buffer.addLine("rawURL := fmt.Sprintf(\"%%s%%s\", p.%1$s.Connection.URL(), p.%1$s.Path)",
            goNames.getPrivateMemberStyleName(serviceClassName));
        buffer.addImport("net/url");
        buffer.addLine("values := make(url.Values)");
        method.parameters()
            .filter(Parameter::isIn)
            .filter(p -> p.getType() instanceof PrimitiveType)
            .sorted()
            .forEach(this::generateRequestParameterQueryBuilder);
        generateAdditionalQueryParameters();
        // Generate the final URL
        buffer.addLine("if len(values) > 0 {");
        buffer.addLine(  "rawURL = fmt.Sprintf(\"%%s?%%s\", rawURL, values.Encode())");
        buffer.addLine("}");

        // Construct the net/http request
        buffer.addImport("net/http");
        buffer.addLine("req, err := http.NewRequest(\"GET\", rawURL, nil)");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");

        generateCommonRequestImplementation(method, service, new String[]{"200"});
        generateResponseParseImplementation(method, service);
    }

    private void generateRemoveRequestImplementation(Method method, Service service) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        buffer.addLine("rawURL := fmt.Sprintf(\"%%s%%s\", p.%1$s.Connection.URL(), p.%1$s.Path)",
            goNames.getPrivateMemberStyleName(serviceClassName));
        buffer.addImport("net/url");
        buffer.addLine("values := make(url.Values)");
        method.parameters()
            .filter(Parameter::isIn)
            .filter(p -> p.getType() instanceof PrimitiveType)
            .sorted()
            .forEach(this::generateRequestParameterQueryBuilder);
        generateAdditionalQueryParameters();
        // Generate the final URL
        buffer.addLine("if len(values) > 0 {");
        buffer.addLine(  "rawURL = fmt.Sprintf(\"%%s?%%s\", rawURL, values.Encode())");
        buffer.addLine("}");

        // Construct the net/http request
        buffer.addImport("net/http");
        buffer.addLine("req, err := http.NewRequest(\"DELETE\", rawURL, nil)");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");

        generateCommonRequestImplementation(method, service, new String[]{"200"});
        generateResponseParseImplementation(method, service);
    }

    private void generateUpdateRequestImplementation(Method method, Service service) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        buffer.addLine("rawURL := fmt.Sprintf(\"%%s%%s\", p.%1$s.Connection.URL(), p.%1$s.Path)",
            goNames.getPrivateMemberStyleName(serviceClassName));
        buffer.addImport("net/url");
        buffer.addLine("values := make(url.Values)");

        getSecondaryParameters(method)
            .forEach(this::generateRequestParameterQueryBuilder);
        generateAdditionalQueryParameters();
        
        // Generate the final URL
        buffer.addLine("if len(values) > 0 {");
        buffer.addLine(  "rawURL = fmt.Sprintf(\"%%s?%%s\", rawURL, values.Encode())");
        buffer.addLine("}");

        // Generate the net/http request.Body (via bytes.Buffer)
        buffer.addImport("bytes");
        buffer.addLine("var body bytes.Buffer");
        generateWriteRequestBody(getFirstParameter(method));

        // Construct the net/http request
        buffer.addImport("net/http");
        buffer.addLine("req, err := http.NewRequest(\"PUT\", rawURL, &body)");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");

        generateCommonRequestImplementation(method, service, new String[]{"200"});
        generateResponseParseImplementation(method, service);
    }

    private void generateActionRequestImplementation(Method method, Service service) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        buffer.addLine("rawURL := fmt.Sprintf(\"%%s%%s/%1$s\", p.%2$s.Connection.URL(), p.%2$s.Path)",
            getPath(method.getName()),
            goNames.getPrivateMemberStyleName(serviceClassName));
        buffer.addLine("actionBuilder := NewActionBuilder()");
        method.parameters()
            .filter(Parameter::isIn)
            .sorted()
            .forEach(parameter -> {
                String paraArgName = goNames.getParameterStyleName(parameter.getName());
                String paraMethodName = goNames.getPublicMemberStyleName(parameter.getName());
                if (goTypes.isGoPrimitiveType(parameter.getType())) {
                    buffer.addLine("actionBuilder.%1$s(*p.%2$s);", paraMethodName, paraArgName);
                } else {
                    buffer.addLine("actionBuilder.%1$s(p.%2$s);", paraMethodName, paraArgName);
                }
            });
        buffer.addLine("action, err := actionBuilder.Build()");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");
        buffer.addImport("net/url");
        buffer.addLine("values := make(url.Values)");

        generateAdditionalQueryParameters();
        
        // Generate the final URL
        buffer.addLine("if len(values) > 0 {");
        buffer.addLine(  "rawURL = fmt.Sprintf(\"%%s?%%s\", rawURL, values.Encode())");
        buffer.addLine("}");

        // Generate the net/http request.Body (via bytes.Buffer)
        buffer.addImport("bytes");
        buffer.addLine("var body bytes.Buffer");
        buffer.addLine("writer := NewXMLWriter(&body)");
        buffer.addLine("err = XMLActionWriteOne(writer, action, \"\")");
        buffer.addLine("writer.Flush()");
        
        // Construct the net/http request
        buffer.addImport("net/http");
        buffer.addLine("req, err := http.NewRequest(\"PUT\", rawURL, &body)");
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");

        generateCommonRequestImplementation(method, service, null);
        // Check action
        List<Parameter> parameters = method.parameters().filter(Parameter::isOut).collect(Collectors.toList());
        if (parameters.isEmpty()) {
            buffer.addLine("_, errCheckAction := CheckAction(resp)");
        } else {
            buffer.addLine("action, errCheckAction := CheckAction(resp)");
        }
        buffer.addLine("if errCheckAction != nil {");
        buffer.addLine(  "return nil, errCheckAction");
        buffer.addLine("}");
        
        if (parameters.isEmpty()) {
            buffer.addLine("return new(%1$s), nil", getResponseClassName(method, service));
        } else {
            Parameter paraFirst = parameters.get(0);
            buffer.addLine("result := action.%1$s()",
                goTypes.getMemberMustGetterMethodName(paraFirst.getName()));
            Type paraType = paraFirst.getType();
            if (goTypes.isGoPrimitiveType(paraType) || paraType instanceof EnumType) {
                buffer.addLine("return &%1$s{%2$s: &result}, nil",
                    getResponseClassName(method, service),
                    goNames.getPrivateMemberStyleName(paraFirst.getName()));
            }
            else {
                buffer.addLine("return &%1$s{%2$s: result}, nil",
                    getResponseClassName(method, service),
                    goNames.getPrivateMemberStyleName(paraFirst.getName()));
            }
        }
    }

    private void generateRequestParameterQueryBuilder(Parameter parameter) {
        String tag = schemaNames.getSchemaTagName(parameter.getName());
        String value = goNames.getPrivateMemberStyleName(parameter.getName());
        buffer.addLine("if p.%1$s != nil {", value);
        buffer.addLine(  "values[\"%1$s\"] = []string{fmt.Sprintf(\"%%v\", *p.%2$s)}", tag, value);
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateAdditionalQueryParameters() {
        buffer.addLine("if p.query != nil {");
        buffer.addLine(  "for k, v := range p.query {");
        buffer.addLine(    "values[k] = []string{v}");
        buffer.addLine(  "}");
        buffer.addLine("}");
    }

    private void generateWriteRequestBody(Parameter parameter) {
        // var body bytes.Buffer
        if (parameter != null) {
            Type type = parameter.getType();
            buffer.addLine("writer := NewXMLWriter(&body)");
            if (type instanceof StructType) {
                buffer.addLine("err := %1$s(writer, p.%2$s, \"\")",
                    goTypes.getXmlWriteOneFuncName(type),
                    goNames.getParameterStyleName(parameter.getName())
                );
            }
            else if (type instanceof ListType) {
                ListType listType = (ListType) type;
                Type elementType = listType.getElementType();
                buffer.addLine("err := %1$s(writer, p.%2$s, \"\", \"\")",
                    goTypes.getXmlWriteManyFuncName(elementType),
                    goNames.getParameterStyleName(parameter.getName())
                );
            }
            buffer.addLine("if err != nil {");
            buffer.addLine(" return nil, err");
            buffer.addLine("}");
            buffer.addLine("writer.Flush()");
        }
    }

    private void generateCommonRequestImplementation(Method method, Service service, String[] codes) {
        String serviceClassName = goNames.getServiceName(service).getClassName();
        String serviceAsPrivateMemberName = goNames.getPrivateMemberStyleName(serviceClassName);
        
        generateAdditionalHeadersParameters();
        buffer.addLine("req.Header.Add(\"User-Agent\", fmt.Sprintf(\"GoSDK/%%s\", SDK_VERSION))");
        buffer.addLine("req.Header.Add(\"Version\", \"4\")");
        buffer.addLine("req.Header.Add(\"Content-Type\", \"application/xml\")");
        buffer.addLine("req.Header.Add(\"Accept\", \"application/xml\")");
        buffer.addLine("rawAuthStr := fmt.Sprintf(\"%%s:%%s\", p.%1$s.Connection.username, p.%1$s.Connection.password)",
            serviceAsPrivateMemberName);
        buffer.addCommentLine("Generate base64(username:password)");
        buffer.addImport("encoding/base64");
        buffer.addLine("auth := fmt.Sprintf(\"Basic %%s\", base64.StdEncoding.EncodeToString([]byte(rawAuthStr)))");
        buffer.addLine("req.Header.Add(\"Authorization\", auth)");
        // Send the request and wait for the response
        buffer.addCommentLine("Send the request and wait for the response");
        buffer.addLine("resp, err := p.%1$s.Connection.client.Do(req)",
            serviceAsPrivateMemberName);
        buffer.addLine("if err != nil {");
        buffer.addLine(  "return nil, err");
        buffer.addLine("}");;
        buffer.addLine("defer resp.Body.Close()");

        // Check the response status code
        if (codes != null && codes.length > 0) {
            buffer.addLine("if !Contains(resp.StatusCode, []int{%1$s}) {", String.join(",", codes));
            buffer.addLine("  return nil, CheckFault(resp)");
            buffer.addLine("}");
        }

        // Read resp.Body (if method is ActionMethod, no need the resp reading)
        Name methodName = method.getName();
        if (methodName.equals(ADD) || methodName.equals(GET) || 
            methodName.equals(LIST) || methodName.equals(REMOVE) ||
            methodName.equals(UPDATE)) {
                buffer.addImport("io/ioutil");
                List<Parameter> parameters = method.parameters().filter(Parameter::isOut).collect(Collectors.toList());
                if (parameters.isEmpty()) {
                    buffer.addLine("_, errReadBody := ioutil.ReadAll(resp.Body)");
                } else {
                    buffer.addLine("respBodyBytes, errReadBody := ioutil.ReadAll(resp.Body)");
                }
                buffer.addLine("if errReadBody != nil {");
                buffer.addLine(  "return nil, errReadBody");
                buffer.addLine("}");
        }
    }


    private void generateResponseParseImplementation(Method method, Service service) {
        List<Parameter> parameters = method.parameters().filter(Parameter::isOut).collect(Collectors.toList());
        if (parameters.isEmpty()) {
            buffer.addLine("return new(%1$s), nil", getResponseClassName(method, service));
        } else {
            for (Parameter para : parameters) {
                generateResponseParameterParseImplementation(para, service);
            }
        }
    }

    private void generateResponseParameterParseImplementation(Parameter parameter, Service service) {
        Type type = parameter.getType();
        String response = getResponseClassName(parameter.getDeclaringMethod(), service);

        buffer.addLine("reader := NewXMLReader(respBodyBytes)");
        if (type instanceof PrimitiveType) {
            Model model = type.getModel();
            if (type == model.getBooleanType()) {
                buffer.addLine("result, err := reader.ReadBool(nil)");
            }
            else if (type == model.getIntegerType()) {
                buffer.addLine("result, err := reader.ReadInt64(nil)");
            }
            else if (type == model.getDecimalType()) {
                buffer.addLine("result, err := reader.ReadFloat64(nil)");
            }
            else if (type == model.getStringType()) {
                buffer.addLine("result, err := reader.ReadString(nil)");
            }
            else if (type == model.getDateType()) {
                buffer.addLine("result, err := reader.ReadTime(nil)");
            }
            else {
                throw new IllegalArgumentException(
                    "XMLParsing: Don't know how to build reference for primitive type \"" + type + "\""
                );
            }
        }
        else if (type instanceof StructType) {
            buffer.addLine("result, err := %1$s(reader, nil, \"\")", goTypes.getXmlReadOneFuncName(type));
        } else if (type instanceof ListType) {
            ListType listype = (ListType) type;
            Type elementType = listype.getElementType();
            buffer.addLine("result, err := %1$s(reader, nil)", goTypes.getXmlReadManyFuncName(elementType));
        }
        buffer.addLine("if err != nil {");
        buffer.addLine("  return nil, err");
        buffer.addLine("}");
        buffer.addLine("return &%1$s{%2$s: result}, nil",
            response, goNames.getPrivateMemberStyleName(parameter.getName()));
    }

    private void generateAdditionalHeadersParameters() {
        buffer.addLine();
        buffer.addLine("if p.header != nil {");
        buffer.addLine(  "for hk, hv := range p.header {");
        buffer.addLine(    "req.Header.Add(hk, hv)");
        buffer.addLine(  "}");
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateResponse(Method method, Service service) {
        String response = getResponseClassName(method, service);
        buffer.addLine("type %1$s struct {", response);
        method.parameters()
            .filter(Parameter::isOut)
            .sorted()
            .forEach(this::generateResponseParameter);
        buffer.addLine("}");

        // Generate Repsonse parameter getter method
        List<Parameter> parameters = method.parameters()
            .filter(Parameter::isOut)
            .sorted()
            .collect(Collectors.toList());
        for (Parameter para : parameters) {
            generateResponseParameterGetterMethod(para, service);
            generateResponseParameterMustGetterMethod(para, service);
        }
    }

    private void generateResponseParameter(Parameter parameter) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        String memberName = goNames.getPrivateMemberStyleName(name);
        GoTypeReference reference = goNames.getRefTypeReference(type);
        buffer.addLine("%1$s %2$s", memberName, reference.getText());
    }

    private void generateResponseParameterGetterMethod(Parameter parameter, Service service) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        GoTypeReference reference = goNames.getTypeReference(type);
        String response = getResponseClassName(parameter.getDeclaringMethod(), service);
        buffer.addLine("func (p *%1$s) %2$s() (%3$s, bool) {",
            response,
            goTypes.getMemberGetterMethodName(name),
            reference.getText()
            );
        buffer.addLine(" if p.%1$s != nil {", goNames.getPrivateMemberStyleName(name));
        if (goTypes.isGoPrimitiveType(type) || type instanceof EnumType) {
            buffer.addLine("  return *p.%1$s, true", goNames.getPrivateMemberStyleName(name));
            buffer.addLine(" }");
        }
        else {
            buffer.addLine("  return p.%1$s, true", goNames.getPrivateMemberStyleName(name));
            buffer.addLine(" }");
        }
        buffer.addLine(" var zero %1$s", reference.getText());
        buffer.addLine(" return zero, false");
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateResponseParameterMustGetterMethod(Parameter parameter, Service service) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        GoTypeReference reference = goNames.getTypeReference(type);
        String response = getResponseClassName(parameter.getDeclaringMethod(), service);
        buffer.addLine("func (p *%1$s) %2$s() %3$s {",
            response,
            goTypes.getMemberMustGetterMethodName(name),
            reference.getText()
            );
        if (goTypes.isGoPrimitiveType(type) || type instanceof EnumType) {
            buffer.addLine(" return *p.%1$s", goNames.getPrivateMemberStyleName(name));
        }
        else {
            buffer.addLine(" return p.%1$s", goNames.getPrivateMemberStyleName(name));
        }

        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateStr(Service service) {
        GoClassName serviceName = goNames.getServiceName(service);
        buffer.addImport("fmt");
        buffer.addLine("func (op *%1$s) String() string {", serviceName.getClassName());
        buffer.addLine(  "return fmt.Sprintf(\"%1$s:%%s\", op.Path)", serviceName.getClassName());
        buffer.addLine("}");
        buffer.addLine();
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
        String methodName = goNames.getPublicMemberStyleName(locator.getName());
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

        buffer.addImport("fmt");
        buffer.addLine(
            "return %1$s(op.Connection, fmt.Sprintf(\"%%s/%%s\", op.Path, %2$s))",
            goTypes.getServiceConstructorFuncName(locatorServiceName),
            argName);
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateLocatorWithoutParameters(Locator locator, Service service) {
        String methodName = goNames.getPublicMethodStyleName(locator.getName());
        String urlSegment = getPath(locator.getName());
        GoClassName locatorServiceName = goNames.getServiceName(locator.getService());
        generateDoc(locator);

        // Get receiver class
        GoClassName receiverClassName = goNames.getServiceName(service);
        // Generate *Service function
        buffer.addLine("func (op *%1$s) %2$sService() *%3$s {",
            receiverClassName.getClassName(), methodName, locatorServiceName.getClassName());

        buffer.addImport("fmt");
        buffer.addLine(
            "return %1$s(op.Connection, fmt.Sprintf(\"%%s/%2$s\", op.Path))",
            goTypes.getServiceConstructorFuncName(locatorServiceName),
            urlSegment);
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
        buffer.addLine(  "if path == \"\" {");
        buffer.addLine(    "return op, nil");
        buffer.addLine(  "}");

        // Generate the code that checks if the path corresponds to any of the locators without parameters:
        service.locators().filter(x -> x.getParameters().isEmpty()).sorted().forEach(locator -> {
            Name name = locator.getName();
            String segment = getPath(name);
            buffer.addImport("strings");
            buffer.addLine("if path == \"%1$s\" {", segment);
            buffer.addLine(  "return op.%1$sService(), nil", goNames.getPublicMethodStyleName(name));
            buffer.addLine("}");
            buffer.addLine("if strings.HasPrefix(path, \"%1$s/\") {", segment);
            buffer.addLine(
                "return op.%1$sService().Service(path[%2$d:])",
                goNames.getPublicMemberStyleName(name),
                segment.length() + 1
            );
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
            buffer.addLine(  "return *(op.%1$sService(path)), nil", goNames.getPublicMemberStyleName(name));
            buffer.addLine("}");
            buffer.addLine(
                "return op.%1$sService(path[:index]).Service(path[index + 1:])",
                goNames.getPublicMemberStyleName(name)
            );
        }
        else {
            buffer.addLine("return nil, fmt.Errorf(\"The path <%%s> doesn't correspond to any service\", path)");
        }

        // End method:
        buffer.addLine("}");
        buffer.addLine();
    }

    private String getPath(Name name) {
        return name.words().map(String::toLowerCase).collect(joining());
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

    private String getRequestClassName(Method method, Service service) {
        return goNames.getServiceName(service).getClassName() + 
            goNames.getClassStyleName(method.getName()) + "Request";
    }

    private String getResponseClassName(Method method, Service service) {
        return goNames.getServiceName(service).getClassName() + 
            goNames.getClassStyleName(method.getName()) + "Response";
    }

}
