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
	"encoding/xml"
	"fmt"
)

// This is the base for all the services of the SDK. It contains the
// utility methods used by all of them.
type BaseService struct {
	Connection *Connection
	Path       string
}

func (service *BaseService) internalGet(headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("GET", service.Path, headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Executes an `add` method.
func (service *BaseService) internalAdd(object interface{}, headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("POST", service.Path, headers, query, "")
	xmlBytes, err := xml.Marshal(object)
	if err != nil {
		return nil, err
	}
	req.Body = string(xmlBytes)
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Executes a `update` method
func (service *BaseService) internalUpdate(object interface{}, headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("PUT", service.Path, headers, query, "")
	xmlBytes, err := xml.Marshal(object)
	if err != nil {
		return nil, err
	}
	req.Body = string(xmlBytes)
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Executes a `remove` method
func (service *BaseService) internalRemove(headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("DELETE", service.Path, headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Executes an `action` method
func (service *BaseService) internalAction(action *Action, path string, headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("POST", fmt.Sprintf("%s/%s", service.Path, path), headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
