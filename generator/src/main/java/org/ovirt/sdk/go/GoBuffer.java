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

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.io.Writer;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.Formatter;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.apache.commons.io.FileUtils;

/**
 * This class is a buffer intended to simplify generation of Go source code. It stores the name of the module, the
 * list of imports and the rest of the source separately, so that imports can be added on demand while generating the
 * rest of the source.
 */
public class GoBuffer {

    // The name of the file:
    private String fileName;

    // The name of the package:
    private String packageName;

    // The imports:
    private Set<String> imports = new HashSet<>();

    // The source lines:
    private List<String> lines = new ArrayList<>();

    /**
     * Sets the file name.
     */
    public void setFileName(String newFileName) {
        fileName = newFileName;
    }

    /**
     * Sets the module name:
     */
    public void setPackageName(String newPackageName) {
        packageName = newPackageName;
    }

    /**
     * Adds one import.
     */
    public void addImport(String format, Object ... args) {
        // Format the line:
        StringBuilder buffer = new StringBuilder();
        Formatter formatter = new Formatter(buffer);
        formatter.format(format, args);

        // Add the line to the set:
        imports.add(buffer.toString());
        formatter.close();
    }

    /**
     * Adds multiple imports.
     */
    public void addImports(List<String> imports) {
        if (imports != null && imports.size() > 0) {
            imports.forEach(this::addImport);
        }
    }

    /**
     * Adds an empty line to the body of the class.
     */
    public void addLine() {
        lines.add("");
    }

    /**
     * Adds a formatted line to the body of the class. The given {@code args} are formatted using the
     * provided {@code format} using the {@link String#format(String, Object...)} method.
     */
    public void addLine(String format, Object ... args) {
        // Format the line:
        StringBuilder buffer = new StringBuilder();
        Formatter formatter = new Formatter(buffer);
        formatter.format(format, args);

        // Add the line to the list:
        lines.add(buffer.toString());
        formatter.close();
    }

    /**
     * Adds a non-formatted line to the body of the class.
     */
    public void addRawLine(String line) {
        StringBuilder buffer = new StringBuilder(line);

        // Add the line to the list:
        lines.add(buffer.toString());
    }

    /**
     * Adds a non-formatted comment line
     */
    public void addCommentLine(String line) {
        addRawLine("// " + line);
    }

    /**
     * Starts a multi line comment.
     */
    public void startComment() {
        addLine("//");
    }

    /**
     * Ends a multi line comment.
     */
    public void endComment() {
        addLine("//");
    }

    /**
     * Generates the complete source code of the class.
     */
    public String toString() {
        StringBuilder buffer = new StringBuilder();

        // License:
        buffer.append("//\n");
        buffer.append("// The oVirt Project - oVirt Engine Go SDK");
        buffer.append("//\n");
        buffer.append("//Copyright (c) oVirt Authors\n");
        buffer.append("//\n");
        buffer.append("// Licensed under the Apache License, Version 2.0 (the \"License\");\n");
        buffer.append("// you may not use this file except in compliance with the License.\n");
        buffer.append("// You may obtain a copy of the License at\n");
        buffer.append("//\n");
        buffer.append("//   http://www.apache.org/licenses/LICENSE-2.0\n");
        buffer.append("//\n");
        buffer.append("// Unless required by applicable law or agreed to in writing, software\n");
        buffer.append("// distributed under the License is distributed on an \"AS IS\" BASIS,\n");
        buffer.append("// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n");
        buffer.append("// See the License for the specific language governing permissions and\n");
        buffer.append("// limitations under the License.\n");
        buffer.append("//\n");
        buffer.append("");

        // Package:
        buffer.append("package ovirtsdk");
        buffer.append("\n");

        // Add the imports:
        if (!imports.isEmpty()) {
            buffer.append("import (\n");
            imports.stream().sorted().forEach(line -> {
                buffer.append("\"" + line + "\"");
                buffer.append("\n");
            });
            buffer.append(")\n");
            buffer.append("\n");
        }

        // Body:
        for (String line : lines) {
            buffer.append(line);
            buffer.append("\n");
        }

        return buffer.toString();
    }

    /**
     * Creates a {@code .go} source file and writes the source. The required intermediate directories will be created
     * if they don't exist.
     *
     * @param dir the base directory for the source code
     * @throws IOException if something fails while creating or writing the file
     */
    public void write(File dir) throws IOException {
        // Calculate the complete file name:
        if (fileName == null) {
            fileName = packageName;
        }
        File file = new File(dir, fileName + ".go");

        // Create the directory and all its parent if needed:
        File parent = file.getParentFile();
        FileUtils.forceMkdir(parent);

        // Write the file:
        System.out.println("Writing file \"" + file.getAbsolutePath() + "\".");
        try (Writer writer = new OutputStreamWriter(new FileOutputStream(file), StandardCharsets.UTF_8)) {
            writer.write(toString());
        }

        // Run the gofmt tool to format the file according to the standard Go rules:
        Process gofmt = new ProcessBuilder()
            .command("gofmt", "-w", file.getAbsolutePath())
            .inheritIO()
            .start();
        try {
            int code = gofmt.waitFor();
            if (code != 0) {
                throw new IOException(
                    "Execution of \"gofmt\" on file \"" + file.getAbsolutePath() + "\" failed with exit code " + code
                );
            }
        }
        catch (InterruptedException exception) {
            throw new IOException(
                "Interrupted while waiting for to execute \"gofmt\" on file \"" + file.getAbsolutePath() + "\"",
                exception
            );
        }
    }
}
