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
const {getArgs} = require('./utils');
let library; 
const args = getArgs();
let ip = args.ip_address ||  "192.168.1.12";
describe('Channel should be launched', () => {
    before(() => {
        library = new rokuLibrary.Library(ip);
    });

    it('should launch the channel', async function() { 
        this.timeout(25000);
        await library.launchTheChannel('dev');
        await library.verifyIsChannelLoaded('dev');
    });

    it('Verify is initial screen loaded', async function() { 
        this.timeout(10000);
        const res = await library.verifyIsScreenLoaded({ 
            'elementData': [{'using': 'tag', 'value': 'GridView'}
        ]});
        expect(res).to.equal(true);
    });
    
    it('Verify posters', async function() { 
        this.timeout(50000);
        const elements = await library.getElements({'elementData': [{'using': 'attr', 'attribute': 'name', 'value': 'poster'}]}, 4);
        let poster = library.getAttribute(elements[0], 'uri');
        expect(poster).to.equal('https://roku-blog.s3.amazonaws.com/developer/files/2017/04/Roku-Recommends-thumbnail.png');
        poster = library.getAttribute(elements[1], 'uri');
        expect(poster).to.equal('https://blog.roku.com/developer/files/2016/10/twitch-poster-artwork.png');
    });

    after(async () => {
        await library.close();
    });
});

