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

from robot.api.deco import keyword
from Library.webDriver import WebDriver
from robot.libraries.BuiltIn import BuiltIn
from time import sleep
from robot.api import logger
import subprocess
import json
import re
import urllib
from datetime import datetime, timedelta

class RobotLibrary:

    ROBOT_LIBRARY_SCOPE = 'GLOBAL'
    ROBOT_LISTENER_API_VERSION = 2

    def __init__(self, ip, timeout = 0, pressDelay = 0, path = ""):
        self._process = None
        if len(path) > 0:
            self._process = subprocess.Popen(path)
        self.ROBOT_LIBRARY_LISTENER = self
        self.locatorHandlers = {
            "attr": self._checkAttribute,
            "tag": self._checkTag,
            "text": self._checkText
        }
        self._client = WebDriver(ip, timeout, pressDelay)
        self.markTimer()
       
    def close(self):
        self._client.quiet()
        if self._process != None:
            self._process.kill()
    
    @keyword("Mark timer")
    def markTimer(self):
        self._startTime = datetime.now()
    
    @keyword("Get timer")
    def getTimer(self):
        currentTime = datetime.now()
        delta = currentTime - self._startTime
        return int(delta / timedelta(milliseconds=1))

    @keyword("Side load")
    def sideLoad(self, path, user, password):
        multipart_form_data = {
            'channel': ('channel.zip', open(path, 'rb')),
            'username': (None, user),
            'password': (None, password)
        }
        response = self._client.side_load(multipart_form_data)
        self._checkResponse(response)

    @keyword("Launch the channel")
    def launchTheChannel(self, channel_code, contentId = "", mediaType = ""):
        launch_response = self._client.send_launch_channel(channel_code, contentId, mediaType)
        self._checkResponse(launch_response)
    
    @keyword("Get apps")
    def getApps(self):
        apps_response = self._client.get_apps()
        self._checkResponse(apps_response)
        res = json.loads(apps_response.text)
        return res['value']

    @keyword("Verify is channel exist")
    def verifyIsChannelExist(self, apps, id):
        for app in apps:
            if app['ID'] == id:
                return True
        raise Exception("Channel doesn't exist")
    
    @keyword("Verify is screen loaded")
    def verifyIsScreenLoaded(self, data: object, retries = 10, delay = 1):
        print(data)
        while retries > 0:
            ui_layout_response = self._client.get_ui_element(data)
            if ui_layout_response.status_code != 200:
                retries -= 1
                sleep(delay)
            else:
                return True
        raise Exception("Can't find element")

    @keyword("Send key")
    def pressBtn(self, key_press: str, delay = 2):
        sleep(delay)
        key_press_response = self._client.send_keypress(key_press)
        self._checkResponse(key_press_response)

    @keyword("Send word")
    def sendWord(self, word: str, delay = 2):
        sleep(delay)
        for c in word:
            sleep(0.5)
            key_press_response = self._client.send_keypress(urllib.parse.quote(f"LIT_{c}"))
            self._checkResponse(key_press_response)

    
    @keyword("Send keys")
    def sendButtonSequence(self, sequence, delay = 2):
        sleep(delay)
        key_press_response = self._client.send_sequence(sequence)
        self._checkResponse(key_press_response)

    
    @keyword("Get element")
    def getElement(self, data: object, delay = 1):
        sleep(delay)
        ui_layout_response = self._client.get_ui_element(data)
        self._checkResponse(ui_layout_response)
        res = json.loads(ui_layout_response.text)
        return res['value']
    
    @keyword("Get elements")
    def getElements(self, data: object, delay = 1):
        sleep(delay)
        ui_layout_response = self._client.get_ui_elements(data)
        self._checkResponse(ui_layout_response)
        res = json.loads(ui_layout_response.text)
        return res['value']

    @keyword("Get focused element")
    def getFocusedElement(self):
        ui_layout_response = self._client.get_active_element()
        self._checkResponse(ui_layout_response)
        res = json.loads(ui_layout_response.text)
        return res['value']
    
    @keyword("Verify is channel loaded")
    def verifyIsChannelLoaded(self, id, retries = 10, delay = 1):
        while retries > 0:
            app_response = self._client.get_current_app()
            self._checkResponse(app_response)
            res = json.loads(app_response.text)
            if res['value']['ID'] != id:
                retries -= 1
                sleep(delay)
            else:
                return True    
        raise Exception("Channel isn't launched")

    @keyword("Get current channel info")
    def getCurrentChannelInfo(self):
        app_response = self._client.get_current_app()
        self._checkResponse(app_response)
        res = json.loads(app_response.text)
        return res['value']

    @keyword("Get device info")
    def getDeviceInfo(self):
        response = self._client.get_device_info()
        self._checkResponse(response)
        res = json.loads(response.text)
        return res['value']
    
    @keyword("Get player info")
    def getPlayerInfo(self):
        response = self._client.get_player_info()
        self._checkResponse(response)
        res = json.loads(response.text)
        value = res['value']
        position = value['Position']
        duration = value['Duration']
        value['Position'] = self._getMsFromString(position)
        value['Duration'] = self._getMsFromString(duration)
        return value

    @keyword("Verify is playback started")
    def verifyIsPlaybackStarted(self, retries = 10, delay = 1):
        while retries > 0:
            response = self._client.get_player_info()
            res = json.loads(response.text)
            if response.status_code != 200 or res['value']['State'] != 'play':
                retries -= 1
                sleep(delay)
            else:
                return True
        raise Exception("Invalid player state")
    
    @keyword("Set timeout")
    def setTimeout(self, timeout: int):
        response = self._client.set_timeouts("implicit", timeout)
        self._checkResponse(response)
    
    @keyword("Set press delay")
    def setDelay(self, delay: int):
        response = self._client.set_timeouts("pressDelay", delay)
        self._checkResponse(response)
    
    @keyword("Get attribute")
    def getAttribute(self, element, attr):
        for attrObj in element['Attrs']:
            if attrObj['Name']["Local"] == attr:
                return  attrObj['Value']
        raise Exception("Can't find attribute")
    
    @keyword("Input deep linking data")
    def inputDeepLinkingData(self, channelId, contentId, mediaType):
        launch_response = self._client.send_input_data(channelId, contentId, mediaType)
        self._checkResponse(launch_response)

    @keyword("Get child nodes")
    def getChildNodes(self, parentNode, locators):
        childNodes = parentNode.Nodes
        result = []
        if childNodes == None:
            return result
        if locators == None:
            return result
        for node in childNodes:
            if self._isElementMatchLocators(node, locators) == True:
                result.append(node)
            result.extend(self.getChildNodes(node, locators))
        return result
        
    def _isElementMatchLocators(self, node, locators):
        for locator in locators:
            if hasattr(locator, 'using') == False:
                return False
            handler = self.locatorHandlers[locator.using]
            if handler == None:
                return False
            isMatch = handler(node, locator)
            if isMatch == False:
                return False
        return True
    
    def _checkAttribute(self, node, locator):
        if hasattr(node, 'Attrs') == False or hasattr(locator, 'value') == False or hasattr(locator, 'attribute') == False:
            return False
        for attr in node.Attrs:
            matchName = attr.Name.Local.lower() == locator.attribute.lower()
            matchValue = attr.Value.lower() == locator.value.lower()
            if matchName and matchValue:
                return True
        return False
    
    def _checkTag(self, node, locator):
        return node.XMLName.Local.lower() == locator.value.lower()

    def _checkText(self, node, locator):
        if hasattr(node, 'Attrs') == False or hasattr(locator, 'value') == False:
            return False
        for attr in node.Attrs:
            matchName = attr.Name.Local.lower() == "text"
            matchValue = attr.Value.lower() == locator.value.lower()
            if matchName and matchValue:
                return True
        return False

    def _checkResponse(self, response):
        if response.status_code == 400:
            raise Exception(response.text)
        elif response.status_code != 200:
            res = json.loads(response.text)
            raise Exception(res['value']['message'])
    
    def _getMsFromString(self, strWithMs):
        ms = 0
        if type(strWithMs) == str:
            searchRes = re.search(r'\d+', strWithMs)
            if searchRes != None:
                ms = int(searchRes.group())
        return ms
    
    
    
    
    
     


    
        
    
    

    
