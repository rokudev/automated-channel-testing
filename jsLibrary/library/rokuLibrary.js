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

const ecp = require("./client");
const FormData = require('form-data');
const fs = require('fs');

const getMsFromString = (value) => {
    const result = value.split(' ');
    return result[0];
}

class Library {
    constructor(ip, timeout = 20000, delay = 2000) {
        this.client = new ecp.Client(ip, timeout, delay);
        this.markTimer();
    }

    async close() {
        await this.client.deleteSession();
    }

    async launchTheChannel(channelCode, contentId = "", contentType = "") {
        await this.client.launch(channelCode, contentId, contentType);
        return true;
    }

    async getApps() {
        const result = await this.client.getApps();
        return result.data.value;
    }

    sleep(milliseconds) {
        return new Promise(resolve => setTimeout(resolve, milliseconds));
    }

    async sendKey(key, delay = 2) {
        await this.sleep(delay*1000);
        await this.client.sendKeypress(key);
        return true;
    }

    async getElement(data, delay = 2) {
        await this.sleep(delay*1000);
        const result = await this.client.getUiElement(data);
        return result.data.value;
    }

    async getElements(data, delay = 2) {
        await this.sleep(delay*1000);
        const result = await this.client.getUiElements(data);
        return result.data.value;
    }

    async getFocusedElement(delay = 2) {
        await this.sleep(delay*1000);
        const result = await this.client.getActiveElement();
        return result.data.value;
    }

    async sendKeys(sequence, delay = 2){
        await this.sleep(delay*1000);
        await this.client.sendSequence(sequence);
        return true;
    }

    async verifyIsScreenLoaded(data, retries = 10, delay = 1) {
        while (retries > 0) {
            try {
                await this.client.getUiElement(data);
                return true;
            } catch {
                retries -= 1;
                if (retries == 0) {
                    return false;
                }
                await this.sleep(delay*1000);
            }
        }
    }

    async getCurrentChannelInfo(self) {
        const response = await this.client.getCurrentApp();
        return response.data.value;
    }

    async getDeviceInfo() {
        const response = await this.client.getDeviceInfo();
        return response.data.value;
    }

    async getPlayerInfo() {
        const response = await this.client.getPlayerInfo();
        let value = response.data.value;
        value.Position = parseInt(getMsFromString(value.Position));
        value.Duration = parseInt(getMsFromString(value.Duration));
        return value;
    }

    async setTimeout(timeout) {
        await this.client.setTimeouts("implicit", timeout);
        return true;
    }

    async sideLoad(path, user, pass) {
        const form = new FormData();
        const stream = fs.createReadStream(path);
        form.append('channel', stream, {
            contentType: 'application/zip'
        });
        form.append('username', user);
        form.append('password', pass);
        await this.client.sideLoadChannel(form);
        return true;
    }

    async setDelay(delay) {
        await this.client.setTimeouts("pressDelay", delay);
        return true;
    }

    async sendWord(word, delay = 2) {
        await this.sleep(delay*1000);
        let symbols = [];
        word.split('').forEach((el) =>  {
            symbols.push(`LIT_${el}`);
        });
        await this.sendKeys(symbols, 0);
        return true;   
    }

    async verifyIsPlaybackStarted(retries = 10, delay = 1) {
        while (retries > 0) {  
            let response = await this.client.getPlayerInfo();
            if (response.data.value.State == 'play') {
                return true;
            } else {
                retries -= 1;
                if (retries == 0) {
                    return false;
                }
                await this.sleep(delay*1000);
            }
        }
    }

    async verifyIsChannelLoaded(id, retries = 10, delay = 1) {
        while (retries > 0) {  
            let response = await this.client.getCurrentApp();
            if (response.data.value.ID == id) {
                return true;
            } else {
                retries -= 1;
                if (retries == 0) {
                    return false;
                }
                await this.sleep(delay*1000);
            }         
        }
    }

    async inputDeepLinkingData(channelId, contentId, mediaType) {
        await this.client.sendInputData(channelId, contentId, mediaType);
        return true;
    }

    markTimer() {
        this.startTime = new Date().getTime();
    }
    
    getTimer() {
        const currentTime = new Date().getTime();
        return currentTime - this.startTime;
    }

    getAttribute(element, attr) {
        const result = element.Attrs.find((attrObj) => attrObj.Name.Local == attr);
        const value = result ? result.Value : null;
        return value;
    }

    verifyIsChannelExist(apps, id) {
        let index =  apps.findIndex((channel) =>  channel.ID == id);
        return index > -1;
    }
}

module.exports.Library = Library;
