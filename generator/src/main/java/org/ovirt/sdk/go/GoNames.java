/*
Copyright (c) 2016-2017 Red Hat, Inc.

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

import java.util.List;
import java.util.Set;
import javax.enterprise.context.ApplicationScoped;
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
import org.ovirt.api.metamodel.tool.ReservedWords;
import org.ovirt.api.metamodel.tool.Words;

/**
 * This class contains the rules used to calculate the names of generated Go concepts.
 */
@ApplicationScoped
public class GoNames {
    // The names of the base classes:
    public static final Name READER_NAME = NameParser.parseUsingCase("Reader");
    public static final Name SERVICE_NAME = NameParser.parseUsingCase("Service");
    public static final Name WRITER_NAME = NameParser.parseUsingCase("Writer");

    // The relative names of the packages:
    public static final String READERS_PACKAGE = "readers";
    public static final String SERVICES_PACKAGE = "services";
    public static final String TYPES_PACKAGE = "types";
    public static final String WRITERS_PACKAGE = "writers";
    public static final String VERSION_PACKAGE = "version";

    // Reference to the object used to do computations with words.
    @Inject
    private Words words;

    // We need the Go reserved words in order to avoid producing names that aren't legal:
    @Inject
    @ReservedWords(language = "go")
    private Set<String> reservedWords;

    // The name of the root package:
    private String rootPackageName = "ovirtsdk4";

    // The version number:
    private String version;

    // root package url prefix
    private String rootPackageUrlPrefix = "github.com/imjoey/sdk";

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
        if (relativeNames != null || relativeNames.length > 0) {
            for (String relativeName : relativeNames) {
                buffer.append('/');
                buffer.append(relativeName);
            }
        }
        return buffer.toString();
    }

    /**
     * Calculates the Go name that corresponds to the given type.
     */
    public GoClassName getTypeName(Type type) {
        return buildClassName(type.getName(), null, TYPES_PACKAGE);
    }

    /**
     * Calculates that should be used in Go to reference the given type. For example, for the boolean type it will
     * return the {@code bool} string.
     */
    public GoTypeReference getTypeReference(Type type) {
        GoTypeReference reference = new GoTypeReference();
        if (type instanceof PrimitiveType) {
            Model model = type.getModel();
            if (type == model.getBooleanType()) {
                reference.setText("bool");
            }
            else if (type == model.getIntegerType()) {
                reference.setText("int64");
            }
            else if (type == model.getDecimalType()) {
                reference.setText("float64");
            }
            else if (type == model.getStringType()) {
                reference.setText("string");
            }
            else if (type == model.getDateType()) {
                reference.addImport("time");
                reference.setText("time.Time");
            }
            else {
                throw new IllegalArgumentException(
                    "Don't know how to build reference for primitive type \"" + type + "\""
                );
            }
        }
        else if (type instanceof StructType || type instanceof EnumType) {
            reference.setText("*" + getTypeName(type).getClassName());
        }
        else if (type instanceof EnumType) {
            reference.setText(getTypeName(type).getClassName());
        }
        else if (type instanceof ListType) {
            ListType listtype = (ListType)type;
            reference.setText("[]*" + getClassStyleName(listtype.getElementType().getName()));
        }
        else {
            throw new IllegalArgumentException("Don't know how to build reference for type \"" + type + "\"");
        }
        return reference;
    }
    /**
     * Calculates the Go name of the base class of the services.
     */
    public GoClassName getBaseServiceName() {
        return buildClassName(SERVICE_NAME, null, SERVICES_PACKAGE);
    }

    /**
     * Calculates the Go name that corresponds to the given service.
     */
    public GoClassName getServiceName(Service service) {
        return buildClassName(service.getName(), SERVICE_NAME, SERVICES_PACKAGE);
    }

    /**
     * Calculates the Go name of the reader for the given type.
     */
    public GoClassName getReaderName(Type type) {
        return buildClassName(type.getName(), READER_NAME, READERS_PACKAGE);
    }

    /**
     * Calculates the Go name of the writer for the given type.
     */
    public GoClassName getWriterName(Type type) {
        return buildClassName(type.getName(), WRITER_NAME, WRITERS_PACKAGE);
    }

    /**
     * Builds a Go name from the given base name, suffix, and package.
     *
     * The suffix can be {@code null} or empty, in that case then won't be added.
     *
     * @param base the base name
     * @param suffix the suffix to add to the name
     * @param package the package name
     * @return the calculated Go class name
     */
    private GoClassName buildClassName(Name base, Name suffix, String pkg) {
        List<String> words = base.getWords();
        if (suffix != null) {
            words.addAll(suffix.getWords());
        }
        Name name = new Name(words);
        GoClassName result = new GoClassName();
        result.setClassName(getClassStyleName(name));
        result.setPackageName(getPackageName(pkg));
        return result;
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go classes.
     */
    public String getClassStyleName(Name name) {
        return name.words().map(words::capitalize).collect(joining());
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go members.
     */
    public String getMemberStyleName(Name name) {
        return getClassStyleName(name);
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go method name.
     */
    public String getMethodStyleName(Name name) {
        return getClassStyleName(name);
    }

    public String getTagStyleName(Name name) {
        String result = name.words().map(String::toLowerCase).collect(joining("_"));
        if (reservedWords.contains(result)) {
            result += "_";
        }
        return result;
    }

    /**
     * Returns a representation of the given name using the non-capitalization style typically used for Go method parameters.
     */
    public String getParameterStyleName(Name name) {
        String result = getClassStyleName(name);
        return result.substring(0, 1).toLowerCase() + result.substring(1);
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go constants.
     */
    public String getConstantStyleName(Name name) {
        return name.words().map(String::toUpperCase).collect(joining("_"));
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go packages.
     */
    public String getModuleStyleName(Name name) {
        String result = name.words().map(String::toLowerCase).collect(joining("_"));
        if (reservedWords.contains(result)) {
            result += "_";
        }
        return result;
    }
}

