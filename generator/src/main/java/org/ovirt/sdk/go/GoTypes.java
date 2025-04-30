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

import java.util.List;

import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.EnumType;
import org.ovirt.api.metamodel.concepts.ListType;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.NameParser;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.Service;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;

/**
 * This class calculates the type references for the Go structs and the Go funcs generated from the model.
 */
public class GoTypes {

    // Suffixes to add to the different types of classes generated for each type:
    private static final Name SLICE_NAME = NameParser.parseUsingCase("Slice");
    private static final Name BUILDER_NAME = NameParser.parseUsingCase("Builder");
    private static final Name READER_NAME = NameParser.parseUsingCase("Reader");
    private static final Name WRITER_NAME = NameParser.parseUsingCase("Writer");
    private static final Name SERVICE_NAME = NameParser.parseUsingCase("Service");
    
    // Suffixes to add to the different functions for each type
    private static final Name READ_ONE_SUFFIX = NameParser.parseUsingCase("ReadOne");
    private static final Name READ_MANY_SUFFIX = NameParser.parseUsingCase("ReadMany");
    private static final Name WRITE_ONE_SUFFIX = NameParser.parseUsingCase("WriteOne");
    private static final Name WRITE_MANY_SUFFIX = NameParser.parseUsingCase("WriteMany");
    private static final Name BUILDER_SUFFIX = NameParser.parseUsingCase("Builder");
    private static final Name PRESENT_SUFFIX = NameParser.parseUsingCase("Present");

    // Prefixes for the XML and JSON readers and writers:
    private static final Name XML_PREFIX = NameParser.parseUsingCase("XML");
    // Prefixes for functions
    private static final Name NEW_PREFIX = NameParser.parseUsingCase("New");
    private static final Name SET_PREFIX = NameParser.parseUsingCase("Set");
    private static final Name MUST_PREFIX = NameParser.parseUsingCase("Must");


    // Reference to the object used to do computations with words.
    @Inject private GoNames goNames;
    @Inject private GoPackages goPackages;

    /**
     * Calculates the Go name that corresponds to the given type.
     */
    public GoClassName getTypeName(Type type) {
        return getClassName(type, goPackages.getTypesPackageName(), null, null, true);
    }

    public GoClassName getTypeSliceName(Type type) {
        return getClassName(type, goPackages.getTypesPackageName(), null, SLICE_NAME, true);
    }

    /**
     * Calculates the name of the builder class that should be generated for the given type. For example,
     * for the {@code Vm} type it will generate {@code org.ovirt.engine.model} as the package name and
     * {@code VmBuilder} as the simple class name.
     */
    public GoClassName getBuilderName(Type type) {
        return getClassName(type, goPackages.getTypeBuildersPackageName(), null, BUILDER_NAME, true);
    }

    /**
     * Calculates the Go name of the xml-reader for the given type.
     */
    public GoClassName getXMLReaderName(Type type) {
        return getClassName(type, goPackages.getReadersPackageName(), XML_PREFIX, READER_NAME, true);
    }

    /**
     * Calculates the Go name of the xml-writer for the given type.
     */
    public GoClassName getXMLWriterName(Type type) {
        return getClassName(type, goPackages.getWritersPackageName(), XML_PREFIX, WRITER_NAME, true);
    }

