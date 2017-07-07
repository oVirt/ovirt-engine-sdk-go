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

import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.Type;

/**
 * For in type in Go
 */
public class GoTypes {

    public static Boolean isGoPrimitiveType(Type type) {
        if (type instanceof PrimitiveType) {
            if (type == type.getModel().getBooleanType() ||
                type == type.getModel().getIntegerType() ||
                type == type.getModel().getDecimalType() ||
                type == type.getModel().getStringType()) {
                    return true;
                }
        }
        return false;
    }

}
