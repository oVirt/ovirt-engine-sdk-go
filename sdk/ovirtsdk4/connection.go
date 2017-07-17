//
// Copyright (c) 2017 Joey
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

// Some codes of this file is from https://github.com/CpuID/ovirt-engine-sdk-go/blob/master/sdk/http/http.go.
// And I made some bug fixes, Thanks to @CpuID

package ovirtsdk4

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Connection ... This type (and its attached functions) are responsible for managing an HTTP connection to the engine server.
// It is intended as the entry point for the SDK, and it provides access to the `system` service and, from there,
// to the rest of the services provided by the API.
type Connection struct {
	url      *url.URL
	username string
	password string
	token    string
	insecure bool
	caFile   string
	kerberos bool
	timeout  time.Duration
	compress bool
	//
	client *http.Client
}

// URL ... Returns the base URL of this connection.
func (c *Connection) URL() string {
	return c.url.String()
}

// SystemService : Returns a reference to the root of the services tree.
func (c *Connection) SystemService() *SystemService {
	return NewSystemService(c, "")
}

// Send : Sends an HTTP request and waits for the response.
func (c *Connection) Send(r *OvRequest) (*OvResponse, error) {
	var result OvResponse

	// Check if we already have an SSO access token:
	c.token, _ = c.getAccessToken()

	// Build the URL:
	useRawURL := c.buildRawURL(r.Path, r.Query)

	// Validate the method selected:
	if Contains(r.Method, []string{"DELETE", "GET", "PUT", "HEAD", "POST"}) == false {
		return &result, fmt.Errorf("The HTTP method '%s' is invalid, we expected one of DELETE/GET/PUT/HEAD/POST", r.Method)
	}

	// Build the net/http request:
	req, err := http.NewRequest(r.Method, useRawURL, nil)
	if err != nil {
		return &result, err
	}

	// Add request headers:
	for reqHK, reqHV := range r.Headers {
		req.Header.Add(reqHK, reqHV)
	}
	req.Header.Add("User-Agent", fmt.Sprintf("GoSDK/%s", SDK_VERSION))
	req.Header.Add("Version", "4")
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")
	// 		Generate base64(username:password)
	rawAuthStr := fmt.Sprintf("%s:%s", c.username, c.password)
	req.Header.Add("Authorization",
		fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(rawAuthStr))))

	// Send the request and wait for the response:
	resp, err := c.client.Do(req)
	if err != nil {
		return &result, err
	}

	// Return the response:
	defer resp.Body.Close()
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	result.Body = string(respBodyBytes)
	if err != nil {
		return &result, err
	}
	result.Code = resp.StatusCode
	result.Headers = make(map[string]string)
	for respHK, respHV := range resp.Header {
		result.Headers[respHK] = respHV[0]
	}

	return &result, nil
}

// Obtains the access token from SSO to be used for bearer authentication.
func (c *Connection) getAccessToken() (string, error) {
	// Build the URL and parameters required for the request:
	rawURL, parameters := c.buildSsoAuthRequest()

	// Send the response and wait for the request:
	response, _ := c.getSsoResponse(rawURL, parameters)

	// Top level array already handled in getSsoResponse() generically.
	if len(response.ssoError) > 0 {
		return "", fmt.Errorf("Error during SSO authentication %s: %s", response.ssoErrorCode, response.ssoError)
	}

	return response.accessToken, nil
}

// Revoke the SSO access token.
func (c *Connection) revokeAccessToken() error {
	// Build the URL and parameters required for the request:
	url, parameters := c.buildSsoRevokeRequest()

	// Send the response and wait for the request:
	response, _ := c.getSsoResponse(url, parameters)

	// Top level array already handled in getSsoResponse() generically.

	if len(response.ssoError) > 0 {
		return fmt.Errorf("Error during SSO revoke %s: %s", response.ssoErrorCode, response.ssoError)
	}
	return nil
}

// NewConnectionBuilder : Create the `ConnectionBuilder struct instance
func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{conn: &Connection{}, err: nil}
}

// ConnectionBuilder :  Builds the `Connection` struct
type ConnectionBuilder struct {
	conn *Connection
	err  error
}

// URL : Set the url field for `Connection` instance
func (connBuilder *ConnectionBuilder) URL(inputRawURL string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	// Save the URL:
	useURL, err := url.Parse(inputRawURL)
	if err != nil {
		connBuilder.err = err
		return connBuilder
	}
	connBuilder.conn.url = useURL
	return connBuilder
}

