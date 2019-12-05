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
)

const RequestTimeoutMilliseconds = 30000

const successStatusCode = 200

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
}

type EcpClient struct {
	BaseURL    *url.URL
	HttpClient *Client
}

type RequestData struct {
	Endpoint    *url.URL
	ParamsMap   map[string]string
	HeadersMap  map[string]string
	Method      string
	RequestBody io.Reader
}

const baseUrlStructure = "http://%s:8060"

func GetEcpClient(ip string) (*EcpClient, error) {
	baseUrl, err := url.Parse(fmt.Sprintf(baseUrlStructure, ip))
	if err != nil {
		return nil, err
	}

	client := SetHTTPClient(http.DefaultClient)
	timeout := SetRequestTimeout(RequestTimeoutMilliseconds * time.Millisecond)

	defaultClient := NewClient(client, timeout)
	return &EcpClient{
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

func (ec *EcpClient) createRequest(data *RequestData) (*http.Request, error) {
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

// Do api call with given http.Request
func (ec *EcpClient) call(req *http.Request) (*http.Response, error) {
	resp, err := ec.HttpClient.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ec *EcpClient) makeRequest(data *RequestData) (*http.Response, error) {
	req, err := ec.createRequest(data)
	if err != nil {
		return nil, err
	}

	return ec.call(req)
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
	requestObject := &RequestData{
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
