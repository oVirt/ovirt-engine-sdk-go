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

import java.util.ArrayList;
import java.util.List;

/**
 * This class represents a reference to a Go type, including all the imports that are necessary to use it.
 */
public class GoTypeReference {
    private String text;
    private List<String> imports = new ArrayList<>(1);

    public String getText() {
        return text;
    }

    public void setText(String newText) {
        text = newText;
    }

    public void setText(Class<?> clazz) {
        text = clazz.getSimpleName();
    }

    public List<String> getImports() {
        return new ArrayList<>(imports);
    }

    public void setImports(List<String> newImports) {
        imports.clear();
        imports.addAll(newImports);
    }

    public void addImport(String newImport) {
        if (newImport != null && newImport.length() >0 ) {
            imports.add(newImport);
        }
    }

    @Override
    public String toString() {
        return text;
    }
}