// Username : Set the username field for `Connection` instance
func (connBuilder *ConnectionBuilder) Username(username string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.username = username
	return connBuilder
}

// Password : Set the password field for `Connection` instance
func (connBuilder *ConnectionBuilder) Password(password string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.password = password
	return connBuilder
}

// Insecure : Set the insecure field for `Connection` instance
func (connBuilder *ConnectionBuilder) Insecure(insecure bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.insecure = insecure
	return connBuilder
}

// Timeout : Set the timeout field for `Connection` instance
func (connBuilder *ConnectionBuilder) Timeout(timeout time.Duration) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.timeout = timeout
	return connBuilder
}

// CAFile : Set the caFile field for `Connection` instance
func (connBuilder *ConnectionBuilder) CAFile(caFilePath string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.caFile = caFilePath
	return connBuilder
}

// Kerberos : Set the kerberos field for `Connection` instance
func (connBuilder *ConnectionBuilder) Kerberos(kerbros bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	// TODO: kerbros==true is not implemented
	if kerbros == true {
		connBuilder.err = errors.New("Kerberos is not currently implemented")
		return connBuilder
	}
	connBuilder.conn.kerberos = kerbros
	return connBuilder
}

// Compress : Set the compress field for `Connection` instance
func (connBuilder *ConnectionBuilder) Compress(compress bool) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.compress = compress
	return connBuilder
}

// Build : Contruct the `Connection` instance
func (connBuilder *ConnectionBuilder) Build() (*Connection, error) {
	// If already has errors, just return
	if connBuilder.err != nil {
		return nil, connBuilder.err
	}

	// Check parameters
	if connBuilder.conn.url == nil {
		return nil, errors.New("The URL must not be empty")
	}
	if len(connBuilder.conn.username) == 0 {
		return nil, errors.New("The Username must not be empty")
	}
	if len(connBuilder.conn.password) == 0 {
		return nil, errors.New("The Password must not be empty")
	}
	// Construct http.Client
	var tlsConfig *tls.Config
	if connBuilder.conn.url.Scheme == "https" {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: connBuilder.conn.insecure,
		}
		if len(connBuilder.conn.caFile) > 0 {
			// Check if the CA File specified exists.
			if _, err := os.Stat(connBuilder.conn.caFile); os.IsNotExist(err) {
				return nil, fmt.Errorf("The ca file '%s' doesn't exist", connBuilder.conn.caFile)
			}
			pool := x509.NewCertPool()
			caCerts, err := ioutil.ReadFile(connBuilder.conn.caFile)
			if err != nil {
				return nil, err
			}
			if ok := pool.AppendCertsFromPEM(caCerts); ok == false {
				return nil, fmt.Errorf("Failed to parse CA Certificate in file '%s'", connBuilder.conn.caFile)
			}
			tlsConfig.RootCAs = pool
		}
	}
	connBuilder.conn.client = &http.Client{
		Timeout: connBuilder.conn.timeout,
		Transport: &http.Transport{
			DisableCompression: !connBuilder.conn.compress,
			TLSClientConfig:    tlsConfig,
		},
	}
	return connBuilder.conn, nil
}

type ssoResponseJSONParent struct {
	children []ssoResponseJson
}

type ssoResponseJson struct {
	accessToken  string `json:"access_token"`
	ssoError     string `json:"error"`
	ssoErrorCode string `json:"error_code"`
}

