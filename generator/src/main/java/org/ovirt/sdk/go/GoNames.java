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

import java.util.Set;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.tool.ReservedWords;
import org.ovirt.api.metamodel.tool.SchemaNames;
import org.ovirt.api.metamodel.tool.Words;

/**
 * This class contains the rules used to calculate the names of generated Go concepts.
 */
@ApplicationScoped
public class GoNames {

    // Reference to the object used to do computations with words.
    @Inject
    private Words words;

    // We need the Go reserved words in order to avoid producing names that aren't legal:
    @Inject
    @ReservedWords(language = "go")
    private Set<String> reservedWords;

    // Reference to the object used to calculate XML schema names:
    @Inject private SchemaNames schemaNames;

    /**
     * pointer represention of Name
     */
    public String withPointer(String name) {
        return "*" + name;
    }

    /**
     * slice represention of Name
     */
    public String withSlice(String name) {
        return "[]" + name;
    }

    /**
     * Returns a representation of the given name typically used for Go Exportable classes.
     */
    public String getExportableClassStyleName(Name name) {
        return renameReserved(exportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for Go Un-Exportable classes.
     */
    public String getUnexportableClassStyleName(Name name) {
        return renameReserved(unexportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for Go Exportable members.
     */
    public String getExportableMemberStyleName(Name name) {
        return renameReserved(exportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for Go Un-Exportable members.
     */
    public String getUnexportableMemberStyleName(Name name) {
        return renameReserved(unexportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for the name of Go Exported method received by struct.
     */
    public String getExportableMethodStyleName(Name name) {
        return renameReserved(exportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for the name of Go Un-Exported method received by struct.
     */
    public String getUnexportableMethodStyleName(Name name) {
        return renameReserved(unexportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for the name of Go Exportable functions.
     */
    public String getExportableFuncStyleName(Name name) {
        return renameReserved(exportableStyle(name));
    }

    /**
     * Returns a representation of the given name typically used for the name of Go Un-Exportable functions.
     */
    public String getUnexportableFuncStyleName(Name name) {
        return renameReserved(unexportableStyle(name));
    }

    /**
     * Returns a representation of the given name using the lower style typically used for Go tag name.
     * if name is `id` `href` `rel`, the tag is the **attribute**
     */
    public String getTagStyleName(Name name) {
        return schemaNames.getSchemaTagName(name);
    }

    /**
     * Returns a representation of the given name typically used for Go method parameters.
     */
    public String getParameterStyleName(Name name) {
        return renameReserved(unexportableStyle(name));
    }

    /**
     * In case of the same with the parameter name, so add _ as suffix
     */
    public String getVariableStyleName(Name name) {
        String result = name.words().map(words::capitalize).collect(joining());
        return renameReserved(String.join("", result, "Var"));
    }

    /**
     * Returns a representation of the given name typically used for Go packages.
     */
    public String getModuleStyleName(Name name) {
        String result = name.words().map(String::toLowerCase).collect(joining("_"));
        return renameReserved(result);
    }

    /**
     * Returns a representation of the given name typically used for Go constants.
     */
    public String getConstantStyleName(Name name) {
        return name.words().map(String::toUpperCase).collect(joining("_"));
    }

    /*
     * Builds the Exportable style string from the given name
     */
    private String exportableStyle(Name name) {
        return name.words().map(words::capitalize).collect(joining());
    }

    /*
     * Builds the Un-Exportable style string from the given name
     */
    private String unexportableStyle(Name name) {
        StringBuilder buffer = new StringBuilder();
        name.words().findFirst().map(String::toLowerCase).ifPresent(buffer::append);
        name.words().skip(1).map(words::capitalize).forEach(buffer::append);
        return buffer.toString();
    }

    private String renameReserved(String result) {
        if (reservedWords.contains(result)) {
            result += "_";
        }
        return result;
    }
}

