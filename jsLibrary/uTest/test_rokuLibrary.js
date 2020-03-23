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

const ecp = require("../library/rokuLibrary");
const expect  = require("chai").expect;
const nock = require('nock');
const {
    respWithEmptyValue, 
    requestLaunch, 
    resp, 
    requestPress,
    requestInstall,
    requestSequence,
    respWithDeviceInfo,
    respWithApps,
    respWithApp,
    respWithPlayerInfo,
    respWithElement,
    requestElement,
    respWithElements,
    requestTimeout,
    respWithSource,
    responseWithError,
    uiElement
} = require('./responses')
let library; 
const baseUrl = 'http://localhost:9000/v1/session';
const sessionId = resp.sessionId;

describe('rokuLibrary tests', () => {
    before(async () => {
        nock('http://localhost:9000/v1')
        .post('/session', {'ip' : '127.0.0.1', 'timeout': 20000, 'pressDelay': 2000})
        .reply(200, resp);
        library = new ecp.Library('127.0.0.1', 20000, 2000);
    });

    it('test launchTheChannel success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/launch`, requestLaunch)
        .reply(200, respWithEmptyValue);
        const result = await library.launchTheChannel('dev', '', '');
        expect(result).to.equal(true);
    });

    it('test launchTheChannel error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/launch`, requestLaunch)
        .reply(500, responseWithError);
        try{
            const result = await library.launchTheChannel('dev');
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test sendKeypress success', async function() { 
        this.timeout(4000);
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestPress)
        .reply(200, respWithEmptyValue);
        const result = await library.sendKey('select', 0.5);
        expect(result).to.equal(true);
    });

    it('test sendKeypress error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestPress)
        .reply(500, responseWithError);
        try{
            const result = await library.sendKey('select', 0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get apps success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/apps`)
        .reply(200, respWithApps);
        const result = await library.getApps();
        expect(JSON.stringify(respWithApps.value)).to.equal(JSON.stringify(result));
    });

    it('test get apps error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/apps`)
        .reply(500, responseWithError);
        try {
            const result = await library.getApps();
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test verify is screen loaded success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element`,requestElement)
        .reply(200, respWithElement);
        const result = await library.verifyIsScreenLoaded({
            "elementData" :[{
                "using": "tag",
                "value": "Label"
            }]
        });
        expect(result).to.equal(true);
    });

    it('test verify is screen loaded error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element`,requestElement)
        .reply(500, responseWithError);
        try {
            const result = await library.verifyIsScreenLoaded({
                "elementData" :[{
                    "using": "tag",
                    "value": "Label"
                }]
            }, 1, 1);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test send keys success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestSequence)
        .reply(200, respWithEmptyValue);
        const result = await library.sendKeys(['up', 'select'], 0.5);
        expect(result).to.equal(true);
    });

    it('test send keys error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestSequence)
        .reply(500, responseWithError);
        try {
            const result = await library.sendKeys(['up', 'select'], 0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test send word success',  async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestSequence)
        .reply(200, respWithEmptyValue);
        const result = await library.sendWord('ts', 0.5);
        expect(result).to.equal(true);
    });

    it('test send word error',  async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestSequence)
        .reply(500, responseWithError);
        try {
            const result = await library.sendWord('ts', 0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get element success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element`, requestElement)
        .reply(200, respWithElement);
        const result = await library.getElement({
            "elementData" :[{
                "using": "tag",
                "value": "Label"
            }]
        }, 0.5);
        expect(JSON.stringify(respWithElement.value)).to.equal(JSON.stringify(result));
    });

    it('test get element error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element`, requestElement)
        .reply(500, responseWithError);
        try {
            const result = await library.getElement({
                "elementData" :[{
                    "using": "tag",
                    "value": "Label"
                }]
            }, 0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get elements success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/elements`, requestElement)
        .reply(200, respWithElements);
        const result = await library.getElements({
            "elementData" :[{
                "using": "tag",
                "value": "Label"
            }]
        }, 0.5);
        expect(JSON.stringify(respWithElements.value)).to.equal(JSON.stringify(result));
    });

    it('test get elements error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/elements`, requestElement)
        .reply(500, responseWithError);
        try {
            const result = await library.getElements({
                "elementData" :[{
                    "using": "tag",
                    "value": "Label"
                }]
            }, 0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get focused element success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element/active`)
        .reply(200, respWithElement);
        const result = await library.getFocusedElement(0.5);
        expect(JSON.stringify(respWithElement.value)).to.equal(JSON.stringify(result));
    });

    it('test get focused element error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element/active`)
        .reply(200, respWithElement);
        try {
            const result = await library.getFocusedElement(0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test verify is channel loaded success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/current_app`)
        .reply(200, respWithApp);
        const result = await library.verifyIsChannelLoaded('id');
        expect(result).to.equal(true);
    });

    it('test verify is channel loaded error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/current_app`)
        .reply(500, responseWithError);
        try {
            const result = await library.verifyIsChannelLoaded('id');
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get current channel info success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/current_app`)
        .reply(200, respWithApp);
        const result = await library.getCurrentChannelInfo();
        expect(JSON.stringify(respWithApp.value)).to.equal(JSON.stringify(result));
    });

    it('test get current channel info error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/current_app`)
        .reply(500, responseWithError);
        try {
            const result = await library.getCurrentChannelInfo();
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get device info success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}`)
        .reply(200, respWithDeviceInfo);
        const response = await library.getDeviceInfo();
        expect(JSON.stringify(respWithDeviceInfo.value)).to.equal(JSON.stringify(response));
    });

    it('test get device info error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}`)
        .reply(500, responseWithError);
        try {
            const result = await library.getDeviceInfo();
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get player success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/player`)
        .reply(200, respWithPlayerInfo);
        const result = await library.getPlayerInfo();
        expect(result.Position).to.equal(1000);
    });

    it('test get player error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/player`)
        .reply(500, responseWithError);
        try {
            await library.getPlayerInfo();
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test verify is playack started success', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/player`)
        .reply(200, respWithPlayerInfo);
        const result = await library.verifyIsPlaybackStarted(0.5);
        expect(result).to.equal(true);
    });

    it('test verify is playack started error', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/player`)
        .reply(500, responseWithError);
        try {
            const result = await library.verifyIsPlaybackStarted(0.5);
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

    it('test get attribute success', async () => { 
        const result = await library.getAttribute(uiElement, 'color');
        expect(result).to.equal('#ffffff6f');
    });

    it('test get attribute error', async () => { 
        const result = await library.getAttribute(uiElement, 'attrNotExist');
        expect(result).to.equal(null);
    });

    it('test send input data success', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/input`)
        .reply(200, respWithEmptyValue);
        const result = await library.inputDeepLinkingData('dev', '12', 'movie');
        expect(result).to.equal(true);
    });

    it('test send input data error', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/input`)
        .reply(500, responseWithError);
        try {
            const result = await library.inputDeepLinkingData('dev', '12', 'movie');
            expect(result).to.not.equal(true);
        } catch(e) {
            expect(e.message).to.equal('Error');
        }
    });

});