    public GoFuncName getNewBuilderFuncName(Type type) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(type.getName(), NEW_PREFIX, BUILDER_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getNewServiceFuncName(Service service) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(service.getName(), NEW_PREFIX, SERVICE_NAME);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getXmlReadOneFuncName(Type type) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(type.getName(), XML_PREFIX, READ_ONE_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getXmlReadManyFuncName(Type type) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(type.getName(), XML_PREFIX, READ_MANY_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getXmlWriteOneFuncName(Type type) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(type.getName(), XML_PREFIX, WRITE_ONE_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getXmlWriteManyFuncName(Type type) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(type.getName(), XML_PREFIX, WRITE_MANY_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getMemberSetterMethodName(Name name) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(name, SET_PREFIX, null);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getMemberGetterMethodName(Name name) {
        GoFuncName funcName = new GoFuncName();
        funcName.setSimpleName(goNames.getExportableFuncStyleName(name));
        return funcName;
    }

    public GoFuncName getMemberMustGetterMethodName(Name name) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(name, MUST_PREFIX, null);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    public GoFuncName getMemberPresentMethodName(Name name) {
        GoFuncName funcName = new GoFuncName();
        Name newName = decorateName(name, null, PRESENT_SUFFIX);
        funcName.setSimpleName(goNames.getExportableFuncStyleName(newName));
        return funcName;
    }

    private GoClassName getClassName(Type type, String packageName, Name prefix, Name suffix, boolean exportable) {
        if (type instanceof PrimitiveType) {
            return getPrimitiveTypeName((PrimitiveType) type, false);
        }
        if (type instanceof StructType || type instanceof EnumType) {
            Name name = decorateName(type.getName(), prefix, suffix);
            GoClassName typeName = new GoClassName();
            // typeName.setPackageName(packageName);
            if (exportable) {
                typeName.setSimpleName(goNames.getExportableClassStyleName(name));
            } else {
                typeName.setSimpleName(goNames.getUnexportableClassStyleName(name));
            }
            return typeName;
        }
        throw new RuntimeException("Don't know how to calculate the Java type name for type \"" + type + "\"");
    }

    private GoClassName getPrimitiveTypeName(PrimitiveType type, boolean pointer) {
        GoClassName name = new GoClassName();
        Model model = type.getModel();
        if (type == model.getBooleanType()) {
            name.setSimpleName(pointer ? "*bool" : "bool");
        }
        else if (type == model.getIntegerType()) {
            name.setSimpleName(pointer ? "*int64" : "int64");
        }
        else if (type == model.getDecimalType()) {
            name.setSimpleName(pointer ? "*float64" : "float64");
        }
        else if (type == model.getStringType()) {
            name.setSimpleName(pointer ? "*string" : "string");
        }
        else if (type == model.getDateType()) {
            name.setPackageName("time");
            name.setSimpleName(pointer ? "*time.Time" : "time.Time");
        }
        else {
            throw new RuntimeException("Don't know how to calculate the Java type reference for type \"" + type + "\"");
        }
        return name;
    }

    private GoTypeReference getTypeReference(Type type, boolean pointer) {
        if (type instanceof PrimitiveType) {
            GoClassName name = getPrimitiveTypeName((PrimitiveType) type, pointer);
            GoTypeReference reference = new GoTypeReference();
            reference.addImport(name.getPackageName());
            reference.setText(name.getSimpleName());
            return reference;
        }
        if (type instanceof StructType) {
            return getStructReference((StructType) type, pointer);
        }
        if (type instanceof EnumType) {
            return getEnumReference((EnumType) type, pointer);
        }
        if (type instanceof ListType) {
            return getListReference((ListType) type, pointer);
        }
        throw new RuntimeException("Don't know how to calculate the Java type reference for type \"" + type + "\"");
    }

    public GoTypeReference getTypeReferenceAsParameter(Type type) {
        GoTypeReference reference = getTypeReference(type, false);
        if (type instanceof StructType) {
            reference.setText(goNames.withPointer(reference.getText()));
        }
        return reference;
    }

    public GoTypeReference getTypeReferenceAsAttribute(Type type) {
        return getTypeReference(type, true);
    }

    public GoTypeReference getTypeReferenceAsVariable(Type type) {
        if (type instanceof StructType) {
            return getTypeReference(type, true);
        }
        return getTypeReference(type, false);
    }

    public GoTypeReference getTypeReferenceAsReturnvalue(Type type) {
        if (type instanceof StructType) {
            return getTypeReference(type, true);
        }
        return getTypeReference(type, false);
    }

    private GoTypeReference getStructReference(StructType type, boolean pointer) {
        GoTypeReference reference = new GoTypeReference();
        String text = goNames.getExportableClassStyleName(type.getName());
        if (pointer) {
            text = goNames.withPointer(text);
        }
        reference.setText(text);
        return reference;
    }

    private GoTypeReference getEnumReference(EnumType type, boolean pointer) {
        GoTypeReference reference = new GoTypeReference();
        String text = goNames.getExportableClassStyleName(type.getName());
        if (pointer) {
            text = goNames.withPointer(text);
        }
        reference.setText(text);
        return reference;
    }

    private GoTypeReference getListReference(ListType type, boolean pointer) {
        GoTypeReference reference = new GoTypeReference();
        Type elementType = type.getElementType();
        if (elementType instanceof StructType) {
            GoClassName typeSlice = getTypeSliceName(elementType);
            reference.addImport(typeSlice.getPackageName());
            reference.setText(goNames.withPointer(typeSlice.getSimpleName()));
        } else {
            GoTypeReference elementTypeReference = getTypeReference(elementType, false);
            // use Recursion to return []EnumType / []string
            reference.setImports(elementTypeReference.getImports());
            reference.setText("[]" + elementTypeReference.getText());
        }
        return reference;
    }

    public Boolean isGoPrimitiveType(Type type) {
        if (type instanceof PrimitiveType) {
            if (type == type.getModel().getBooleanType() ||
                type == type.getModel().getIntegerType() ||
                type == type.getModel().getDecimalType() ||
                type == type.getModel().getStringType()  ||
                type == type.getModel().getDateType()) {
                    return true;
                }
        }
        return false;
    }

    /**
     * Decorates a name with optional prefix and suffix.
     */
    private Name decorateName(Name name, Name prefix, Name suffix) {
        List<String> words = name.getWords();
        if (prefix != null) {
            words.addAll(0, prefix.getWords());
        }
        if (suffix != null) {
            words.addAll(suffix.getWords());
        }
        return new Name(words);
    }

}
