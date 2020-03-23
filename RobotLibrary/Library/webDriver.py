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

import requests
import json
from time import sleep

class WebDriver:
    def __init__(self, roku_ip_address: str, timeout: int, pressDelay: int):
        data = {'ip' : roku_ip_address, 'timeout': timeout, 'pressDelay': pressDelay}
        request_url = self._build_request_url('')
        response = self._post(request_url, data)
        res = json.loads(response.text)
        self._session_id = res['sessionId']

    def send_launch_channel(self, channel_code: str, contentID, mediaType):
        data = {'channelId' : channel_code, 'contentId': contentID, 'contentType': mediaType}
        request_url = self._build_request_url(f"/{self._session_id}/launch")
        return self._post(request_url, data)
    
    def send_input_data(self, channelId, contentID, mediaType):
        data = {'channelId': channelId, 'contentId': contentID, 'contentType': mediaType}
        request_url = self._build_request_url(f"/{self._session_id}/input")
        return self._post(request_url, data)

    def get_device_info(self):
        request_url = self._build_request_url(f"/{self._session_id}")
        return self._get(request_url)

    def side_load(self, form):
        request_url = self._build_request_url(f"/{self._session_id}/load")
        return requests.post(request_url, files=form)
    
    def get_player_info(self):
        request_url = self._build_request_url(f"/{self._session_id}/player")
        return self._get(request_url)
        
    def send_install_channel(self, channel_code: str):
        data = {'channelId' : channel_code}
        request_url = self._build_request_url(f"/{self._session_id}/install")
        return self._post(request_url, data)
    
    def send_sequence(self, sequence):
        data = {'button_sequence' : sequence}
        request_url = self._build_request_url(f"/{self._session_id}/press")
        return self._post(request_url, data)

    def get_ui_element(self, data: object):
        request_url = self._build_request_url(f"/{self._session_id}/element")
        return self._post(request_url, data)

    def set_timeouts(self, timeoutType: str, delay: int):
        print(str(delay))
        data = {'type': timeoutType, 'ms': delay}
        request_url = self._build_request_url(f"/{self._session_id}/timeouts")
        return self._post(request_url, data)

    def get_ui_elements(self, data: object):
        request_url = self._build_request_url(f"/{self._session_id}/elements")
        return self._post(request_url, data)
    
    def get_apps(self):
        request_url = self._build_request_url(f"/{self._session_id}/apps")
        return self._get(request_url)
    
    def get_current_app(self):
        request_url = self._build_request_url(f"/{self._session_id}/current_app")
        return self._get(request_url)
    
    def get_screen_source(self):
        request_url = self._build_request_url(f"/{self._session_id}/source")
        return self._get(request_url)

    def send_keypress(self, key_press: str):
        data = {'button' : key_press}
        request_url = self._build_request_url(f"/{self._session_id}/press")
        return self._post(request_url, data)
    
    def get_active_element(self):
        request_url = self._build_request_url(f"/{self._session_id}/element/active")
        return self._post(request_url, {})

    def _build_request_url(self, endpoint: str):
        return f"http://localhost:9000/v1/session{endpoint}"

    def quiet(self):
        request_url = self._build_request_url(f"/{self._session_id}")
        return self._delete(request_url)

    def _post(self, request_url: str, data: object):
        return requests.post(url = request_url, data = json.dumps(data))

    def _get(self, request_url: str):
        return requests.get(request_url)
    
    def _delete(self, request_url: str):
        return requests.delete(request_url)

    