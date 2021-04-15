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
from Library.RobotLibrary import RobotLibrary
from response_mock import *

class LibTests(unittest.TestCase):
    def setUp(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(resp)
            self._library = RobotLibrary("1.1.0.127")

    def test_press_btn_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.pressBtn("select")
        a = str(context.exception)
        self.assertTrue(errorMessage in str(context.exception))

    def test_press_btn_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.pressBtn("select")
            except Exception:
                self.fail("pressBtn() raised Exception unexpectedly!")
    
    def test_launch_channel_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.launchTheChannel("dev")
        self.assertTrue(errorMessage in str(context.exception))

    def test_launch_channel_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.launchTheChannel("dev")
            except Exception:
                self.fail("launchTheChannel() raised Exception unexpectedly!")

    def test_get_apps_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.getApps()
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_apps_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithApps)  
            try:
                response = self._library.getApps()
                self.assertEqual(response, respWithApps['value'])
            except Exception:
                self.fail("getApps() raised Exception unexpectedly!")

    def test_verify_is_screen_loaded_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.verifyIsScreenLoaded([{"using": "text", "value": "text"}], True, 1)
        self.assertTrue('Can\'t find element' in str(context.exception))

    def test_verify_is_screen_loaded_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)  
            try:
                response = self._library.verifyIsScreenLoaded([{"using": "text", "value": "text"}], True, 1)
                self.assertEqual(response, True)
            except Exception:
                self.fail("verifyIsScreenLoaded() raised Exception unexpectedly!")

    def test_send_word_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.sendWord("text")
        self.assertTrue(errorMessage in str(context.exception))

    def test_send_word_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.sendWord("text")
            except Exception:
                self.fail("sendWord() raised Exception unexpectedly!")

    def test_send_button_sequence_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.sendButtonSequence(["up", "up"])
        self.assertTrue(errorMessage in str(context.exception))

    def test_send_button_sequence_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.sendButtonSequence(["up", "up"])
            except Exception:
                self.fail("send_button_sequence() raised Exception unexpectedly!")

    def test_get_element_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.getElement([{"using": "text", "value": "text"}])
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_element_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)  
            try:
                response = self._library.getElement([{"using": "text", "value": "text"}])
                self.assertEqual(response, respWithElement['value'])
            except Exception:
                self.fail("getElement() raised Exception unexpectedly!")
    
    def test_get_elements_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.getElements([{"using": "text", "value": "text"}])
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_elements_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)  
            try:
                response = self._library.getElements([{"using": "text", "value": "text"}])
                self.assertEqual(response, respWithElement['value'])
            except Exception:
                self.fail("getElements() raised Exception unexpectedly!")
    
    def test_get_focused_element_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.getFocusedElement()
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_focused_element_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithElement)  
            try:
                response = self._library.getFocusedElement()
                self.assertEqual(response, respWithElement['value'])
            except Exception:
                self.fail("getFocusedElement() raised Exception unexpectedly!")

    def test_verify_is_channel_loaded_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.verifyIsChannelLoaded("id", 1)
        self.assertTrue(errorMessage in str(context.exception))

    def test_verify_is_channel_loaded_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithApp)  
            try:
                response = self._library.verifyIsChannelLoaded("id", 1)
                self.assertEqual(response, True)
            except Exception:
                self.fail("verifyIsChannelLoaded() raised Exception unexpectedly!")

    def test_get_current_channel_info_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.getCurrentChannelInfo()
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_current_channel_info_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithApp)  
            try:
                response = self._library.getCurrentChannelInfo()
                self.assertEqual(response, respWithApp['value'])
            except Exception:
                self.fail("getCurrentChannelInfo() raised Exception unexpectedly!")

    def test_get_device_info_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.getDeviceInfo()
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_device_info_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithDeviceInfo)  
            try:
                response = self._library.getDeviceInfo()
                self.assertEqual(response, respWithDeviceInfo['value'])
            except Exception:
                self.fail("getDeviceInfo() raised Exception unexpectedly!")


    def test_get_player_info_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.getPlayerInfo()
        self.assertTrue(errorMessage in str(context.exception))

    def test_get_player_info_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithPlayerInfo)  
            try:
                response = self._library.getPlayerInfo()
                self.assertEqual(response, resultPlayerInfo['value'])
            except Exception:
                self.fail("getPlayerInfo() raised Exception unexpectedly!")

    def test_verify_is_playback_started_error(self):
        with patch('Library.webDriver.requests.get') as mock_get, self.assertRaises(Exception) as context:
            mock_get.return_value.status_code = 500
            mock_get.return_value.text = json.dumps(responseWithError)
            self._library.verifyIsPlaybackStarted(1)
        self.assertTrue('Invalid player state' in str(context.exception))

    def test_verify_is_playback_started_success(self):
        with patch('Library.webDriver.requests.get') as mock_get:
            mock_get.return_value.status_code = 200
            mock_get.return_value.text = json.dumps(respWithPlayerInfo)  
            try:
                self._library.verifyIsPlaybackStarted(1)
            except Exception:
                self.fail("verifyIsPlaybackStarted() raised Exception unexpectedly!")

    def test_get_attribute_success(self):
        res = self._library.getAttribute(uiElement, "color")
        self.assertTrue('#ffffff6f' in str(res))

    def test_get_attribute_error(self):
        with  self.assertRaises(Exception) as context:
            self._library.getAttribute(uiElement, "attrNotExist")
        self.assertTrue('Can\'t find attribute' in str(context.exception))

        
    def test_input_deep_linking_data_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.inputDeepLinkingData("dev", "12", "movie")
        self.assertTrue(errorMessage in str(context.exception))


    def test_input_deep_linking_data_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.inputDeepLinkingData("dev", "12", "movie")
            except Exception:
                self.fail("inputDeepLinkingData() raised Exception unexpectedly!")

    def test_side_load_error(self):
        with patch('Library.webDriver.requests.post') as mock_post, self.assertRaises(Exception) as context:
            mock_post.return_value.status_code = 500
            mock_post.return_value.text = json.dumps(responseWithError)
            self._library.sideLoad("channel.zip", "user", "pass")
        self.assertTrue(errorMessage in str(context.exception))
    
    def test_side_load_data_success(self):
        with patch('Library.webDriver.requests.post') as mock_post:
            mock_post.return_value.status_code = 200
            mock_post.return_value.text = json.dumps(respWithEmptyValue)  
            try:
                self._library.sideLoad("channel.zip", "user", "pass")
            except Exception:
                self.fail("inputDeepLinkingData() raised Exception unexpectedly!")

        

    

    

    
        

