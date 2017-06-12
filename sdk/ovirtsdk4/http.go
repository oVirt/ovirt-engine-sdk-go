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

// OvRequest represents an HTTP request.
// This class is intended for internal use by other components of the SDK.
// Refrain from using it directly as there is no backwards compatibility
// guarantee.
type OvRequest struct {
	Method  string
	Path    string
	Query   map[string]string
	Headers map[string]string
	Body    string
}

// NewOvRequest  ...
func NewOvRequest(method string, path string, headers, query map[string]string, body string) *OvRequest {
	if headers == nil {
		headers = make(map[string]string)
	}
	if query == nil {
		query = make(map[string]string)
	}
	return &OvRequest{
		Method:  method,
		Headers: headers,
		Query:   query,
		Body:    body,
	}
}

// This class represents an HTTP response.
// This class is intended for internal use by other components of the SDK.
// Refrain from using it directly as there is no backwards compatibility
// guarantee.
type OvResponse struct {
	Body    string
	Code    int
	Headers map[string]string
	Message string
}
