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

import static java.util.stream.Collectors.toCollection;
import static java.util.stream.Collectors.toList;
import static java.util.stream.Collectors.toSet;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.stream.Stream;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.EnumType;
import org.ovirt.api.metamodel.concepts.EnumValue;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.StructMember;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.Names;

/**
 * This class is responsible for generating the classes that represent the types of the model.
 */
public class TypesGenerator implements GoGenerator {
    // The directory were the output will be generated:
    protected File out;

    // Reference to the objects used to generate the code:
    @Inject
    private Names names;
    @Inject
    private GoNames goNames;

    // The buffer used to generate the code:
    private GoBuffer buffer;

    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Prepare the buffer:
        buffer = new GoBuffer();
        buffer.setPackageName(goNames.getTypesPackageName());

        // Generate the code:
        generateTypes(model);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new RuntimeException("Error writing types module", exception);
        }
    }

    private void generateTypes(Model model) {
        // Get the list of struct types:
        List<StructType> structs = model.types()
            .filter(StructType.class::isInstance)
            .map(StructType.class::cast)
            .sorted()
            .collect(toList());

        // Generate the type:
        structs.forEach(this::generateStruct);

        // Generate the builder:
        structs.forEach(this::generateStructBuilder);

        // Enum types don't need any special order, so we sort them only by name:
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .sorted()
            .forEach(this::generateEnum);
    }

    private void generateStruct(StructType type) {
        // Begin class:
        GoClassName typeName = goNames.getTypeName(type);
        Type base = type.getBase();
        // Define []Struct
        buffer.addLine("type %1$ss struct {", typeName.getClassName());
        //  Add xml.Name
        buffer.addImport("encoding/xml");
        buffer.addLine(  "XMLName xml.Name `xml:\"%1$ss\"`", typeName.getClassName().toLowerCase());
        buffer.addLine(  "%1$ss []%1$s `xml:\"%2$s,omitempty\"`", typeName.getClassName(), goNames.getTagStyleName(type.getName()));
        buffer.addLine("}");
        buffer.addLine();
        // Define Struct
        buffer.addLine("type %1$s struct {", typeName.getClassName());
        // Ignore Base-class mixin, fill in all
        buffer.addLine("OvStruct");

        // Constructor with a named parameter for each attribute and link:
        Set<StructMember> allMembers = Stream.concat(type.attributes(), type.links())
            .collect(toSet());
        Set<StructMember> declaredMembers = Stream.concat(type.declaredAttributes(), type.declaredLinks())
            .collect(toSet());
        allMembers.addAll(declaredMembers);
        allMembers.stream().sorted().forEach(this::generateMemberFormalParameter);

        buffer.addLine();

        // End class:
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateStructBuilder(StructType type) {
        // Get Type names
        GoClassName typeName = goNames.getTypeName(type);
        String typeClassName = typeName.getClassName();
        String typePrivateClassName = typeName.getPrivateClassName();
        String typePrivateMemberName = goNames.getPrivateMemberStyleName(type.getName());
        // Get struct members
        Set<StructMember> allMembers = Stream.concat(type.attributes(), type.links())
            .collect(toSet());
        Set<StructMember> declaredMembers = Stream.concat(type.declaredAttributes(), type.declaredLinks())
            .collect(toSet());
        allMembers.addAll(declaredMembers);

        // Begin class:
        buffer.addLine("type %1$sBuilder struct {", typePrivateClassName);
        //      Add properties of TypeBuilder
        buffer.addLine(  "%1$s *%2$s", typePrivateMemberName, typeName.getClassName());
        buffer.addLine(  "err error");
        // End class:
        buffer.addLine("}");
        buffer.addLine();

        // Define NewStructBuilder function
        buffer.addLine("func New%1$sBuilder() *%2$sBuilder {",
            typeClassName, typePrivateClassName);
        buffer.addLine("return &%1$sBuilder{%2$s: &%3$s{}, err: nil}",
            typePrivateClassName,
            typePrivateMemberName,
            typeClassName
            );
        buffer.addLine("}");
        buffer.addLine();

        // Construct TypeBuilder methods for member-settings
        List<StructMember> members = allMembers.stream().sorted().collect(toCollection(ArrayList::new));
        for (StructMember member : members) {
            this.generateBuilderMemberMethods(type, member);
        }

        // Generate Build method
        buffer.addLine("func (builder *%1$sBuilder) Build() (*%2$s, error) {", typePrivateClassName, typeClassName);
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return nil, builder.err");
        buffer.addLine(  "}");
        buffer.addLine(  "return builder.%1$s, nil", typePrivateMemberName);
        buffer.addLine("}");
    }

    private void generateEnum(EnumType type) {
        // Begin class:
        GoClassName typeName = goNames.getTypeName(type);

        // Type declaration
        buffer.addLine("type %1$s string", typeName.getClassName());

        // Type definition
        buffer.addLine("const (");
        type.values().sorted().forEach(this::generateEnumValue);
        buffer.addLine(")");

        // End definition:
        buffer.addLine();
    }

    private void generateEnumValue(EnumValue value) {
        Name name = value.getName();
        String constantName = goNames.getConstantStyleName(name);
        String className = goNames.getTypeName(value.getDeclaringType()).getClassName();
        String constantValue = names.getLowerJoined(name, "_");

        // To avoid constant-name conflict, eg: VMSTATUS_DOWN
        constantName = className.toUpperCase() + "_" + constantName;

        buffer.addLine("%1$s %2$s = \"%3$s\"", constantName, className, constantValue);
    }

    private void generateMemberFormalParameter(StructMember member) {
        GoTypeReference memberTypeReference = goNames.getRefTypeReference(member.getType());
        buffer.addImports(memberTypeReference.getImports());
        buffer.addLine(
            "%1$s %2$s `xml:\"%3$s,omitempty\"` ",
            goNames.getPublicMemberStyleName(member.getName()),
            memberTypeReference.getText(),
            goNames.getTagStyleName(member.getName())
        );
    }

    private void generateBuilderMemberMethods(StructType type, StructMember member) {
        // Get Type names
        GoClassName typeName = goNames.getTypeName(type);
        String typeClassName = typeName.getClassName();
        String typePrivateClassName = typeName.getPrivateClassName();
        String typePrivateMemberName = goNames.getPrivateMemberStyleName(type.getName());

        // Get member names
        GoTypeReference memberTypeReference = goNames.getTypeReference(member.getType());
        // Define method for TypeBuilder
        buffer.addLine("func (builder *%1$sBuilder) %2$s(%3$s %4$s) *%1$sBuilder {",
            typePrivateClassName, goNames.getPublicMethodStyleName(member.getName()),
            goNames.getParameterStyleName(member.getName()), memberTypeReference.getText());
        //      Check if has errors
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();
        //      Method Body
        String settedValue = goNames.getParameterStyleName(member.getName());
        if (GoTypes.isGoPrimitiveType(member.getType())) {
            settedValue = "&" + settedValue;
        }
        
        buffer.addLine(  "builder.%1$s.%2$s = %3$s",
            typePrivateMemberName,
            goNames.getPublicMemberStyleName(member.getName()),
            settedValue
            );
        buffer.addLine(  "return builder");
        buffer.addLine("}");
        buffer.addLine();
    }

}

