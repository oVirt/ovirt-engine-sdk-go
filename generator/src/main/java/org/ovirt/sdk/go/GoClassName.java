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

/**
 * This class represents the fully qualified name of a Go class, composed by the package name and the struct.
 */
public class GoClassName {
    private String packageName;
    private String simpleName;

    public String getPackageName() {
        return packageName;
    }

    public void setPackageName(String newPackageName) {
        packageName = newPackageName;
    }

    public String getSimpleName() {
        return simpleName;
    }

    public void setSimpleName(String newClassName) {
        simpleName = newClassName;
    }

    public void setClass(Class<?> newClass) {
        packageName = newClass.getPackage().getName();
        simpleName = newClass.getSimpleName();
    }

    @Override
    public String toString() {
        StringBuilder buffer = new StringBuilder();
        if (packageName != null && packageName.length() > 0) {
            buffer.append(packageName);
            buffer.append("/");
        }
        buffer.append(simpleName);
        return buffer.toString();
    }
}

