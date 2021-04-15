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

describe('test_Roku_Recommends_deeplinking', () => {
    before(() => {
        library = new rokuLibrary.Library("192.168.1.64");
    });

    it('side loading', async function() { 
        this.timeout(25000);
        await library.sideLoad("../channels/Roku_Recommends.zip", "rokudev", "aaaa");
        await library.sendKey('Home');
    });

    it('should launch the channel', async function() { 
        this.timeout(5000);
        await library.launchTheChannel('dev', '12', 'movie');
        await library.verifyIsChannelLoaded('dev');
    });
    
    it('Verify is playback started', async function() { 
        this.timeout(50000);
        const res = await library.verifyIsPlaybackStarted(20, 2);
        expect(res).equal(true);
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

