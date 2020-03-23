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

package httpServer

type BuildInfo struct {
	Version string `json:"version"`
	Time string `json:"time"`
}

type OsInfo struct {
	Arch string `json:"arch"`
	Name string `json:"name"`
}

type Status struct {
	Build BuildInfo `json:"build"`
	Os OsInfo `json:"os"`
}

type Capability struct {
	Ip string `json:"ip"`
	Timeout int `json:"timeout"`
	PressDelay int `json:"pressDelay"`
	VendorName string   `json:"vendorName"`
	ModelName string   `json:"modelName"`
	Language  string   `json:"language"`
	Country   string   `json:"country"`
}

type SessionResponse struct {
	Id string `json:"sessionId"`
	Status int `json:"status"`
	Value interface{} `json:"value"`
}

var responseStatuses = map[string]int {
	"Success": 0,
	"NoSuchDriver": 6,
	"NoSuchElement": 7,
	"UnknownCommand": 9,
	"UnknownError": 13,
	"SessionNotCreatedException": 33,
}

type ButtonRequest struct {
	Button string `json:"button"`
	Button_sequence []string `json:"button_sequence"`
	Button_delays []string `json:"button_delays"`
}

type Element struct {
	Using string `json:"using"`
	Value string `json:"value"`
	Attribute string `json:"attribute"`
}

type ElementRequest struct {
	ParentData []Element `json:"parentData"`
	ElementData []Element `json:"elementData"`
} 

type ChannelRequest struct { 
	ChannelId string `json:"channelId"`
	ContentType string `json:"contentType"`
	ContentId string `json:"contentId"`
}

type TimeoutRequest struct {
	Type string `json:"type"`
	Ms int `json:"ms"`
}

type IntentRequest struct {
	Intent string `json:"intent"`
}