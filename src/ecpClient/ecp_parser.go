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
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"image/png"
	"image"
)

func (ec *EcpClient) parseApps(res *http.Response) (*[]App, error) {
	var apps AppsResponse
	result, err := ec.parseResponse(res, &apps)
	if err != nil {
		return nil, err
	}
	parsedRes := result.(*AppsResponse)
	return &parsedRes.Apps, err
}

func (ec *EcpClient) parsePlayer(res *http.Response) (*Player, error) {
	var player Player
	result, err := ec.parseResponse(res, &player)
	if err != nil {
		return nil, err
	}
	parsedRes := result.(*Player)
	return parsedRes, err
}

func (ec *EcpClient) parseActiveApp(res *http.Response) (*App, error) {
	var app ActiveAppResponse
	result, err := ec.parseResponse(res, &app)
	if err != nil {
		return nil, err
	}
	parsedRes := result.(*ActiveAppResponse)
	return &parsedRes.ActiveApp, err
}

func (ec *EcpClient) parseDeviceInfo(res *http.Response) (*DeviceInfo, error) {
	var responseStructure DeviceInfo
	result, err := ec.parseResponse(res, &responseStructure)
	if err != nil {
		return nil, err
	}
	parsedRes := result.(*DeviceInfo)
	return parsedRes, err
}

func (ec *EcpClient) parseAppUi(res *http.Response) (*Node, error) {
	var responseStructure Node
	result, err := ec.parseResponse(res, &responseStructure)
	if err != nil {
		return nil, err
	}
	parsedRes := result.(*Node)
	return parsedRes, err
}

func (ec *EcpClient) parseImage(res *http.Response) (image.Image, error) {
	result, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}


func (ec *EcpClient) parseAppUiSource(res *http.Response) ([]byte, error) {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ec *EcpClient) parseResponse(res *http.Response, v interface{}) (interface{}, error) {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}