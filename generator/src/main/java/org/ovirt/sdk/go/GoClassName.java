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

/**
 * This class represents the fully qualified name of a Go class, composed by the package name and the struct.
 */
public class GoClassName {
    private String packageName;
    private String className;

    public String getPackageName() {
        return packageName;
    }

    public void setPackageName(String newPackageName) {
        packageName = newPackageName;
    }

    public String getClassName() {
        return className;
    }

    public String getPrivateClassName() {
        return getClassName().substring(0, 1).toLowerCase() + getClassName().substring(1);
    }

    public void setClassName(String newClassName) {
        className = newClassName;
    }

    @Override
    public String toString() {
        StringBuilder buffer = new StringBuilder();
        buffer.append(packageName);
        buffer.append("/");
        buffer.append(className);
        return buffer.toString();
    }
}

