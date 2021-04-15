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

import unittest
from unittest.mock import patch
import sys
import json
from Library.webDriver import WebDriver
from response_mock import *

class BasicTests(unittest.TestCase):

    def setUp(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(resp)
            self._webDriver = WebDriver("1.1.0.127", 5000, 2000)


    def test_send_keypress(self):
        with patch('Library.webDriver.requests.post') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.send_keypress("select")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)
    
    def test_send_launch_channel(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.send_launch_channel("dev", "", "")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

    def test_send_install_channel(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.send_install_channel("1111")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

    def test_get_device_info(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithDeviceInfo)
            response = self._webDriver.get_device_info()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithDeviceInfo)

    def test_send_sequence(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.send_sequence(["up", "left"])
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

    def test_get_apps(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithApps)
            response = self._webDriver.get_apps()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithApps)

    def test_get_player(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithPlayerInfo)
            response = self._webDriver.get_player_info()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithPlayerInfo)
    
    def test_get_current_app(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithApp)
            response = self._webDriver.get_current_app()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithApp)

    def test_quiet(self):
        with patch('Library.webDriver.requests.delete') as mock_delete:
            mock_delete.return_value.status_code = 200
            mock_delete.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.quiet()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)
    
    def test_get_ui_element(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)
            response = self._webDriver.get_ui_element([{"using": "text", "value": "text"}])
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithElement)
    
    def test_get_ui_elements(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElements)
            response = self._webDriver.get_ui_elements([{"using": "text", "value": "text"}])
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithElements)

    def test_get_active_element(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)
            response = self._webDriver.get_active_element()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithElement)

    def test_set_timeouts(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.set_timeouts("implicit", "2000")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

    def test_get_screen_source(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithSource)
            response = self._webDriver.get_screen_source()
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithSource)

    def test_send_input_data(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            response = self._webDriver.send_input_data("dev", "12", "movie")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

    def test_side_load(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)
            multipart_form_data = {
                'username': (None, "user"),
                'password': (None, "pass")
            }
            response = self._webDriver.side_load(multipart_form_data)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(json.loads(response.text), respWithEmptyValue)

