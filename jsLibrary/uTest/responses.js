///////////////////////////////////////////////////////////////////////////
// Copyright 2020 Roku, Inc.
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

const uiElement = {
    "XMLName": {
        "Space": "",
        "Local": "Label"
    },
    "Attrs": [
        {
            "Name": {
                "Space": "",
                "Local": "bounds"
            },
            "Value": "{1, 1, 1, 1}"
        },
        {
            "Name": {
                "Space": "",
                "Local": "color"
            },
            "Value": "#ffffff6f"
        },
        {
            "Name": {
                "Space": "",
                "Local": "index"
            },
            "Value": "0"
        },
        {
            "Name": {
                "Space": "",
                "Local": "text"
            },
            "Value": "text"
        }
    ],
    "Nodes": null
}

module.exports = {
    uiElement: uiElement,
    respWithEmptyValue: {
        "sessionId": "test",
        "status": 0,
        "value": null
    },
    requestLaunch: {
        'channelId': /.+/i, 
        'contentId': '', 
        'contentType': ''
    },
    requestInstall: {
        'channelId': /.+/i
    },
    resp: {
        'sessionId': '11111', 
        'status': 0, 
        'value': {
            'ip': '127.0.0.1', 
            'timeout': 20000, 
            'pressDelay': 2000
        }
    },
    requestPress: {
        'button': /.+/i
    },
    requestSequence: {
        'button_sequence': [/.+/i, /.+/i]
    },
    respWithDeviceInfo: {
        "sessionId": "11111",
        "status": 0,
        "value": {
            "ip": "1.1.0.127",
            "vendorName": "Roku",
            "modelName": "Roku",
            "language": "en",
            "country": "US"
        }
    },
    respWithApps: {
        "sessionId": "11111",
        "status": 0,
        "value": [{
            "Title": "TestChannel",
            "ID": "id",
            "Type": "menu",
            "Version": "1.0.0",
            "Subtype": ""
    
        }]
    },
    respWithApp: {
        "sessionId": "11111",
        "status": 0,
        "value": {
            "Title": "TestChannel",
            "ID": "id",
            "Type": "menu",
            "Version": "1.0.0",
            "Subtype": ""
    
        }
    },
    respWithPlayerInfo: {
        "sessionId": "11111",
        "status": 0,
        "value": {
            "Error": "false",
            "State": "play",
            "Format": {
                "Audio": "",
                "Captions": "",
                "Container": "",
                "Drm": "",
                "Video": "",
                "VideoRes": ""
            },
            "Buffering": {
                "Current": "",
                "Max": "",
                "Target": ""
            },
            "NewStream": {
                "Speed": ""
            },
            "Position": "1000 ms",
            "Duration": "20000 ms",
            "IsLive": "",
            "Runtime": "",
            "StreamSegment": {
                "Bitrate": "",
                "MediaSequence": "",
                "SegmentType": "",
                "Time": ""
            }
        }
    },
    respWithElement: {
        "sessionId": "11111",
        "status": 0,
        "value": uiElement
    },
    respWithElements: {
        "sessionId": "11111",
        "status": 0,
        "value":[uiElement]
    },
    requestElement: {
        "elementData" :[{
            "using":  /.+/i,
            "value":  /.+/i
        }]
    },
    requestTimeout: {
        'type': /.+/i, 
        'ms': /.+/i
    },
    respWithSource: {
        "sessionId": "11111",
        "status": 0,
        "value": "dfverv5vrevrf=="
    },
    responseWithError: {
        "sessionId": "11111",
        "status": 8,
        "value": {
            "message": "Error"
        }
    }
}