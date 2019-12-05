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
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"encoding/xml"
)

var testRequest = &http.Request{Method: "GET", URL: base, Body: nil}

func TestClient_CallKeyPressNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.KeyPress("errorCommnad")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.IsType(t, new(url.Error), err)
}

func TestClient_CallKeyPressInvalidData(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.KeyPress("")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.Equal(t, err.Error(), "the button is required")
}

func TestClient_CallKeyPressSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.KeyPress("up")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestClient_CallKeyUpNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.KeyUp("errorCommnad")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.IsType(t, new(url.Error), err)
}

func TestClient_CallKeyUpNoInvalidData(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.KeyUp("")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.Equal(t, err.Error(), "the button is required")
}

func TestClient_CallKeyUpSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.KeyUp("up")
	assert.Nil(t, err)
	assert.Equal(t, resp, true)
}

func TestClient_CallKeyDownNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.KeyDown("errorCommnad")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.IsType(t, new(url.Error), err)
}

func TestClient_CallKeyDownInvalidData(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.KeyDown("")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.Equal(t, err.Error(), "the button is required")
}

func TestClient_CallKeyDownSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.KeyDown("up")
	assert.Nil(t, err)
	assert.Equal(t, resp, true)
}

func TestClient_CallLaunchNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.LaunchChannel("11111", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.IsType(t, new(url.Error), err)
}

func TestClient_CallLaunchInvalidData(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.LaunchChannel("", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.Equal(t, err.Error(), "the channelId is required")
}

func TestClient_CallLaunchSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.LaunchChannel("dev", "", "")
	assert.Nil(t, err)
	assert.Equal(t, resp, true)
}

func TestClient_CallInstallNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.InstallChannel("11111")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.IsType(t, new(url.Error), err)
}

func TestClient_CallInstallInvalidData(t *testing.T) {
	httpClient := GetMockedClient(nil)

	resp, err := httpClient.InstallChannel("")
	assert.NotNil(t, err)
	assert.Equal(t, resp, false)
	assert.Equal(t, err.Error(), "the channelId is required")
}

func TestClient_CallInstallSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&SuccessResponseMock)

	resp, err := httpClient.InstallChannel("dev")
	assert.Nil(t, err)
	assert.Equal(t, resp, true)
}

func TestClient_CallActiveAppSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&AppResponseMock)
	resp, err := httpClient.GetActiveApp()
	expectedResult := &App{
		Title: "test",
		ID: "test",
		Type: "test",
		Version: "test",
		Subtype: "test",
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestClient_CallActiveAppNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)
	resp, err := httpClient.GetActiveApp()
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestClient_CallAppsSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&AppsResponseMock)
	resp, err := httpClient.GetApps()
	expectedResult := &[]App{
		App{
			Title: "test",
			ID: "test",
			Type: "test",
			Version: "test",
			Subtype: "test",
		},
		App{
			Title: "test2",
			ID: "test2",
			Type: "test2",
			Version: "test2",
			Subtype: "test2",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestClient_CallAppsNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)
	resp, err := httpClient.GetApps()
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestClient_CallDeviceInfoSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&deviceInfoResponse)
	resp, err := httpClient.GetDeviceInfo()
	expectedResult := &DeviceInfo{
	Udn: "test",
	SerialNumber: "test",
	DeviceID: "test",
	AdvertisingID: "test",
	VendorName: "test",
	ModelName: "test",
	ModelNumber: "test",
	ModelRegion: "test",
	IsTv: "test",
	IsStick: "test",
	SupportsEthernet: "test",
	WifiMac: "test",
	WifiDriver: "test",
	EthernetMac: "test",
	NetworkType: "test",
	NetworkName: "test",
	FriendlyDeviceName: "test",
	FriendlyModelName: "test",
	DefaultDeviceName: "test",
	UserDeviceName: "test",
	UserDeviceLocation: "test",
	BuildNumber: "test",
	SoftwareVersion: "test",
	SoftwareBuild: "test",
	SecureDevice: "test",
	Language: "test",
	Country: "test",
	Locale: "test",
	TimeZoneAuto: "test",
	TimeZone: "test",
	TimeZoneName: "test",
	TimeZoneTz: "test",
	TimeZoneOffset: "test",
	ClockFormat: "test",
	Uptime: "test",
	PowerMode: "test",
	SupportsSuspend: "test",
	SupportsFindRemote: "test",
	FindRemoteIsPossible: "test",
	SupportsAudioGuide: "test",
	SupportsRva: "test",
	DeveloperEnabled: "test",
	KeyedDeveloperID: "test",
	SearchEnabled: "test",
	SearchChannelsEnabled: "test",
	NotificationsEnabled: "test",
	NotificationsFirstUse: "test",
	SupportsPrivateListening: "test",
	HeadphonesConnected: "test",
	SupportsEcsTextedit: "test",
	SupportsEcsMicrophone: "test", 
	SupportsWakeOnWlan: "test",
	HasPlayOnRoku: "test",
	HasMobileScreensaver: "test",
	SupportURL: "test",
	GrandcentralVersion: "test",
	TrcVersion: "test",
	TrcChannelVersion: "test",
	DavinciVersion: "test",
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestClient_CallDeviceInfoNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)
	resp, err := httpClient.GetDeviceInfo()
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestClient_CallAppUiSuccessResponse(t *testing.T) {
	httpClient := GetMockedClient(&UiResponseMock)
	resp, err := httpClient.GetAppUi()
	expectedResult :=  &Node{
			XMLName: xml.Name{
				Local: "MainScene",
			},
			Attrs: []xml.Attr{
				{
					Name: xml.Name{
						Local: "bounds",
					},
					Value: "{0, 0, 1920, 1080}",
				},
				{
					Name: xml.Name{
						Local: "children",
					},
					Value: "0",
				},
				{
					Name: xml.Name{
						Local: "extends",
					},
					Value: "BaseScene",
				},
				{
					Name: xml.Name{
						Local: "focusable",
					},
					Value: "true",
				},
			},
			Nodes: []Node{
				{
					XMLName: xml.Name{
						Local: "Poster",
					},
					Attrs: []xml.Attr{
						{
							Name: xml.Name{
								Local: "bounds",
							},
							Value: "{0, 0, 1920, 1080}",
						},
						{
							Name: xml.Name{
								Local: "index",
							},
							Value: "0",
						},
						{
							Name: xml.Name{
								Local: "loadStatus",
							},
							Value: "3",
						},
					},
					Nodes: nil,
				},
			},
		}
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestClient_CallAppUiNoResponse(t *testing.T) {
	httpClient := GetMockedClient(nil)
	resp, err := httpClient.GetAppUi()
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

