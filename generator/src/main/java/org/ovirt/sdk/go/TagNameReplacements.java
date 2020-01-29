/*
Copyright (c) 2017 Joey <majunjiev@gmail.com>.

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
import java.util.HashMap;
import java.util.Map;
import javax.annotation.PostConstruct;
import javax.enterprise.inject.Produces;
import javax.inject.Singleton;


/**
 * This class is a producer of the set of Go reserved words.
 */
@Singleton
public class TagNameReplacements {
    private Map<String, String> words;

    @PostConstruct
    private void init() {
        // Create the set:
        words = new HashMap<>();

        words.put("boot_device", "device");
        words.put("boot_devices", "devices");

        // Wrap the set so that it is unmodifiable:
        words = Collections.unmodifiableMap(words);
    }

    /**
     * Produces the set of Go reserved words.
     */
    @Produces
    @TagNames
    public Map<String, String> getWords() {
        return words;
    }
}
