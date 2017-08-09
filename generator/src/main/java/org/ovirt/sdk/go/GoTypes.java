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

import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.EnumType;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.Words;

/**
 * For in type in Go
 */
public class GoTypes {

    @Inject private GoNames goNames;

    // Reference to the object used to do computations with words.
    @Inject
    private Words words;

    /**
     * Calculates the name of the builder class that should be generated for the given type. For example,
     * for the {@code Vm} type it will generate {@code org.ovirt.engine.model} as the package name and
     * {@code V4VmBuilder} as the simple class name.
     */
    public String getBuilderName(Type type) {
        String typeName = goNames.getTypeName(type).getClassName();
        typeName = typeName.substring(0, 1).toLowerCase() + typeName.substring(1);
        String result = String.join("", typeName, "Builder");
        return goNames.renameReserved(result);
    }


    public String getNewBuilderFuncName(Type type) {
        GoClassName typeName = goNames.getTypeName(type);
        String result = String.join("", "New", typeName.getClassName(), "Builder");
        return result;
    }

    public String getXmlReadOneFuncName(Type type) {
        GoClassName typeName = goNames.getTypeName(type);
        String result = String.join("", "XML", typeName.getClassName(), "ReadOne");
        return goNames.renameReserved(result);
    }

    public String getXmlReadManyFuncName(Type type) {
        GoClassName typeName = goNames.getTypeName(type);
        String result = String.join("", "XML", typeName.getClassName(), "ReadMany");
        return goNames.renameReserved(result);
    }

    public String getXmlWriteOneFuncName(Type type) {
        GoClassName typeName = goNames.getTypeName(type);
        String result = String.join("", "XML", typeName.getClassName(), "WriteOne");
        return goNames.renameReserved(result);
    }

    public String getXmlWriteManyFuncName(Type type) {
        GoClassName typeName = goNames.getTypeName(type);
        String result = String.join("", "XML", typeName.getClassName(), "WriteMany");
        return goNames.renameReserved(result);
    }

    public String getMemberSetterMethodName(Name name) {
        String result = name.words().map(words::capitalize).collect(joining());
        return goNames.renameReserved("Set" + result);
    }

    public String getMemberGetterMethodName(Name name) {
        String result = name.words().map(words::capitalize).collect(joining());
        return goNames.renameReserved(result);
    }

    public String getMemberMustGetterMethodName(Name name) {
        return "Must" + getMemberGetterMethodName(name);
    }

    public String getMemberPresentMethodName(Name name) {
        String result = name.words().map(words::capitalize).collect(joining());
        return goNames.renameReserved(result + "Present");
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

}
