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

import javax.enterprise.context.ApplicationScoped;

/**
 * This class contains the rules used to calculate the names of generated Go packages.
 */
@ApplicationScoped
public class GoPackages {

    // The relative names of the packages:
    public static final String READERS_PACKAGE = "readers";
    public static final String SERVICES_PACKAGE = "services";
    public static final String TYPES_PACKAGE = "types";
    public static final String WRITERS_PACKAGE = "writers";
    public static final String VERSION_PACKAGE = "version";

    // The name of the root package:
    private String rootPackageName = "ovirtsdk";

    // The version number:
    private String version;

    // root package url prefix
    private String rootPackageUrlPrefix = "github.com/ovirt/go-ovirt/v4";

    public void setRootPackageUrlPrefix(String newRootPackageUrlPrefix) {
        rootPackageUrlPrefix = newRootPackageUrlPrefix;
    }

    public String getRootPackageUrlPrefix() {
        return rootPackageUrlPrefix;
    }

    /**
     * Sets the version.
     */
    public void setVersion(String newVersion) {
        version = newVersion;
    }

    /**
     * Get the version.
     */
    public String getVersion() {
        return version;
    }

    /**
     * Get the name of the root package.
     */
    public String getRootPackageName() {
        return rootPackageName;
    }

    /**
     * Get the name of the types package.
     */
    public String getTypesPackageName() {
        return getPackageName(TYPES_PACKAGE);
    }

    /**
     * Get the name of the models builder package.
     *  - just use type package name.
     */
    public String getTypeBuildersPackageName() {
        return getTypesPackageName();
    }

    /**
     * Get the name of the readers package.
     */
    public String getReadersPackageName() {
        return getPackageName(READERS_PACKAGE);
    }

    /**
     * Get the name of the writers package.
     */
    public String getWritersPackageName() {
        return getPackageName(WRITERS_PACKAGE);
    }

    /**
     * Get the name of the services package.
     */
    public String getServicesPackageName() {
        return getPackageName(SERVICES_PACKAGE);
    }

    /**
     * Get the name of the version package
     */
    public String getVersionPackageName() {
        return getPackageName(VERSION_PACKAGE);
    }

    /**
     * Get the complete name of the given package.
     */
    public String getPackageName(String... relativeNames) {
        StringBuilder buffer = new StringBuilder();
        buffer.append(rootPackageName);
        if (relativeNames != null && relativeNames.length > 0) {
            for (String relativeName : relativeNames) {
                buffer.append('/');
                buffer.append(relativeName);
            }
        }
        return buffer.toString();
    }

}