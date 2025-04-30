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
import org.ovirt.api.metamodel.tool.SchemaNames;

/**
 * This class is responsible for generating the functions that create instances of model types from XML documents.
 */
public class ReadersGenerator implements GoGenerator {
    // The directory were the output will be generated:
    protected File out;

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
        buffer.setPackageName(goPackages.getReadersPackageName());

        // Generate classes for each struct type:
        model.types()
            .filter(StructType.class::isInstance)
            .map(StructType.class::cast)
            .forEach(this::generateStructReader);

        // Generate classes for each enum type:
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .forEach(this::generateEnumReader);

        // Write the file:
        try {
            buffer.write(out);
        } catch (IOException exception) {
            throw new RuntimeException("Error writing readers", exception);
        }
    }

    private void generateStructReader(StructType type) {
        // Generate methods to read one instance and a list of instances:
        buffer.addImport("encoding/xml");
        buffer.addImport("io");
        generateStructReadOne(type);
        generateStructReadMany(type);
        generateProcessLink(type);
    }

    private void generateStructReadOne(StructType type) {
        GoClassName typeName = goTypes.getTypeName(type);

        // Generate the tag name
        Name singularName = type.getName();
        String singularTag = goNames.getTagStyleName(singularName);

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

        GoClassName typeSliceName = goTypes.getTypeSliceName(type);

        buffer.addLine("func %1$s(reader *XMLReader, start *xml.StartElement, expectedTag string) (*%2$s, error) {",
            goTypes.getXmlReadOneFuncName(type).getSimpleName(), typeName.getSimpleName());
        // Generate the function body
        buffer.addLine(" builder := %1$s()", goTypes.getNewBuilderFuncName(type).getSimpleName());

        //      Generate `find start element`
        buffer.addLine("  if start == nil {");
        buffer.addLine("    st, err := reader.FindStartElement()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        return nil, nil");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    start = st");
        buffer.addLine("  }");

        buffer.addLine(" if expectedTag == \"\" {");
        buffer.addLine("   expectedTag = \"%1$s\"", singularTag);
        buffer.addLine(" }");

        buffer.addLine("  if start.Name.Local != expectedTag {");
        buffer.addLine("    return nil, XMLTagNotMatchError{start.Name.Local, expectedTag}");
        buffer.addLine("  }");

        //      Process the attributes
        if (!asAttributes.isEmpty()) {
            buffer.addLine("  // Process the attributes");
	        buffer.addLine("  for _, attr := range start.Attr {");
	        buffer.addLine("  	name := attr.Name.Local");
	        buffer.addLine("  	value := attr.Value");
            buffer.addLine("  	switch name {");
            for (StructMember member : asAttributes) {
                this.generateStructReadMemberFromAttribute(type, member);
            }
            buffer.addLine("    case \"href\":");
            buffer.addLine("      builder.Href(value)");
	        buffer.addLine("  	}");
	        buffer.addLine("  }");
        }

        // Define links slice
        buffer.addLine("  var links []Link");
        //      Process the inner elements:
        if (!asElements.isEmpty()) {
            buffer.addLine("  depth := 1");
            buffer.addLine("  for depth >0 {");
            buffer.addLine("    t, err := reader.Next()");
            buffer.addLine("    if err != nil {");
            buffer.addLine("      if err == io.EOF {");
            buffer.addLine("        break");
            buffer.addLine("      }");
            buffer.addLine("      return nil, err");
            buffer.addLine("    }");
            buffer.addLine("    t = xml.CopyToken(t)");
            buffer.addLine("    switch t := t.(type) {");
            buffer.addLine("    case xml.StartElement:");
            // Generate attributes switch-case
            buffer.addLine("      switch t.Name.Local {");
            for (StructMember member : asElements) {
                this.generateStructReadMemberFromElement(type, member);
            }
            //      It's <link>
            buffer.addLine("      case \"link\":");
            buffer.addLine("        var rel, href string");
	        buffer.addLine("        for _, attr := range t.Attr {");
	        buffer.addLine("  	      name := attr.Name.Local");
	        buffer.addLine("  	      value := attr.Value");
            buffer.addLine("  	      switch name {");
            buffer.addLine("            case \"href\":");
            buffer.addLine("              href = value");
            buffer.addLine("            case \"rel\":");
            buffer.addLine("              rel = value");
	        buffer.addLine("  	      }");
            buffer.addLine("        }");
            buffer.addLine("        if rel != \"\" && href != \"\" {");
            buffer.addLine("          links = append(links, Link{&href, &rel})");
            buffer.addLine("        }");
            buffer.addLine("        // <link> just has attributes, so must skip manually");
            buffer.addLine("        reader.Skip()");
            
            buffer.addLine("      default:");
            buffer.addLine("        reader.Skip()");
            buffer.addLine("      }");
            buffer.addLine("    case xml.EndElement:");
            buffer.addLine("      depth--");
            buffer.addLine("    }");
            buffer.addLine("  }");
        } else {
            buffer.addLine("  reader.Skip()");
        }

        buffer.addLine("  one, err := builder.Build()");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    return nil, err");
        buffer.addLine("  }");
        // Generate processing the links
        buffer.addLine("  for _, link := range links {");
        buffer.addLine("    switch *link.rel {");
        List<Link> links = type.links()
            .sorted()
            .filter(link -> link.getType() instanceof ListType)
            .collect(toList());
        links.forEach(
            link -> {
                String field = goNames.getUnexportableMemberStyleName(link.getName());
                Type elementType = ((ListType) link.getType()).getElementType();
                String rel = link.getName().words().map(String::toLowerCase).collect(joining());
                buffer.addLine("      case \"%1$s\":", rel);
                buffer.addLine("        if one.%1$s == nil {", field);
                buffer.addLine("          one.%1$s = new(%2$s)", field, goTypes.getTypeSliceName(elementType));
                buffer.addLine("        }");
                buffer.addLine("        one.%1$s.href = link.href", field);
            }
        );
        buffer.addLine("    }  // end of switch");
        buffer.addLine("  }  // end of for-links");

        buffer.addLine("  return one, nil");
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateStructReadMany(StructType type) {
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
        GoClassName typeSliceName = goTypes.getTypeSliceName(type);
        buffer.addLine("func %1$s(reader *XMLReader, start *xml.StartElement) (*%2$s, error) {",
            goTypes.getXmlReadManyFuncName(type).getSimpleName(), typeSliceName.getSimpleName());

        // Generate the function body
        //      Generate `find start element`
        buffer.addLine("  if start == nil {");
        buffer.addLine("    st, err := reader.FindStartElement()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        return nil, nil");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    start = st");
        buffer.addLine("  }");

        //      Generate slice of type definition
        buffer.addLine("  var result %1$s", typeSliceName.getSimpleName());
        //      Process the inner elements:
        buffer.addLine("  depth := 1");
        buffer.addLine("  for depth >0 {");
        buffer.addLine("    t, err := reader.Next()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        break");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    t = xml.CopyToken(t)");
        buffer.addLine("    switch t := t.(type) {");
        buffer.addLine("    case xml.StartElement:");
        String singularTag = goNames.getTagStyleName(type.getName());
        buffer.addLine("      switch t.Name.Local {");
        buffer.addLine("      case \"%1$s\":", singularTag);
        buffer.addLine("        one, err := %1$s(reader, &t, \"%2$s\")", goTypes.getXmlReadOneFuncName(type).getSimpleName(), singularTag);
        buffer.addLine("        if err != nil {");
        buffer.addLine("          return nil, err");
        buffer.addLine("        }");
        buffer.addLine("        if one != nil {");
        buffer.addLine("          result.slice = append(result.slice, one)");
        buffer.addLine("        }");
        buffer.addLine("      default:");
        buffer.addLine("        reader.Skip()");
        buffer.addLine("      }");
        buffer.addLine("	case xml.EndElement:");
        buffer.addLine("      depth--");
        buffer.addLine("    }");
        // End of for
        buffer.addLine("  }");

        buffer.addLine("  return &result, nil");
        // End of function
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateStructReadMemberFromAttribute(StructType structType, StructMember member) {
        Name memberName = member.getName();
        Type memberType = member.getType();
        String publicMethodName = goNames.getExportableMethodStyleName(memberName);
        String tag = goNames.getTagStyleName(memberName);
        buffer.addLine("      case \"%1$s\":", tag);
        if (memberType instanceof PrimitiveType) {
            Model model = memberType.getModel();
            if (memberType == model.getBooleanType()) {
                buffer.addImport("strconv");
                buffer.addLine("        v, err := strconv.ParseBool(value)");
                buffer.addLine("        if err != nil {");
                buffer.addLine("          return nil, err");
                buffer.addLine("        }");
                buffer.addLine("        builder.%1$s(v)", publicMethodName);
            }
            else if (memberType == model.getIntegerType()) {
                buffer.addImport("strconv");
                buffer.addLine("        v, err := strconv.ParseInt(value, 10, 64)");
                buffer.addLine("        if err != nil {");
                buffer.addLine("          return nil, err");
                buffer.addLine("        }");
                buffer.addLine("        builder.%1$s(v)", publicMethodName);
            }
            else if (memberType == model.getDecimalType()) {
                buffer.addImport("strconv");
                buffer.addLine("        v, err := strconv.ParseBool(value)", publicMethodName);                
                buffer.addLine("        if err != nil {");
                buffer.addLine("          return nil, err");
                buffer.addLine("        }");
                buffer.addLine("        builder.%1$s(v)", publicMethodName);
            }
            else if (memberType == model.getStringType()) {
                buffer.addLine("        builder.%1$s(value)", publicMethodName);
            }
            else if (memberType == model.getDateType()) {
                buffer.addImport("time");
                buffer.addLine("        builder.%1$s(time.Parse(time.RFC3339Nano, value))", publicMethodName);
                buffer.addLine("        if err != nil {");
                buffer.addLine("          return nil, err");
                buffer.addLine("        }");
                buffer.addLine("        builder.%1$s(v)", publicMethodName);
            }
        }
        else if (memberType instanceof EnumType) {
            buffer.addLine("        builder.%1$s(%2$s(value))", publicMethodName, goTypes.getTypeName(memberType).getSimpleName());
        }
    }

    private void generateStructReadMemberFromElement(StructType structType, StructMember member) {
        Name memberName = member.getName();
        Type memberType = member.getType();
        String publicMethodName = goNames.getExportableMethodStyleName(memberName);
        String tag = goNames.getTagStyleName(memberName);

        buffer.addLine("      case \"%1$s\":", tag);
        if (memberType instanceof PrimitiveType) {
            Model model = memberType.getModel();
            if (memberType == model.getBooleanType()) {
                buffer.addLine("        v, err := reader.ReadBool(&t)");
            } else if (memberType == model.getIntegerType()) {
                buffer.addLine("        v, err := reader.ReadInt64(&t)");
            } else if (memberType == model.getDecimalType()) {
                buffer.addLine("        v, err := reader.ReadFloat64(&t)");
            } else if (memberType == model.getStringType()) {
                buffer.addLine("        v, err := reader.ReadString(&t)");
            } else if (memberType == model.getDateType()) {
                buffer.addLine("        v, err := reader.ReadTime(&t)");
            } else {
                buffer.addLine("        reader.Skip()");
            }
        }
        else if (memberType instanceof StructType) {
            GoFuncName readOneFuncName = goTypes.getXmlReadOneFuncName(memberType);
            buffer.addLine("        v, err := %1$s(reader, &t, \"%2$s\")", readOneFuncName.getSimpleName(), tag);
        }
        else if (memberType instanceof EnumType) {
            GoFuncName readOneFuncName = goTypes.getXmlReadOneFuncName(memberType);
            buffer.addLine("        vp, err := %1$s(reader, &t)", readOneFuncName.getSimpleName());
            buffer.addLine("        v := *vp");
        }
        else if (memberType instanceof ListType) {
            ListType listType = (ListType) memberType;
            Type elementType = listType.getElementType();
            GoFuncName readManyFuncName = goTypes.getXmlReadManyFuncName(elementType);
            if (elementType instanceof StructType || elementType instanceof EnumType) {
                buffer.addLine("        v, err := %1$s(reader, &t)", readManyFuncName.getSimpleName());
            }
            else if (elementType instanceof PrimitiveType) {
                Model model = memberType.getModel();
                if (elementType == model.getBooleanType()) {
                    buffer.addLine("        v, err := reader.ReadBools(&t)");
                }
                else if (elementType == model.getIntegerType()) {
                    buffer.addLine("        v, err := reader.ReadInt64s(&t)");
                }
                else if (elementType == model.getDecimalType()) {
                    buffer.addLine("        v, err := reader.ReadFloat64s(&t)");
                }
                else if (elementType == model.getStringType()) {
                    buffer.addLine("        v, err := reader.ReadStrings(&t)");
                }
                else if (elementType == model.getDateType()) {
                    buffer.addLine("        v, err := reader.ReadTimes(&t)");
                }
                else {
                    buffer.addLine("        reader.Skip()");
                }
            }
        } else {
            buffer.addLine("        reader.Skip()");
        }
        
        buffer.addLine("        if err != nil {");
        buffer.addLine("          return nil, err");
        buffer.addLine("        }");
        buffer.addLine("        builder.%1$s(v)", publicMethodName);
    }

    private void generateProcessLink(StructType type) {

    }

    private void generateEnumReader(EnumType type) {
        // Generate methods to read one instance and a list of instances:
        generateEnumReadOne(type);
        generateEnumReadMany(type);
    }

    private void generateEnumReadOne(EnumType type) {
        GoClassName typeName = goTypes.getTypeName(type);

        buffer.addLine("func %1$s(reader *XMLReader, start *xml.StartElement) (*%2$s, error) {",
            goTypes.getXmlReadOneFuncName(type).getSimpleName(), typeName.getSimpleName());
        // Generate the function body
        //      Generate `find start element`
        buffer.addLine("  if start == nil {");
        buffer.addLine("    st, err := reader.FindStartElement()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        return nil, nil");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    start = st");
        buffer.addLine("  }");
        
        buffer.addLine("  s, err := reader.ReadString(start)");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    return nil, err");
        buffer.addLine("  }");
        buffer.addLine("  result := new(%1$s)", typeName.getSimpleName());
        buffer.addLine("  *result = %1$s(s)", typeName.getSimpleName());
        buffer.addLine("  return result, nil", typeName.getSimpleName());

        // End of function
        buffer.addLine("}");;
        buffer.addLine();
    }

    private void generateEnumReadMany(EnumType type) {
        GoClassName typeName = goTypes.getTypeName(type);

        buffer.addLine("func %1$s(reader *XMLReader, start *xml.StartElement) ([]%2$s, error) {",
            goTypes.getXmlReadManyFuncName(type), typeName.getSimpleName());
        // Generate the function body
        //      Generate `find start element`
        buffer.addLine("  if start == nil {");
        buffer.addLine("    st, err := reader.FindStartElement()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        return nil, nil");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    start = st");
        buffer.addLine("  }");
        
        //      Generate slice of type definition
        buffer.addLine("  var results []%1$s", typeName.getSimpleName());
        
        //      Process the inner elements:
        buffer.addLine("  depth := 1");
        buffer.addLine("  for depth >0 {");
        buffer.addLine("    t, err := reader.Next()");
        buffer.addLine("    if err != nil {");
        buffer.addLine("      if err == io.EOF {");
        buffer.addLine("        break");
        buffer.addLine("      }");
        buffer.addLine("      return nil, err");
        buffer.addLine("    }");
        buffer.addLine("    t = xml.CopyToken(t)");
        buffer.addLine("    switch t := t.(type) {");
        buffer.addLine("    case xml.StartElement:");
        buffer.addLine("      one, err := reader.ReadString(&t)");
        buffer.addLine("      if err != nil {");
        buffer.addLine("        return nil, err");
        buffer.addLine("      }");
        buffer.addLine("      results = append(results, %1$s(one))", typeName.getSimpleName());
        buffer.addLine("	case xml.EndElement:");
        buffer.addLine("      depth--");
        buffer.addLine("    }");
        // End of for
        buffer.addLine("  }");

        buffer.addLine("  return results, nil");
        // End of function
        buffer.addLine("}");
        buffer.addLine();
    }

}

