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
	"encoding/xml"
)
type AppsResponse struct {
	Apps   []App `xml:"app"`
} 

type App struct {
    Title   string `xml:",chardata"`
    ID      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
	Subtype string `xml:"subtype,attr"`
}

type Player struct {
	Error   string   `xml:"error,attr"`
	State   string   `xml:"state,attr"`
	Format struct {
		Audio     string `xml:"audio,attr"`
		Captions  string `xml:"captions,attr"`
		Container string `xml:"container,attr"`
		Drm       string `xml:"drm,attr"`
		Video     string `xml:"video,attr"`
		VideoRes  string `xml:"video_res,attr"`
	} `xml:"format"`
	Buffering struct {
		Current string `xml:"current,attr"`
		Max     string `xml:"max,attr"`
		Target  string `xml:"target,attr"`
	} `xml:"buffering"`
	NewStream struct {
		Speed string `xml:"speed,attr"`
	} `xml:"new_stream"`
	Position      string `xml:"position"`
	Duration      string `xml:"duration"`
	IsLive        string `xml:"is_live"`
	Runtime       string `xml:"runtime"`
	StreamSegment struct {
		Bitrate       string `xml:"bitrate,attr"`
		MediaSequence string `xml:"media_sequence,attr"`
		SegmentType   string `xml:"segment_type,attr"`
		Time          string `xml:"time,attr"`
	} `xml:"stream_segment"`
} 

type ActiveAppResponse struct {
	ActiveApp App `xml:"app"`
}

type Node struct {
	XMLName xml.Name
    Attrs   []xml.Attr `xml:",any,attr"`
    Nodes   []Node     `xml:",any"`
}

type DeviceInfo struct {
	Udn                      string   `xml:"udn"`
	SerialNumber             string   `xml:"serial-number"`
	DeviceID                 string   `xml:"device-id"`
	AdvertisingID            string   `xml:"advertising-id"`
	VendorName               string   `xml:"vendor-name"`
	ModelName                string   `xml:"model-name"`
	ModelNumber              string   `xml:"model-number"`
	ModelRegion              string   `xml:"model-region"`
	IsTv                     string   `xml:"is-tv"`
	IsStick                  string   `xml:"is-stick"`
	SupportsEthernet         string   `xml:"supports-ethernet"`
	WifiMac                  string   `xml:"wifi-mac"`
	WifiDriver               string   `xml:"wifi-driver"`
	EthernetMac              string   `xml:"ethernet-mac"`
	NetworkType              string   `xml:"network-type"`
	NetworkName              string   `xml:"network-name"`
	FriendlyDeviceName       string   `xml:"friendly-device-name"`
	FriendlyModelName        string   `xml:"friendly-model-name"`
	DefaultDeviceName        string   `xml:"default-device-name"`
	UserDeviceName           string   `xml:"user-device-name"`
	UserDeviceLocation       string   `xml:"user-device-location"`
	BuildNumber              string   `xml:"build-number"`
	SoftwareVersion          string   `xml:"software-version"`
	SoftwareBuild            string   `xml:"software-build"`
	SecureDevice             string   `xml:"secure-device"`
	Language                 string   `xml:"language"`
	Country                  string   `xml:"country"`
	Locale                   string   `xml:"locale"`
	TimeZoneAuto             string   `xml:"time-zone-auto"`
	TimeZone                 string   `xml:"time-zone"`
	TimeZoneName             string   `xml:"time-zone-name"`
	TimeZoneTz               string   `xml:"time-zone-tz"`
	TimeZoneOffset           string   `xml:"time-zone-offset"`
	ClockFormat              string   `xml:"clock-format"`
	Uptime                   string   `xml:"uptime"`
	PowerMode                string   `xml:"power-mode"`
	SupportsSuspend          string   `xml:"supports-suspend"`
	SupportsFindRemote       string   `xml:"supports-find-remote"`
	FindRemoteIsPossible     string   `xml:"find-remote-is-possible"`
	SupportsAudioGuide       string   `xml:"supports-audio-guide"`
	SupportsRva              string   `xml:"supports-rva"`
	DeveloperEnabled         string   `xml:"developer-enabled"`
	KeyedDeveloperID         string   `xml:"keyed-developer-id"`
	SearchEnabled            string   `xml:"search-enabled"`
	SearchChannelsEnabled    string   `xml:"search-channels-enabled"`
	NotificationsEnabled     string   `xml:"notifications-enabled"`
	NotificationsFirstUse    string   `xml:"notifications-first-use"`
	SupportsPrivateListening string   `xml:"supports-private-listening"`
	HeadphonesConnected      string   `xml:"headphones-connected"`
	SupportsEcsTextedit      string   `xml:"supports-ecs-textedit"`
	SupportsEcsMicrophone    string   `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan       string   `xml:"supports-wake-on-wlan"`
	HasPlayOnRoku            string   `xml:"has-play-on-roku"`
	HasMobileScreensaver     string   `xml:"has-mobile-screensaver"`
	SupportURL               string   `xml:"support-url"`
	GrandcentralVersion      string   `xml:"grandcentral-version"`
	TrcVersion               string   `xml:"trc-version"`
	TrcChannelVersion        string   `xml:"trc-channel-version"`
	DavinciVersion           string   `xml:"davinci-version"`
} 