// Execute a get request to the SSO server and return the response.
func (c *Connection) getSsoResponse(inputRawURL string, parameters map[string]string) (ssoResponseJson, error) {
	useURL, err := url.Parse(inputRawURL)
	if err != nil {
		return ssoResponseJson{}, err
	}
	// Configure TLS parameters:
	var tlsConfig *tls.Config
	if useURL.Scheme == "https" {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: c.insecure,
		}
		if len(c.caFile) > 0 {
			// Check if the CA File specified exists.
			if _, err := os.Stat(c.caFile); os.IsNotExist(err) {
				return ssoResponseJson{}, fmt.Errorf("The CA File '%s' doesn't exist", c.caFile)
			}
			pool := x509.NewCertPool()
			caCerts, err := ioutil.ReadFile(c.caFile)
			if err != nil {
				return ssoResponseJson{}, err
			}
			if ok := pool.AppendCertsFromPEM(caCerts); ok == false {
				return ssoResponseJson{}, fmt.Errorf("Failed to parse CA Certificate in file '%s'", c.caFile)
			}
			tlsConfig.RootCAs = pool
		}
	}

	// Create the HTTP client handle for SSO:
	client := &http.Client{
		Timeout: time.Duration(c.timeout),
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Configure authentication:
	// TODOLATER: implement, skipped for now. kerberos support in Golang seems "interesting"?
	// if c.kerberos == true {
	// }

	// POST request body:
	bodyValues := make(url.Values)
	for k1, v1 := range parameters {
		bodyValues[k1] = []string{v1}
	}

	// Build the net/http request:
	req, err := http.NewRequest("POST", inputRawURL, strings.NewReader(bodyValues.Encode()))
	if err != nil {
		return ssoResponseJson{}, err
	}

	// Add request headers:
	req.Header.Add("User-Agent", fmt.Sprintf("GoSDK/%s", SDK_VERSION))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Send the request and wait for the response:
	resp, err := client.Do(req)
	if err != nil {
		return ssoResponseJson{}, err
	}
	defer resp.Body.Close()

	// Parse and return the JSON response:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ssoResponseJson{}, err
	}
	var json1 ssoResponseJson
	err1 := json.Unmarshal(body, &json1)
	if err1 != nil {
		// Maybe it's array encapsulated, try the other approach.
		var json2 ssoResponseJSONParent
		err2 := json.Unmarshal(body, &json2)
		if err2 != nil {
			return ssoResponseJson{}, fmt.Errorf("Errors for both JSON unmarshal methods (array/non-array) for SSO response: %s / %s", err1.Error(), err2.Error())
		}
		json1.accessToken = json2.children[0].accessToken
		json1.ssoError = json2.children[0].ssoError
		return json1, nil
	}
	return json1, nil
}

// Builds a the URL and parameters to acquire the access token from SSO.
func (c *Connection) buildSsoAuthRequest() (string, map[string]string) {
	// Compute the entry point and the parameters:
	parameters := map[string]string{
		"scope": "ovirt-app-api",
	}

	var entryPoint string
	if c.kerberos == true {
		entryPoint = "token-http-auth"
		parameters["grant_type"] = "urn:ovirt:params:oauth:grant-type:http"
	} else {
		entryPoint = "token"
		parameters["grant_type"] = "password"
		parameters["username"] = c.username
		parameters["password"] = c.password
	}

	// Compute the URL:
	ssoUrl := c.url
	ssoUrl.Path = fmt.Sprintf("/ovirt-engine/sso/oauth/%s", entryPoint)

	// Return the URL and the parameters:
	return ssoUrl.String(), parameters
}

// Builds a the URL and parameters to revoke the SSO access token.
// string = the URL of the SSO service
// map = hash containing the parameters required to perform the revoke
func (c *Connection) buildSsoRevokeRequest() (string, map[string]string) {
	// Compute the parameters:
	parameters := map[string]string{
		"scope": "",
		"token": c.token,
	}

	// Compute the URL:
	ssoUrl := c.url
	ssoUrl.Path = "/ovirt-engine/services/sso-logout"

	// Return the URL and the parameters:
	return ssoUrl.String(), parameters
}

// Tests the connectivity with the server. If connectivity works correctly it returns a nil error. If there is any
// connectivity problem it will return an error containing the reason as the message.
func (c *Connection) Test() error {
	return nil
}

// Performs the authentication process and returns the authentication token. Usually there is no need to
// call this method, as authentication is performed automatically when needed. But in some situations it
// may be useful to perform authentication explicitly, and then use the obtained token to create other
// connections, using the `token` parameter of the constructor instead of the user name and password.
func (c *Connection) Authenticate() {
	c.token, _ = c.getAccessToken()
}

// Indicates if the given object is a link. An object is a link if it has an `href` attribute.
func (c *Connection) IsLink(object string) bool {
	// TODO: implement
	return false
}

// Follows the `href` attribute of the given object, retrieves the target object and returns it.
func (c *Connection) FollowLink(object string) error {
	// TODO: implement
	return nil
}

// Close : Releases the resources used by this connection.
func (c *Connection) Close() {
	if len(c.token) > 0 {
		c.revokeAccessToken()
	}
}

// Builds a request URL from a path, and the set of query parameters.
func (c *Connection) buildRawURL(path string, query map[string]string) string {
	rawURL := fmt.Sprintf("%s%s", c.url.String(), path)
	if len(query) > 0 {
		values := make(url.Values)
		for k, v := range query {
			values[k] = []string{v}
		}
		rawURL = fmt.Sprintf("%s?%s", rawURL, values.Encode())
	}
	return rawURL
}
