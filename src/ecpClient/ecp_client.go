///////////////////////////////////////////////////////////////////////////
// Copyright 2019 Roku, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//////////////////////////////////////////////////////////////////////////

package ecpClient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"image"
	"errors"
	"bytes"
	"mime/multipart"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
	"strings"
)

const RequestTimeoutMilliseconds = 30000

const successStatusCode = 200

const requestRetries = 2

var endpointsMap = map[string]string{
	"appUI":     "query/app-ui",
	"activeApp": "query/active-app",
	"apps":      "query/apps",
	"install":   "install/%s",
	"launch":    "launch/%s?contentId=%s&mediaType=%s",
	"icon":      "query/icon/%s",
	"device":    "query/device-info",
	"keypress":  "keypress/%s",
	"keydown":   "keydown/%s",
	"keyup":     "keyup/%s",
	"player":    "query/media-player",
	"input":     "input/%s?contentId=%s&mediaType=%s",
	"load":      "/plugin_install",
}

type BaseClient struct {
	BaseURL    *url.URL
	HttpClient *Client
}

type EcpClient struct {
	*BaseClient
}

type PluginClient struct {
	*BaseClient
}

type RequestData struct {
	Endpoint    *url.URL
	ParamsMap   map[string]string
	HeadersMap  map[string]string
	Method      string
	RequestBody io.Reader
}

const baseUrlStructure = "http://%s:8060"
const pluginBaseUrlStructure = "http://%s"

func GetEcpClient(ip string) (*EcpClient, error) {
	client, err := getClient(baseUrlStructure, ip)
	if err != nil {
		return nil, err
	}
	return &EcpClient{
		BaseClient: client,
	}, nil
}

func GetPluginClient(ip string) (*PluginClient, error) {
	client, err := getClient(pluginBaseUrlStructure, ip)
	if err != nil {
		return nil, err
	}
	return &PluginClient{
		BaseClient: client,
	}, nil
}

func getClient(baseUrlA string, ip string) (*BaseClient, error) {
	baseUrl, err := url.Parse(fmt.Sprintf(baseUrlA, ip))
	if err != nil {
		return nil, err
	}
	client := SetHTTPClient(http.DefaultClient)
	timeout := SetRequestTimeout(RequestTimeoutMilliseconds * time.Millisecond)

	defaultClient := NewClient(client, timeout)
	return &BaseClient{
		BaseURL:    baseUrl,
		HttpClient: defaultClient,
	}, nil
}

func (ec *EcpClient) SetTimeout(timeout time.Duration) {
	SetRequestTimeout(timeout * time.Millisecond)(ec.HttpClient)
}

func (ec *EcpClient) GetTimeout() int {
	ms := int(ec.HttpClient.HttpClient.Timeout/time.Millisecond)
	return ms
}

func (ec *BaseClient) createRequest(data *RequestData) (*http.Request, error) {
	u := ec.BaseURL.ResolveReference(data.Endpoint)
	req, err := http.NewRequest(data.Method, u.String(), data.RequestBody)
	if err != nil {
		return nil, err
	}

	for hName, hValue := range data.HeadersMap {
		req.Header.Set(hName, hValue)
	}

	query := req.URL.Query()
	for pName, pValue := range data.ParamsMap {
		query.Set(pName, pValue)
	}

	req.URL.RawQuery = query.Encode()
	return req, err
}

func (ec *PluginClient)  Load(file io.Reader, user string, pass string) (bool, error) {
	end := endpointsMap["load"]
	auth, err := ec.getAuthHeader(end, user, pass)
	if err != nil {
		return false, err
	}
	
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("mysubmit", "Delete")
	writer.WriteField("archive", "")
	writer.Close()
    headers :=  map[string]string {
		"Authorization": auth,
		"Content-Type":  "application/json",
	}
	_, err = ec.makePluginRequest("POST", end, body.Bytes(), headers)
	if err != nil {
		return false, err
	}

	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("archive", "name")
	io.Copy(part, file)
	writer.WriteField("mysubmit", "Install")
	writer.Close()
	return  ec.makePluginRequest("POST", end, body.Bytes(), headers)
}

func (ec *PluginClient) makePluginRequest(method string, end string, body []byte, headers map[string]string)  (bool, error) {
	res, err := url.Parse(end)
	if err != nil {
		return false, err
	}
	requestObject := &RequestData {
		Method:   method,
		Endpoint: res,
		RequestBody: bytes.NewBuffer(body),
		HeadersMap: headers,
	}
	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return false, err
	} else if response.StatusCode != successStatusCode {
		return false, errors.New("Command execution failed")
	}

	return true, nil
}
func (ec *PluginClient) getAuthHeader(uri string, user string, pass string)  (string, error) {
	method := "POST"
	res,_ := url.Parse(uri)
	requestObject := &RequestData{
		Method:   method,
		Endpoint: res,
	}
	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return "", err
	}
    digestParts := digestParts(response)
    digestParts["uri"] = uri
    digestParts["method"] = method
    digestParts["username"] = user
    digestParts["password"] = pass
    authHeader := getDigestAuthrization(digestParts)
    return authHeader, nil
}

