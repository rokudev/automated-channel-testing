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

var validAppsResponse = `
{
	"sessionId": "test",
	"status": 0,
	"value": [
		{
			"Title": "test",
			"ID": "test",
			"Type": "test",
			"Version": "test",
			"Subtype": "test"
		},
		{
			"Title": "test2",
			"ID": "test2",
			"Type": "test2",
			"Version": "test2",
			"Subtype": "test2"
		}
	]
}`

var validAppResponse = `
	{
		"sessionId": "test",
		"status": 0,
		"value": {
				"Title": "test",
				"ID": "test",
				"Type": "test",
				"Version": "test",
				"Subtype": "test"
			}
	}`

var validResponseWithNullValue = `
{
	"sessionId": "test",
	"status": 0,
	"value": null
}`

var validSessionResponse = `
{
	"sessionId":"test", 
	"status":0, 
	"value": {
		"country":"", 
		"ip":"", 
		"language":"",
		"modelName":"", 
		"vendorName":"",
		"pressDelay": 0,
		"timeout": 0
	}
}`

var validSessionsResponse = `[
{
	"sessionId":"test", 
	"status":0, 
	"value": {
		"country":"", 
		"ip":"", 
		"language":"",
		"modelName":"", 
		"vendorName":"",
		"pressDelay": 0,
		"timeout": 0
	}
}]`

var validElementResponse = `
{
"sessionId":"test", 
"status":0, 
"value": {
	"XMLName": {
		"Local": "Poster",
		"Space": ""
	},
	"Attrs": [
		{
			"Name": {
				"Local": "bounds",
				"Space": ""
			},
			"Value": "{0, 0, 1920, 1080}"
		},
		{
			"Name": {
				"Local": "index",
				"Space": ""
			},
			"Value": "0"
		},
		{
			"Name": {
				"Local": "loadStatus",
				"Space": ""
			},
			"Value": "3"
		}
	],
	"Nodes": null
	}
}`

var validElementsResponse = `
{
"sessionId":"test", 
"status":0, 
"value": [{
	"XMLName": {
		"Local": "Poster",
		"Space": ""
	},
	"Attrs": [
		{
			"Name": {
				"Local": "bounds",
				"Space": ""
			},
			"Value": "{0, 0, 1920, 1080}"
		},
		{
			"Name": {
				"Local": "index",
				"Space": ""
			},
			"Value": "0"
		},
		{
			"Name": {
				"Local": "loadStatus",
				"Space": ""
			},
			"Value": "3"
		}
	],
	"Nodes": null
	}]
}`