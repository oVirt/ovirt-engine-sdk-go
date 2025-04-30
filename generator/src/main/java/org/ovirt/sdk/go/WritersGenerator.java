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

import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.EnumType;
import org.ovirt.api.metamodel.concepts.Link;
import org.ovirt.api.metamodel.concepts.ListType;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.StructMember;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.Names;
import org.ovirt.api.metamodel.tool.SchemaNames;

/**
 * This class is responsible for generating the functions that render XML documents from instances of model types.
 */
public class WritersGenerator implements GoGenerator {
    // The directory were the output will be generated:
    protected File out;
    
    // Reference to object used to calculate names:
    @Inject
    private Names names;

    // Reference to the objects used to generate the code:
    @Inject private GoNames goNames;

    @Inject private GoTypes goTypes;

    // The buffer used to generate the code:
    private GoBuffer buffer;
    
    // Reference to the object used to calculate XML schema names:
    @Inject private SchemaNames schemaNames;

    @Inject private GoPackages goPackages;

    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Prepare the buffer:
        buffer = new GoBuffer();
        buffer.setPackageName(goPackages.getWritersPackageName());
        buffer.addImport("fmt");
        // Generate classes for each struct type:
        model.types()
            .filter(StructType.class::isInstance)
            .map(StructType.class::cast)
            .forEach(this::generateStructWriter);

        // Generate classes for each enum type:
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .forEach(this::generateEnumWriter);

