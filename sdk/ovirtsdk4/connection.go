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
	"errors"
	"net/url"
	"net/http"
	"log"
	"time"
)


type connection struct {
	Url   *url.URL
	Username string
	Password string
	Insecure bool
	Debug    bool
	Log 	 *log.Logger
	Timeout  uint8
	Headers map[string]string
	Pipeline uint8
	Connections uint8
	Client *http.Client
}

func NewConnection(url string, username string, password string, insecure bool, debug bool, timeout uint8) *connection {
	c := new(Connection)
	c.Username = username
	c.Password = password
	c.Insecure = insecure
	c.Debug = debug
	c.timeout = timeout
	// init Log
	c.Log = log.New(os.Stdout, "[oVirtGo]", log.LstdFlags)
	// Check url is valid
	theUrl, err := url.Parse(url)
	if err != nil {
		return nil
	}
	c.Url = theUrl
	c.Client = &http.Client{
		Timeout: time.Duration(timeout)
	}
	return c
}
