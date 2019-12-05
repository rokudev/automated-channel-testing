########################################################################
# Copyright 2019 Roku, Inc.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.
########################################################################

resp = {
            "sessionId": "test",
            "status": 0,
            "value": {
                "ip": "1.1.0.127"
            }
        }

respWithEmptyValue = {
    "sessionId": "test",
    "status": 0,
    "value": None
}

respWithApps = {
    "sessionId": "test",
    "status": 0,
    "value": [{
        "Title": "TestChannel",
        "ID": "id",
        "Type": "menu",
        "Version": "1.0.0",
        "Subtype": ""

    }]
}

respWithApp = {
    "sessionId": "test",
    "status": 0,
    "value": {
        "Title": "TestChannel",
        "ID": "id",
        "Type": "menu",
        "Version": "1.0.0",
        "Subtype": ""

    }
}

respWithPlayerInfo = {
    "sessionId": "42ff6f0d-2362-5e26-a0c9-800e3390bae0",
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
}

resultPlayerInfo = {
    "sessionId": "42ff6f0d-2362-5e26-a0c9-800e3390bae0",
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
        "Position": 1000,
        "Duration": 20000,
        "IsLive": "",
        "Runtime": "",
        "StreamSegment": {
            "Bitrate": "",
            "MediaSequence": "",
            "SegmentType": "",
            "Time": ""
        }
    }
}

uiElement = {
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
        "Nodes": None
    }

respWithElement = {
    "sessionId": "test",
    "status": 0,
    "value": uiElement
}

respWithElements = {
    "sessionId": "test",
    "status": 0,
    "value":[uiElement]
}

respWithDeviceInfo = {
    "sessionId": "test",
    "status": 0,
    "value": {
        "ip": "1.1.0.127",
        "vendorName": "Roku",
        "modelName": "Roku",
        "language": "en",
        "country": "US"
    }
}

respWithSource = {
    "sessionId": "test",
    "status": 0,
    "value": "dfverv5vrevrf=="
}

errorMessage = "Error occurs"

responseWithError = {
    "sessionId": "test",
    "status": 8,
    "value": {
        "message": errorMessage
    }
}