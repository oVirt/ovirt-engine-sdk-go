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

import java.util.Collections;
import java.util.HashSet;
import java.util.Set;
import javax.annotation.PostConstruct;
import javax.enterprise.inject.Produces;
import javax.inject.Singleton;

import org.ovirt.api.metamodel.tool.ReservedWords;

/**
 * This class is a producer of the set of Python reserved words.
 */
@Singleton
public class GoReservedWords {
    private Set<String> words;

    @PostConstruct
    private void init() {
        // Create the set:
        words = new HashSet<>();

        // type用于声明自定义类型
        // map用于声明map类型数据
        // range用于读取slice、map、channel数据
        // definition
        words.add("var");
        words.add("const");
        words.add("interface");
        words.add("struct");
        words.add("type");
        words.add("func");
        words.add("package");
        words.add("import");
        words.add("make");
        // type
        words.add("false");
        words.add("nil");
        words.add("true");
        words.add("map");
        words.add("bool");
        words.add("float64");
        words.add("int");
        words.add("int8");
        words.add("int16");
        words.add("int32");
        words.add("int64");
        words.add("uint");
        words.add("uint8");
        words.add("uint16");
        words.add("uint32");
        words.add("uint64");
        words.add("uintptr");
        words.add("string");
        // flow control
        words.add("break");
        words.add("case");
        words.add("switch");
        words.add("continue");
        words.add("for");
        words.add("if");
        words.add("else");
        words.add("goto");
        words.add("default");
        words.add("fallthrough");
        words.add("return");
        words.add("defer");
        words.add("range");
        // parallel
        words.add("go");
        words.add("chan");
        words.add("select");
        // others
        words.add("panic");

        // Wrap the set so that it is unmodifiable:
        words = Collections.unmodifiableSet(words);
    }

    /**
     * Produces the set of Go reserved words.
     */
    @Produces
    @ReservedWords(language = "go")
    public Set<String> getWords() {
        return words;
    }
}
