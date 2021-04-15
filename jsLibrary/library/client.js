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

const axios = require('axios')
const BASE_URL = "http://localhost:9000/v1/session"


class Client {

    constructor(ip, timeout, delay) {
        this.ip = ip;
        this.timeout = timeout;
        this.delay = delay;
        this.sessionId = "";
    }
    
    async createSession(ip, timeout, delay) {
        return await this.doRequest('', 'post', {'ip' : ip, 'timeout': timeout, 'pressDelay': delay}, {}, false);
    }

    async launch(channel_code, contentID, contentType) {
        return await this.doRequest('/launch', 'post', {'channelId' : channel_code, 'contentId': contentID, 'contentType': contentType});
    }

    async sendKeypress(key) {
        return await this.doRequest('/press', 'post', {'Button' : key});
    }

    async deleteSession() {
        return await this.doRequest('', 'delete');
    }

    async getDeviceInfo() {
        return await this.doRequest('', 'get');
    }

    async getApps() {
        return await this.doRequest('/apps', 'get');
    }

    async getCurrentApp() {
        return await this.doRequest('/current_app', 'get');
    }

    async sendSequence(sequence) {
        return await this.doRequest('/press', 'post', {'button_sequence' : sequence});
    }

    async sendInputData(channelId, contentID, mediaType) {
        return await this.doRequest('/input', 'post', {'channelId': channelId, 'contentId' : contentID, 'contentType': mediaType });
    }

    async getPlayerInfo() {
        return await this.doRequest('/player', 'get');
    }

    async getScreenSource() {
        return await this.doRequest('/source', 'get');
    }

    async getActiveElement() {
        return await this.doRequest('/element/active', 'post');
    }

    async setTimeouts( timeoutType, delay) {
        return await this.doRequest('/timeouts', 'post',  {'type': timeoutType, 'ms': delay});
    }

    async sendInstallChannel(channelCode) {
        return await this.doRequest('/install', 'post', {'channelId' : channelCode});
    }

    async getUiElements(data) {
        return await this.doRequest('/elements', 'post', data);
    }

    async getUiElement(data) {
        return await this.doRequest('/element', 'post', data);
    }

    async sideLoadChannel(form) {
        console.log
        return await this.doRequest('/load', 'post', form, form.getHeaders());
    }

    async addSessionId(url) {
        if (this.sessionId == '') {
            let result = await this.createSession(this.ip, this.timeout, this.delay);
            if (result.status == 200) {
                this.sessionId = result.data.sessionId;
            }
        }
        url = `/${this.sessionId}${url}`;
        return url;
    }

    async doRequest(url, method, body = {}, headers = {}, addSessionId = true) {
        try {
            if (addSessionId == true) {
                url = await this.addSessionId(url);
            }
            const result = await axios({
                    method: method,
                    url: `${BASE_URL}${url}`,
                    data: body,
                    headers: {
                        ...headers
                    }
                });
            return result;
        } catch (errorResponse) {    
            const response = errorResponse.response;
            if (response == undefined) {
                throw new Error('Could not get any response');
            }
            const status = response.status;
            let errorMessage;
            if (status == 400) {
                errorMessage = response.data;
            } else {
                errorMessage = response.data.value.message;
            }
            throw new Error(errorMessage);
        }
    }
}

module.exports.Client = Client