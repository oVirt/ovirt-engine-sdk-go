/*
The oVirt Project - oVirt Engine Go SDK

Copyright (c) oVirt Authors

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
import org.ovirt.api.metamodel.concepts.ListType;
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
    @Inject private Names names;
    @Inject private GoNames goNames;
    @Inject private GoPackages goPackages;

    // The buffer used to generate the code:
    private GoBuffer buffer;

    // Reference to the object used to calculate Go types:
    @Inject private GoTypes goTypes;


    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Prepare the buffer:
        buffer = new GoBuffer();
        buffer.setPackageName(goPackages.getTypesPackageName());

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
            .collect(toList());

        // Generate the type:
        structs.forEach(this::generateStruct);

        // Generate the slices of the type
        structs.forEach(this::generateStructSlice);

        // Generate the builder:
        structs.forEach(this::generateStructBuilder);

        // Enum types don't need any special order, so we sort them only by name:
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .forEach(this::generateEnum);
    }

    private void addDocTag(String name, Type type) {
        final String doc = type.getDoc();
        if (doc != null && !doc.isEmpty()) {
            buffer.addLine(
                "// %1$s %2$s",
                name,
                doc.replace("\n", "\n// ")
            );
        }
    }

    private void generateStruct(StructType type) {
        // Begin class:
        GoClassName typeName = goTypes.getTypeName(type);

        addDocTag(typeName.getSimpleName(), type);
        // Define Struct
        buffer.addLine("type %1$s struct {", typeName.getSimpleName());
        // Ignore Base-class mixin, fill in all
        buffer.addLine("Struct");

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

        // Generate the member methods of type
        List<StructMember> members = allMembers.stream().sorted().collect(toCollection(ArrayList::new));
        for (StructMember m : members) {
            this.generateMemberMethods(type, m);
        }

        buffer.addLine();
    }

    private void generateStructSlice(StructType type) {
        GoClassName typeName = goTypes.getTypeName(type);
        GoClassName typeSliceName = goTypes.getTypeSliceName(type);
        // Define Struct Slice
        buffer.addLine("type %1$s struct {", typeSliceName.getSimpleName());
        buffer.addLine("  href *string");
        buffer.addLine("  slice []*%1$s", typeName.getSimpleName());
        buffer.addLine("}");
        buffer.addLine();

        // Define the methods of Struct Slice
        buffer.addLine("func (op *%1$s) Href() (string, bool) {",
            typeSliceName.getSimpleName());
        buffer.addLine(" if op.href == nil {");
        buffer.addLine("   return \"\", false");
        buffer.addLine(" }");
        buffer.addLine(" return *op.href, true");
        buffer.addLine("}");
        buffer.addLine();

        buffer.addLine("func (op *%1$s) SetHref(href string) {",
            typeSliceName.getSimpleName());
        buffer.addLine("  op.href = &href");
        buffer.addLine("}");
        buffer.addLine();

        buffer.addLine("func (op *%1$s) Slice() []%2$s {",
            typeSliceName.getSimpleName(), goTypes.getTypeReferenceAsReturnvalue(type));
        buffer.addLine("  return op.slice");
        buffer.addLine("}");
        buffer.addLine();

        buffer.addLine("func (op *%1$s) SetSlice(slice []%2$s) {",
            typeSliceName.getSimpleName(), goTypes.getTypeReferenceAsReturnvalue(type));
        buffer.addLine("  op.slice = slice");
        buffer.addLine("}");
        buffer.addLine();

    }

    private void generateMemberMethods(StructType structType, StructMember member) {
        // Begin class:
        GoClassName structTypeName = goTypes.getTypeName(structType);
        Type type = member.getType();
        // Generate the setter method
        addDocTag(goTypes.getMemberSetterMethodName(member.getName()).getSimpleName(), type);
        buffer.addLine("func (p *%1$s) %2$s(attr %3$s) {",
            structTypeName.getSimpleName(),
            goTypes.getMemberSetterMethodName(member.getName()),
            goTypes.getTypeReferenceAsParameter(member.getType()));
        //      Generate the setter method body
        if (goTypes.isGoPrimitiveType(type) || type instanceof EnumType) {
            buffer.addLine(" p.%1$s = &attr", goNames.getUnexportableMemberStyleName(member.getName()));
        }
        else {
            buffer.addLine(" p.%1$s = attr", goNames.getUnexportableMemberStyleName(member.getName()));
        }
        buffer.addLine("}");    // End of setter method
        buffer.addLine();

        addDocTag(goTypes.getMemberGetterMethodName(member.getName()).getSimpleName(), type);
        // Generate the getter method
        buffer.addLine("func (p *%1$s) %2$s() (%3$s, bool) {",
            structTypeName.getSimpleName(),
            goTypes.getMemberGetterMethodName(member.getName()),
            goTypes.getTypeReferenceAsReturnvalue(member.getType())
        );
        //      Generate the getter method body
        buffer.addLine(" if p.%1$s != nil {", goNames.getUnexportableMemberStyleName(member.getName()));
        if (goTypes.isGoPrimitiveType(type) || type instanceof EnumType) {
            buffer.addLine("  return *p.%1$s, true", goNames.getUnexportableMemberStyleName(member.getName()));
            buffer.addLine(" }");
            buffer.addLine(" var zero %1$s", goTypes.getTypeReferenceAsVariable(member.getType()));;
            buffer.addLine(" return zero, false");
        }
        else {
            buffer.addLine("  return p.%1$s, true", goNames.getUnexportableMemberStyleName(member.getName()));
            buffer.addLine(" }");
            buffer.addLine(" return nil, false");
        }
        buffer.addLine("}");    // End of getter method
        buffer.addLine();

        // Generate the MUST getter method
        addDocTag(goTypes.getMemberMustGetterMethodName(member.getName()).getSimpleName(), type);
        buffer.addLine("func (p *%1$s) %2$s() %3$s {",
            structTypeName.getSimpleName(),
            goTypes.getMemberMustGetterMethodName(member.getName()),
            goTypes.getTypeReferenceAsReturnvalue(member.getType())
        );
        buffer.addLine(" if p.%1$s == nil {", goNames.getUnexportableMemberStyleName(member.getName()));
        buffer.addLine("  panic(\"the %1$s must not be nil, please use %2$s() function instead\")",
            goNames.getUnexportableMemberStyleName(member.getName()),
            goTypes.getMemberGetterMethodName(member.getName()));
        buffer.addLine(" }");
        if (goTypes.isGoPrimitiveType(type) || type instanceof EnumType) {
            buffer.addLine(" return *p.%1$s", goNames.getUnexportableMemberStyleName(member.getName()));
        } else {
            buffer.addLine(" return p.%1$s", goNames.getUnexportableMemberStyleName(member.getName()));
        }
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateStructBuilder(StructType type) {
        // Get Type names
        GoClassName typeName = goTypes.getTypeName(type);
        String typeClassName = typeName.getSimpleName();
        String typePrivateMemberName = goNames.getUnexportableMemberStyleName(type.getName());
        // Get struct members
        Set<StructMember> allMembers = Stream.concat(type.attributes(), type.links())
            .collect(toSet());
        Set<StructMember> declaredMembers = Stream.concat(type.declaredAttributes(), type.declaredLinks())
            .collect(toSet());
        allMembers.addAll(declaredMembers);

        // Begin class:
        buffer.addLine("type %1$s struct {", goTypes.getBuilderName(type).getSimpleName());
        //      Add properties of TypeBuilder
        buffer.addLine(  "%1$s *%2$s", typePrivateMemberName, typeName.getSimpleName());
        buffer.addLine(  "err error");
        // End class:
        buffer.addLine("}");
        buffer.addLine();

        // Define NewStructBuilder function
        buffer.addLine("func %1$s() *%2$s {",
            goTypes.getNewBuilderFuncName(type).getSimpleName(),
            goTypes.getBuilderName(type).getSimpleName());
        buffer.addLine("return &%1$s{%2$s: &%3$s{}, err: nil}",
            goTypes.getBuilderName(type).getSimpleName(),
            typePrivateMemberName,
            typeClassName
            );
        buffer.addLine("}");
        buffer.addLine();

        // Construct TypeBuilder methods for member-settings
        List<StructMember> members = allMembers.stream().sorted().collect(toCollection(ArrayList::new));
        for (StructMember member : members) {
            this.generateBuilderMethods(type, member);
        }
        // Generate Href setting method
        buffer.addLine("func (builder *%1$s) Href(href string) *%1$s {", goTypes.getBuilderName(type).getSimpleName());
        //      Check if has errors
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();
        buffer.addLine(  "builder.%1$s.SetHref(href)", typePrivateMemberName);
        buffer.addLine(  "return builder");
        buffer.addLine("}");
        buffer.addLine();
        // Generate Build method
        buffer.addLine("func (builder *%1$s) Build() (*%2$s, error) {", goTypes.getBuilderName(type).getSimpleName(), typeClassName);
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return nil, builder.err");
        buffer.addLine(  "}");
        buffer.addLine(  "return builder.%1$s, nil", typePrivateMemberName);
        buffer.addLine("}");
        buffer.addLine();
        // Generate MustBuild method
        buffer.addImport("fmt");
        buffer.addLine("func (builder *%1$s) MustBuild() *%2$s {", goTypes.getBuilderName(type).getSimpleName(), typeClassName);
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "panic(fmt.Sprintf(\"Failed to build %1$s instance, reason: %%v\", builder.err))", typeClassName);
        buffer.addLine(  "}");
        buffer.addLine(  "return builder.%1$s", typePrivateMemberName);
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateEnum(EnumType type) {
        // Begin class:
        GoClassName typeName = goTypes.getTypeName(type);

        // Type declaration
        buffer.addLine("type %1$s string", typeName.getSimpleName());

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
        String className = goTypes.getTypeName(value.getDeclaringType()).getSimpleName();
        String constantValue = names.getLowerJoined(name, "_");

        // To avoid constant-name conflict, eg: VMSTATUS_DOWN
        constantName = className.toUpperCase() + "_" + constantName;

        buffer.addLine("%1$s %2$s = \"%3$s\"", constantName, className, constantValue);
    }

    private void generateMemberFormalParameter(StructMember member) {
        Type memberType = member.getType();
        GoTypeReference memberTypeReference = goTypes.getTypeReferenceAsAttribute(memberType);
        buffer.addImports(memberTypeReference.getImports());

        buffer.addLine(
            "%1$s %2$s",
            goNames.getUnexportableMemberStyleName(member.getName()),
            memberTypeReference.getText()
        );
    }

    private void generateBuilderMethods(StructType type, StructMember member) {
        // Get Type names
        String typePrivateMemberName = goNames.getUnexportableMemberStyleName(type.getName());

        // Get member names
        GoTypeReference memberTypeReference = goTypes.getTypeReferenceAsParameter(member.getType());
        // Define method for TypeBuilder
        buffer.addLine("func (builder *%1$s) %2$s(attr %3$s) *%1$s {",
            goTypes.getBuilderName(type).getSimpleName(),
            goNames.getExportableMethodStyleName(member.getName()),
            memberTypeReference.getText());
        //      Check if has errors
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();
        //      Method Body
        String settedValue = goNames.getParameterStyleName(member.getName());
        if (goTypes.isGoPrimitiveType(member.getType()) || 
                member.getType() instanceof EnumType) {
            settedValue = "&" + settedValue;
        }
        
        buffer.addLine(  "builder.%1$s.%2$s(attr)",
            typePrivateMemberName,
            goTypes.getMemberSetterMethodName(member.getName())
            );
        buffer.addLine(  "return builder");
        buffer.addLine("}");
        buffer.addLine();
        // If is TypeSlice (ListType), add the varargs pattern of method
        if (member.getType() instanceof ListType) {
            this.generateBuilderMethodsTakingVararg(type, member);
        } else if (member.getType() instanceof StructType) {
            this.generateBuilderMethodsTakingBuilderArg(type, member);
        }
    }

    private void generateBuilderMethodsTakingVararg(StructType type, StructMember member) {
        // Get Type names
        String typePrivateMemberName = goNames.getUnexportableMemberStyleName(type.getName());
        
        ListType memberType = (ListType) member.getType();
        Type elementType = memberType.getElementType();

        buffer.addLine("func (builder *%1$s) %2$sOfAny(anys ...%3$s) *%1$s {",
            goTypes.getBuilderName(type).getSimpleName(),
            goNames.getExportableMethodStyleName(member.getName()),
            goTypes.getTypeReferenceAsParameter(elementType).getText());

        //      Check if has errors
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();

        // Do appending after TypeSlice initialized
        if (goTypes.isGoPrimitiveType(elementType) || elementType instanceof EnumType) {
            buffer.addLine("  builder.%1$s.%2$s = append(builder.%1$s.%2$s, anys...)",
                typePrivateMemberName,
                goNames.getUnexportableMemberStyleName(member.getName()));
        } else {
            buffer.addLine("  if builder.%1$s.%2$s == nil {",
                typePrivateMemberName,
                goNames.getUnexportableMemberStyleName(member.getName()));
            buffer.addLine("    builder.%1$s.%2$s = new(%3$s)",
                typePrivateMemberName,
                goNames.getUnexportableMemberStyleName(member.getName()),
                goTypes.getTypeSliceName(elementType).getSimpleName());
            buffer.addLine("  }");
            buffer.addLine("  builder.%1$s.%2$s.slice = append(builder.%1$s.%2$s.slice, anys...)",
                typePrivateMemberName,
                goNames.getUnexportableMemberStyleName(member.getName()));
        }
        buffer.addLine("return builder");
        buffer.addLine("}");
        buffer.addLine();
        
        if (elementType instanceof StructType) {
            generateBuilderMethodsTakingBuilderVararg(type, member);
        }
    }

    private void generateBuilderMethodsTakingBuilderArg(StructType type, StructMember member) {
        // Define method for TypeBuilder
        buffer.addLine("func (builder *%1$s) %2$sBuilder(attrBuilder *%3$s) *%1$s {",
            goTypes.getBuilderName(type).getSimpleName(),
            goNames.getExportableMethodStyleName(member.getName()),
            goTypes.getBuilderName(member.getType()).getSimpleName());
        //      Check if has errors
        buffer.addLine(  "if builder.err != nil {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();

        buffer.addLine(  "if attrBuilder.err != nil {");
        buffer.addLine(  "  builder.err = attrBuilder.err");
        buffer.addLine(  "  return builder");
        buffer.addLine(  "}");
        
        //      Build out the attr struct
        buffer.addLine(" attr, err := attrBuilder.Build()");
        buffer.addLine(" if err != nil {");
        buffer.addLine("   builder.err = err");
        buffer.addLine("   return builder");
        buffer.addLine(  "}");

        buffer.addLine(" return builder.%1$s(attr)", 
            goNames.getExportableMethodStyleName(member.getName()));
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateBuilderMethodsTakingBuilderVararg(StructType type, StructMember member) {
        // Get Type names
        ListType memberType = (ListType) member.getType();
        Type elementType = memberType.getElementType();

        buffer.addLine("func (builder *%1$s) %2$sBuilderOfAny(anyBuilders ...%3$s) *%1$s {",
            goTypes.getBuilderName(type).getSimpleName(),
            goNames.getExportableMethodStyleName(member.getName()),
            goTypes.getBuilderName(elementType).getSimpleName());

        //      Check if has errors
        buffer.addLine(  "if builder.err != nil || len(anyBuilders) == 0 {");
        buffer.addLine(    "return builder");
        buffer.addLine(  "}");
        buffer.addLine();

        buffer.addLine("  for _, b := range anyBuilders {");
        buffer.addLine("      if b.err != nil {");
        buffer.addLine("          builder.err = b.err");
        buffer.addLine("          return builder");
        buffer.addLine("      }");
        buffer.addLine("      attr, err := b.Build()");
        buffer.addLine("      if err != nil {");
        buffer.addLine("          builder.err = b.err");
        buffer.addLine("          return builder");
        buffer.addLine("      }");
        buffer.addLine("      builder.%1$sOfAny(attr)", goNames.getExportableMethodStyleName(member.getName()));
        buffer.addLine("  }");
        buffer.addLine("  return builder");
        buffer.addLine("}");
        buffer.addLine();
    }

}

