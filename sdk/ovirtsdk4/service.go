//
// Copyright (c) 2017 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ovirtsdk4

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BaseService represents the base for all the services of the SDK. It contains the
// utility methods used by all of them.
type BaseService struct {
	Connection *Connection
	Path       string
}

// CheckFault procoesses error parsing and returns it back
func CheckFault(response *http.Response) error {
	resBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response, reason: %s", err.Error())
	}

	var fault Fault
	err = xml.Unmarshal(resBytes, &fault)
	if err != nil {
		return fmt.Errorf("Failed to read response, reason: %s", err.Error())
	}

	return BuildError(response, fault)
}

// CheckAction checks if response contains an Action instance
func CheckAction(response *http.Response) (*Action, error) {
	resBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response, reason: %s", err.Error())
	}
	// Check if is Fault instance
	var faultZero Fault
	var fault Fault
	err = xml.Unmarshal(resBytes, &fault)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response, reason: %s", err.Error())
	}
	if fault != faultZero {
		return nil, BuildError(response, fault)
	}

	var action Action
	err = xml.Unmarshal(resBytes, &action)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response, reason: %s", err.Error())
	}
	if action.Fault != nil {
		return nil, BuildError(response, *action.Fault)
	}
	return &action, nil
}

// BuildError constructs error
func BuildError(response *http.Response, fault Fault) error {
	var buffer bytes.Buffer

	if fault.Reason != nil {
		if buffer.Len() > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("Fault reason is \"%s\".", *fault.Reason))
	}
	if fault.Detail != nil {
		if buffer.Len() > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("Fault detail is \"%s\".", *fault.Detail))
	}
	if response != nil {
		if buffer.Len() > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("HTTP response code is \"%d\".", response.StatusCode))
		buffer.WriteString(" ")
		buffer.WriteString(fmt.Sprintf("HTTP response message is \"%s\".", response.Status))
	}
	return errors.New(buffer.String())
}
