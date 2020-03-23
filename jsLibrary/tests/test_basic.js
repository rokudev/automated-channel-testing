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

describe('test_basic', () => {
    before(async function() {
        this.timeout(50000);
        library = new rokuLibrary.Library("192.168.2.121");
        await library.sideLoad("../sample/channel.zip", "rokudev", "aaaa");
    });

    it('should launch the channel', async function() { 
        this.timeout(15000);
        await library.verifyIsChannelLoaded('dev');
    });

    it('Check if details screen showed', async function() { 
        this.timeout(30000);
        await library.sendKey('select', 4);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'text', 'value': 'Barack Gates, Bill Obama'}]});
        expect(res).to.equal(true);
    });
    
    it('Check if playback started', async function() { 
        this.timeout(50000);
        let res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'text', 'value': 'Authenticate to watch'}]}, 2, 2);
        if (res == true) {
            await library.sendKey('select');
            res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'text', 'value': 'Please enter your username'}]});
            if (res == false) {
                expect.fail("Can't enter user name");
            }
            await library.sendWord('user');
            await library.sendKeys(['down', 'down', 'down', 'down', 'select']);
            await library.sendWord('pass');
            await library.sendKeys(['down', 'down', 'down', 'down', 'select']);
        } else {
            await library.sendKey('select');
        }
        res = await library.verifyIsPlaybackStarted();
        expect(res).to.equal(true);
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

