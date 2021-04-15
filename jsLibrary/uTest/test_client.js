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

const ecp = require("../library/client");
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
    respWithSource
} = require('./responses')
let client; 
const baseUrl = 'http://localhost:9000/v1/session';
const sessionId = resp.sessionId;

describe('client tests', () => {
    before(async () => {
        nock('http://localhost:9000/v1')
        .post('/session', {'ip' : '127.0.0.1', 'timeout': 20000, 'pressDelay': 2000})
        .reply(200, resp);
        client = new ecp.Client('127.0.0.1', 20000, 2000);
    });

    it('test launch', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/launch`, requestLaunch)
        .reply(200, respWithEmptyValue);
        const response = await client.launch('dev', '', '');
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test sendKeypress', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestPress)
        .reply(200, respWithEmptyValue);
        const response = await client.sendKeypress('select');
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test install', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/install`, requestInstall)
        .reply(200, respWithEmptyValue);
        const response = await client.sendInstallChannel('1111');
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test send sequence', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/press`, requestSequence)
        .reply(200, respWithEmptyValue);
        const response = await client.sendSequence(['up', 'select']);
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test get device info', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}`)
        .reply(200, respWithDeviceInfo);
        const response = await client.getDeviceInfo();
        expect(JSON.stringify(respWithDeviceInfo)).to.equal(JSON.stringify(response.data));
    });

    it('test get apps', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/apps`)
        .reply(200, respWithApps);
        const response = await client.getApps();
        expect(JSON.stringify(respWithApps)).to.equal(JSON.stringify(response.data));
    });

    it('test get current app', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/current_app`)
        .reply(200, respWithApp);
        const response = await client.getCurrentApp();
        expect(JSON.stringify(respWithApp)).to.equal(JSON.stringify(response.data));
    });

    it('test get player', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/player`)
        .reply(200, respWithPlayerInfo);
        const response = await client.getPlayerInfo();
        expect(JSON.stringify(respWithPlayerInfo)).to.equal(JSON.stringify(response.data));
    });

    it('test get ui element', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element`,requestElement)
        .reply(200, respWithElement);
        const response = await client.getUiElement({
            "elementData" :[{
                "using": "tag",
                "value": "Label"
            }]
        });
        expect(JSON.stringify(respWithElement)).to.equal(JSON.stringify(response.data));
    });

    it('test get ui elements', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/elements`, requestElement)
        .reply(200, respWithElements);
        const response = await client.getUiElements({
            "elementData" :[{
                "using": "tag",
                "value": "Label"
            }]
        });
        expect(JSON.stringify(respWithElements)).to.equal(JSON.stringify(response.data));
    });

    it('test get active element', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/element/active`, {})
        .reply(200, respWithElement);
        const response = await client.getActiveElement();
        expect(JSON.stringify(respWithElement)).to.equal(JSON.stringify(response.data));
    });

    it('test set timeouts', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/timeouts`, requestTimeout)
        .reply(200, respWithEmptyValue);
        const response = await client.setTimeouts("implicit", 1000);
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test get screen source', async () => { 
        nock(baseUrl)
        .get(`/${sessionId}/source`)
        .reply(200, respWithSource);
        const response = await client.getScreenSource();
        expect(JSON.stringify(respWithSource)).to.equal(JSON.stringify(response.data));
    });

    it('test send input data', async () => { 
        nock(baseUrl)
        .post(`/${sessionId}/input`)
        .reply(200, respWithEmptyValue);
        const response = await client.sendInputData('dev', '12', 'movie');
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });

    it('test deleteSession', async () => { 
        nock(baseUrl)
        .delete(`/${sessionId}`)
        .reply(200, respWithEmptyValue);
        const response = await client.deleteSession();
        expect(JSON.stringify(respWithEmptyValue)).to.equal(JSON.stringify(response.data));
    });
});

