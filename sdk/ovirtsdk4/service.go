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

// BaseService : This is the base for all the services of the SDK. It contains the
// utility methods used by all of them.
type BaseService struct {
	Connection *Connection
	Path       string
}

func checkFault(response *OvResponse) error {
	if len(response.Body) == 0 {
		detailStr := fmt.Sprintf("HTTP response code is %d", response.Code)
		return &Fault{Detail: &detailStr}
	}
	// xml unmarshal into Fault struct
	var faultVar Fault
	xml.Unmarshal([]byte(response.Body), &faultVar)
	return &faultVar
}

func (service *BaseService) internalGet(headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("GET", service.Path, headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	// If Get succeed
	if res.Code == 200 {
		return res, nil
	}
	// Get failed
	return nil, checkFault(res)
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
	// If Add succeed
	if Contains(res.Code, []int{200, 201, 202}) {
		return res, nil
	}
	// Add failed
	return nil, checkFault(res)
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
	// If Update succeed
	if res.Code == 200 {
		return res, nil
	}
	// Update faield
	return nil, checkFault(res)
}

// Executes a `remove` method
func (service *BaseService) internalRemove(headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("DELETE", service.Path, headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	// If Remove succeed
	if res.Code == 200 {
		return res, nil
	}
	// Remove failed
	return nil, checkFault(res)
}

// Executes an `action` method
func (service *BaseService) internalAction(action *Action, path string, headers, query map[string]string, wait bool) (*OvResponse, error) {
	req := NewOvRequest("POST", fmt.Sprintf("%s/%s", service.Path, path), headers, query, "")
	res, err := service.Connection.Send(req)
	if err != nil {
		return nil, err
	}
	// If Action succeed
	if Contains(res.Code, []int{200, 201, 202}) {
		return res, nil
		// To unmarsh Action, not used now
		// 		var actionVar Action
		// 		xml.Unmarshal([]byte(res.Body), &actionVar)
	}
	return nil, checkFault(res)
}