func digestParts(resp *http.Response) map[string]string {
    result := map[string]string{}
    if len(resp.Header["Www-Authenticate"]) > 0 {
        wantedHeaders := []string{"nonce", "realm", "qop"}
        responseHeaders := strings.Split(resp.Header["Www-Authenticate"][0], ",")
        for _, r := range responseHeaders {
            for _, w := range wantedHeaders {
                if strings.Contains(r, w) {
                    result[w] = strings.Split(r, `"`)[1]
                }
            }
        }
    }
    return result
}

func getMD5(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func getCnonce() string {
    b := make([]byte, 8)
    io.ReadFull(rand.Reader, b)
    return fmt.Sprintf("%x", b)[:16]
}

func getDigestAuthrization(digestParts map[string]string) string {
    d := digestParts
    ha1 := getMD5(d["username"] + ":" + d["realm"] + ":" + d["password"])
    ha2 := getMD5(d["method"] + ":" + d["uri"])
    nonceCount := 00000001
    cnonce := getCnonce()
    response := getMD5(fmt.Sprintf("%s:%s:%v:%s:%s:%s", ha1, d["nonce"], nonceCount, cnonce, d["qop"], ha2))
    authorization := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
        d["username"], d["realm"], d["nonce"], d["uri"], cnonce, nonceCount, d["qop"], response)
    return authorization
}

// Do api call with given http.Request
func (ec *BaseClient) call(req *http.Request) (*http.Response, error) {
	retries := requestRetries
	var resp *http.Response
	var err error
	for retries > 0 {
        resp, err = ec.HttpClient.HttpClient.Do(req)
        if err != nil {
            retries --
        } else {
            break
        }
    }
	return resp, err
}

func (ec *BaseClient) makeRequest(data *RequestData) (*http.Response, error) {
	req, err := ec.createRequest(data)
	if err != nil {
		return nil, err
	}

	resp, err := ec.call(req)
	
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ec *EcpClient) GetAppUi() (*Node, error) {
	end, err := url.Parse(endpointsMap["appUI"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}
	return ec.parseAppUi(response)
}

func (ec *EcpClient) GetSource() ([]byte, error) {
	end, err := url.Parse(endpointsMap["appUI"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}

	return ec.parseAppUiSource(response)
}

func (ec *EcpClient) GetActiveApp() (*App, error) {
	end, err := url.Parse(endpointsMap["activeApp"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}
	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}
	return ec.parseActiveApp(response)
}

func (ec *EcpClient) GetApps() (*[]App, error) {
	end, err := url.Parse(endpointsMap["apps"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}
	return ec.parseApps(response)
}

func (ec *EcpClient) GetPlayer() (*Player, error) {
	end, err := url.Parse(endpointsMap["player"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}
	return ec.parsePlayer(response)
}

func (ec *EcpClient) InstallChannel(channelId string) (bool, error) {
	if len(channelId) == 0 {
		return false, errors.New("the channelId is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["install"], channelId))
	if err != nil {
		return false, err
	}
 
	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) LaunchChannel(channelId string, contentId  string , mediaType string ) (bool, error) {
	if len(channelId) == 0 {
		return false, errors.New("the channelId is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["launch"], channelId, contentId, mediaType))
	if err != nil {
		return false, err
	}

	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) InputChannel(channelId string, contentId  string , mediaType string ) (bool, error) {
	if len(contentId) == 0 || len(mediaType) == 0 {
		return false, errors.New("contentId and mediaType are required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["input"], channelId, contentId, mediaType))
	if err != nil {
		return false, err
	}

	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) GetIcon(channelId string) (image.Image, error) {
	if len(channelId) == 0 {
		return nil, errors.New("the channelId is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["icon"], channelId))
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}

	return ec.parseImage(response)
}

func (ec *EcpClient) GetDeviceInfo() (*DeviceInfo, error) {
	end, err := url.Parse(endpointsMap["device"])
	if err != nil {
		return nil, err
	}

	requestObject := &RequestData{
		Method:   "GET",
		Endpoint: end,
	}

	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return nil, err
	}
	return ec.parseDeviceInfo(response)
}

func (ec *EcpClient) KeyPress(button string) (bool, error) {
	if len(button) == 0 {
		return false, errors.New("the button is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["keypress"], button))
	if err != nil {
		return false, err
	}

	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) KeyDown(button string) (bool, error) {
	if len(button) == 0 {
		return false, errors.New("the button is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["keydown"], button))
	if err != nil {
		return false, err
	}

	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) KeyUp(button string) (bool, error) {
	if len(button) == 0 {
		return false, errors.New("the button is required")
	}
	end, err := url.Parse(fmt.Sprintf(endpointsMap["keyup"], button))
	if err != nil {
		return false, err
	}

	return ec.makeNavigationRequest("POST", end)
}

func (ec *EcpClient) makeNavigationRequest(method string, end *url.URL) (bool, error) {
	requestObject := &RequestData {
		Method:   method,
		Endpoint: end,
	}
    
	response, err := ec.makeRequest(requestObject)
	if err != nil {
		return false, err
	} else if response.StatusCode != successStatusCode {
		return false, errors.New("Command execution failed")
	}
	return true, nil
}