        // Write the file:
        try {
            buffer.write(out);
        } catch (IOException exception) {
            throw new RuntimeException("Error writing writers", exception);
        }
    }

    private void generateStructWriter(StructType type) {
        // Generate methods to read one instance and a list of instances:
        generateStructWriteOne(type);
        generateStructWriteMany(type);
    }

    private void generateEnumWriter(EnumType type) {
        generateEnumWriteOne(type);
        generateEnumWriteMany(type);
    }

    private void generateStructWriteOne(StructType type) {
        GoClassName typeName = goTypes.getTypeName(type);

        // Generate the method:
        List<StructMember> allMembers = new ArrayList<>();
        allMembers.addAll(type.getAttributes());
        allMembers.addAll(type.getLinks());
        List<StructMember> asAttributes = allMembers.stream()
            .filter(x -> schemaNames.isRepresentedAsAttribute(x.getName()))
            .sorted()
            .collect(toList());
        List<StructMember> asElements = allMembers.stream()
            .filter(x -> !schemaNames.isRepresentedAsAttribute(x.getName()))
            .sorted()
            .collect(toList());

        buffer.addLine("func %1$s(writer *XMLWriter, object *%2$s, tag string) error {",
            goTypes.getXmlWriteOneFuncName(type).getSimpleName(), typeName.getSimpleName());
        // Check the object if is nil
        buffer.addLine("  if object == nil {");
        buffer.addLine("    return fmt.Errorf(\"input object pointer is nil\")");
        buffer.addLine("  }");
        // Generate the `tag` value
        buffer.addLine("  if tag == \"\" {");
        buffer.addLine("    tag = \"%1$s\"", goNames.getTagStyleName(type.getName()));
        buffer.addLine("  }");
        // Generate the attributes
        if (!asAttributes.isEmpty()) {
            buffer.addLine("  var attrs map[string]string");
            for (StructMember member : asAttributes) {
                this.generateStructWriteMemberAsAttribute(type, member);
            }
            buffer.addLine("  writer.WriteStart(\"\", tag, attrs)");
        } else {
            buffer.addLine("  writer.WriteStart(\"\", tag, nil)");
        }
        
        // Generate the XML elements of inner members
        if (!asElements.isEmpty()) {
            for (StructMember member : asElements) {
                this.generateStructReadMemberAsElement(type, member);
            }
        } else {

        }
        
        buffer.addLine("  writer.WriteEnd(tag)");
        buffer.addLine("  return nil");
        buffer.addLine("}");    // End of function
        buffer.addLine();
    }

    private void generateStructWriteMany(StructType type) {
        // Calculate the tag names:
        GoClassName typeName = goTypes.getTypeName(type);

        Name singularName = type.getName();
        Name pluralName = names.getPlural(singularName);
        String singularTag = goNames.getTagStyleName(singularName);
        String pluralTag = goNames.getTagStyleName(pluralName);

        GoClassName typeSliceName = goTypes.getTypeSliceName(type);
        buffer.addLine("func %1$s(writer *XMLWriter, structSlice *%2$s, plural, singular string) error {",
            goTypes.getXmlWriteManyFuncName(type).getSimpleName(), typeSliceName.getSimpleName());

        // Generate the plural and singular name
        buffer.addLine("  if plural == \"\" {");
        buffer.addLine("    plural = \"%1$s\"", pluralTag);
        buffer.addLine("  }");
        buffer.addLine("  if singular == \"\" {");
        buffer.addLine("    singular = \"%1$s\"", singularTag);
        buffer.addLine("  }");
        buffer.addLine("  writer.WriteStart(\"\", plural, nil)");
        buffer.addLine("  for _, o := range structSlice.Slice() {");
        buffer.addLine("    %1$s(writer, o, singular)",
            goTypes.getXmlWriteOneFuncName(type).getSimpleName());
        buffer.addLine("  }");
        buffer.addLine("  writer.WriteEnd(plural)");

        buffer.addLine("  return nil");
        buffer.addLine("}");    // End of function
        buffer.addLine();
    }

    private void generateStructWriteMemberAsAttribute(StructType structType, StructMember member) {
        Name memberName = member.getName();
        Type memberType = member.getType();
        String tag = goNames.getTagStyleName(memberName);
        buffer.addLine("  if r, ok := object.%1$s(); ok {", goTypes.getMemberGetterMethodName(memberName).getSimpleName());
        buffer.addLine("    if attrs == nil {");
        buffer.addLine("      attrs = make(map[string]string)");
        buffer.addLine("    }");
        if (memberType instanceof PrimitiveType) {
            Model model = memberType.getModel();
            if (memberType == model.getBooleanType()) {
                buffer.addLine("    attrs[\"%1$s\"] = writer.FormatBool(r)", tag);
            }
            else if (memberType == model.getIntegerType()) {
                buffer.addLine("    attrs[\"%1$s\"] = writer.FormatInt64(r)", tag);
            }
            else if (memberType == model.getDecimalType()) {
                buffer.addLine("    attrs[\"%1$s\"] = writer.FormatFloat64(r)", tag);
            }
            else if (memberType == model.getStringType()) {
                buffer.addLine("    attrs[\"%1$s\"] = r", tag);
            }
            else if (memberType == model.getDateType()) {
                buffer.addLine("    attrs[\"%1$s\"] = writer.FormatDate(r)", tag);
            }
        }
        else if (memberType instanceof EnumType) {
            buffer.addLine("    attrs[\"%1$s\"] = string(r)", tag);
        }
        buffer.addLine("  }");    // End of if

    }

    public void generateStructReadMemberAsElement(StructType structType, StructMember member) {
        Name memberName = member.getName();
        Type memberType = member.getType();
        String tag = goNames.getTagStyleName(memberName);
        
        buffer.addLine("  if r, ok := object.%1$s(); ok {", goTypes.getMemberGetterMethodName(memberName).getSimpleName());
        if (memberType instanceof PrimitiveType) {
            Model model = memberType.getModel();
            if (memberType == model.getBooleanType()) {
                buffer.addLine("    writer.WriteBool(\"%1$s\", r)", tag);
            }
            else if (memberType == model.getIntegerType()) {
                buffer.addLine("    writer.WriteInt64(\"%1$s\", r)", tag);
            }
            else if (memberType == model.getDecimalType()) {
                buffer.addLine("    writer.WriteFloat64(\"%1$s\", r)", tag);
            }
            else if (memberType == model.getStringType()) {
                buffer.addLine("    writer.WriteCharacter(\"%1$s\", r)", tag);
            }
            else if (memberType == model.getDateType()) {
                buffer.addLine("    writer.WriteDate(\"%1$s\", r)", tag);
            }
        }
        else if (memberType instanceof StructType || memberType instanceof EnumType) {
            buffer.addLine("%1$s(writer, r, \"%2$s\")",
                goTypes.getXmlWriteOneFuncName(memberType).getSimpleName(),
                tag
            );
        }
        else if (memberType instanceof ListType) {
            ListType listType = (ListType) memberType;
            Type elementType = listType.getElementType();
            if (elementType instanceof StructType || elementType instanceof EnumType) {
                String elementSingularTag;
                if(elementType instanceof StructType) {
                    elementSingularTag = goNames.getTagStyleName(elementType.getName());
                } else {
                    elementSingularTag = goNames.getTagStyleName(names.getSingular(member.getName()));
                }

                buffer.addLine("%1$s(writer, r, \"%2$s\", \"%3$s\")",
                    goTypes.getXmlWriteManyFuncName(elementType).getSimpleName(),
                    tag,
                    elementSingularTag
                );
            }
            else if (elementType instanceof PrimitiveType) {
                Model model = memberType.getModel();
                if (elementType == model.getBooleanType()) {
                    buffer.addLine("writer.WriteBools(\"%1$s\", r);", tag);
                }
                else if (elementType == model.getIntegerType()) {
                    buffer.addLine("writer.WriteInt64s(\"%1$s\", r);", tag);
                }
                else if (elementType == model.getDecimalType()) {
                    buffer.addLine("writer.WriteFloat64s(\"%1$s\", r);", tag);
                }
                else if (elementType == model.getStringType()) {
                    buffer.addLine("writer.WriteCharacters(\"%1$s\", r);", tag);
                }
                else if (elementType == model.getDateType()) {
                    buffer.addLine("writer.writeDates(\"%1$s\", r);", tag);
                }
            }
        }

        buffer.addLine("  }");  // End of if

    }

    private void generateEnumWriteOne(EnumType type) {
        GoClassName typeName = goTypes.getTypeName(type);
        String tag = goNames.getTagStyleName(type.getName());
        buffer.addLine("func %1$s(writer *XMLWriter, enum %2$s, tag string) {",
            goTypes.getXmlWriteOneFuncName(type).getSimpleName(), typeName.getSimpleName());
        // Generate the function body
        //      Generate the `tag` value
        buffer.addLine("  if tag == \"\" {");
        buffer.addLine("    tag = \"%1$s\"", tag);
        buffer.addLine("  }");
        buffer.addLine("  writer.WriteCharacter(tag, string(enum))");
        buffer.addLine("}");    // End of function
        buffer.addLine();
    }

    private void generateEnumWriteMany(EnumType type) {
        GoClassName typeName = goTypes.getTypeName(type);

        Name singularName = type.getName();
        Name pluralName = names.getPlural(singularName);
        String singularTag = goNames.getTagStyleName(singularName);
        String pluralTag = goNames.getTagStyleName(pluralName);

        buffer.addLine("func %1$s(writer *XMLWriter, enums []%2$s, plural, singular string) error {",
            goTypes.getXmlWriteManyFuncName(type).getSimpleName(), typeName.getSimpleName());

        // Generate the plural and singular name
        buffer.addLine("  if plural == \"\" {");
        buffer.addLine("    plural = \"%1$s\"", pluralTag);
        buffer.addLine("  }");
        buffer.addLine("  if singular == \"\" {");
        buffer.addLine("    singular = \"%1$s\"", singularTag);
        buffer.addLine("  }");

        buffer.addLine("  writer.WriteStart(\"\", plural, nil)");
        buffer.addLine("  for _, e := range enums {");
        buffer.addLine("    writer.WriteCharacter(singular, string(e))");
        buffer.addLine("  }");
        buffer.addLine("  writer.WriteEnd(plural)");
        buffer.addLine("  return nil");
        buffer.addLine("}");    // End of function
        buffer.addLine();
    }

    
}