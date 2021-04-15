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

const rokuLibrary = require("../library/rokuLibrary");
const expect  = require("chai").expect;
const { spawn } = require('child_process');

const childProcess = spawn('D:/projects/go/webDriver/src/main.exe');

let library; 

describe('test_Roku_Recommends_input', () => {
    before(() => {
        library = new rokuLibrary.Library("192.168.1.64");
    });

    it('side loading', async function() { 
        this.timeout(25000);
        await library.sideLoad("../channels/Roku_Recommends.zip", "rokudev", "aaaa");
        await library.verifyIsChannelLoaded('dev');
        await library.sendKey('Home');
    });

    it('input deep linking', async function() { 
        this.timeout(50000);
        await library.inputDeepLinkingData('dev', '12', 'movie');
        library.markTimer();
        await library.sleep(1000);
        let res = await library.verifyIsPlaybackStarted(25, 1);
        expect(res).to.equal(true);
        let time = library.getTimer();
        expect(14000).greaterThan(time);
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

