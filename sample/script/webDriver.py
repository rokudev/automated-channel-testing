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
    def __init__(self, roku_ip_address: str):
        data = {'ip' : roku_ip_address}
        request_url = self._build_request_url('')
        response = self._post(request_url, data)
        res = json.loads(response.text)
        self._session_id = res['sessionId']

    def _send_launch_channel(self, channel_code: str):
        data = {'channelId' : channel_code}
        request_url = self._build_request_url(f"/{self._session_id}/launch")
        return self._post(request_url, data)
    
    def _send_sequence(self, sequence):
        data = {'button_sequence' : sequence}
        request_url = self._build_request_url(f"/{self._session_id}/press")
        return self._post(request_url, data)

    def _get_ui_element(self, data: object):
        request_url = self._build_request_url(f"/{self._session_id}/element")
        return self._post(request_url, data)

    def _send_keypress(self, key_press: str):
        data = {'button' : key_press}
        request_url = self._build_request_url(f"/{self._session_id}/press")
        return self._post(request_url, data)

    def _build_request_url(self, endpoint: str):
        return f"http://localhost:9000/v1/session{endpoint}"

    def quiet(self):
        request_url = self._build_request_url(f"/{self._session_id}")
        self._delete(request_url)

    def _post(self, request_url: str, data: object):
        return requests.post(url = request_url, data = json.dumps(data))

    def _get(self, request_url: str):
        return requests.get(request_url)
    
    def _delete(self, request_url: str):
        return requests.delete(request_url)

    def launch_the_channel(self, channel_code):
        launch_response = self._send_launch_channel(channel_code)
        if launch_response.status_code != 200:
            raise Exception("Wrong launch response code")

    def verify_is_screen_loaded(self, data: object, invoke_error = True, retries = 10):
        while retries > 0:
            ui_layout_response = self._get_ui_element(data)
            if ui_layout_response.status_code != 200:
                retries -= 1
                sleep(1)
            else:
                return True
        if invoke_error == True:
            raise Exception("Can't find element")
        else:       
            return False

    def press_btn(self, key_press: str):
        sleep(2)
        key_press_response = self._send_keypress(key_press)
        if key_press_response.status_code != 200:
            raise Exception("Wrong keypress response code")

    def send_word(self, word: str):
        sleep(2)
        for c in word:
            key_press_response = self._send_keypress(f"LIT_{c}")
            if key_press_response.status_code != 200:
                raise Exception("Wrong keypress response code")

    def send_button_sequence(self, sequence):
        key_press_response = self._send_sequence(sequence)
        if key_press_response.status_code != 200:
            raise Exception("Wrong keypress response code")