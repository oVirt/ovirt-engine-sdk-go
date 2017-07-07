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
import static java.util.stream.Collectors.toSet;

import java.io.File;
import java.io.IOException;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Set;
import java.util.HashSet;
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
        Deque<StructType> pending = model.types()
            .filter(StructType.class::isInstance)
            .map(StructType.class::cast)
            .sorted()
            .collect(toCollection(ArrayDeque::new));
        Deque<StructType> structs = new ArrayDeque<>(pending.size());
        while (!pending.isEmpty()) {
            StructType current = pending.removeFirst();
            StructType base = (StructType) current.getBase();
            if (base == null || structs.contains(base)) {
                structs.addLast(current);
            }
            else {
                pending.addLast(current);
            }
        }
        structs.stream()
            .forEach(this::generateStruct);

        // Enum types don't need any special order, so we sort them only by name:
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .sorted()
            .forEach(this::generateEnum);
        
        // Customize type generation
        //      Add  Error method for type Fault
        buffer.addLine();
        buffer.addLine("func (fault *Fault) Error() string {");
        buffer.startBlock();
        buffer.addLine("return fmt.Sprintf(\"Error details is %%s, reason is %%s\", fault.Detail, fault.Reason)");
        buffer.endBlock();
        buffer.addLine("}");
    }

    private void generateStruct(StructType type) {
        // Begin class:
        GoClassName typeName = goNames.getTypeName(type);
        Type base = type.getBase();
        // Define []Struct
        buffer.addLine("type %1$ss struct {", typeName.getClassName());
        buffer.startBlock();
        //  Add xml.Name
        buffer.addLine("XMLName xml.Name `xml:\"%1$ss\"`", typeName.getClassName().toLowerCase());
        buffer.addLine("%1$ss []%1$s `xml:\"%2$s,omitempty\"`", typeName.getClassName(), goNames.getTagStyleName(type.getName()));
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
        // Define Struct
        buffer.addLine("type %1$s struct {", typeName.getClassName());
        buffer.startBlock();
        // Ignore Base-class mixin, fill in all 
        buffer.addLine("OvStruct");

        // Constructor with a named parameter for each attribute and link:
        Set<StructMember> allMembers = Stream.concat(type.attributes(), type.links())
            .collect(toSet());
        Set<StructMember> declaredMembers = Stream.concat(type.declaredAttributes(), type.declaredLinks())
            .collect(toSet());
        allMembers.addAll(declaredMembers);
        allMembers.stream().sorted().forEach(this::generateMemberFormalParameter);

        buffer.endBlock();
        buffer.addLine();

        // End class:
        buffer.endBlock();
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateEnum(EnumType type) {
        // Begin class:
        GoClassName typeName = goNames.getTypeName(type);

        // Type declaration
        buffer.addLine("type %1$s string", typeName.getClassName());

        // Type definition
        buffer.addLine("const (");
        buffer.startBlock();
        type.values().sorted().forEach(this::generateEnumValue);
        buffer.endBlock();
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
        GoTypeReference goTypeReference = goNames.getTypeReferenceAsStructMember(member.getType());
        buffer.addImports(goTypeReference.getImports());
        buffer.addLine(
            "%1$s    %2$s   `xml:\"%3$s,omitempty\"` ",
            goNames.getMemberStyleName(member.getName()),
            goTypeReference.getText(),
            goNames.getTagStyleName(member.getName())
        );
    }

}

