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

describe('test_SearchView', () => {
    before(() => {
        library = new rokuLibrary.Library("192.168.1.11");
    });

    it('should launch the channel', async function() { 
        this.timeout(25000);
        await library.sideLoad("../channels/SearchView.zip", "rokudev", "aaaa");
        await library.verifyIsChannelLoaded('dev');
    });

    it('Verify is initial screen loaded', async function() { 
        this.timeout(10000);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'tag', 'value': 'SearchView'}]});
        expect(res).to.equal(true);
    });
    
    it('Verify search input', async function() { 
        this.timeout(50000);
        const input = 'hello@1 r~';
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'text', 'value': 'Enter search term'}]});
        expect(res).to.equal(true);
        await library.sendWord(input);
        const inputLabel = await library.getElement({
            'elementData': [{'using': 'tag', 'value': 'Label'}], 
            'parentData': [{'using': 'tag', 'value': 'TextEditBox'}]
        });
        const text = library.getAttribute(inputLabel, 'text');
        expect(text).to.equal(input);
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

