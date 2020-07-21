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

const getChildNodesByTag = (node, name) => {
    const children = node.Nodes;
    const result = []
    if (children == null) {
        return result
    }
    name = name.toLowerCase()
    children.forEach((value) => {
        if (value.XMLName.Local.toLowerCase() == name) {
            result.push(value)
        }
        const elements = getChildNodesByTag(value, name)
        result.push(...elements)
    })
    return result
}

let library; 

describe('test_7-EpisodePicker_focus', () => {
    before(async function() {
        this.timeout(20000);
        library = new rokuLibrary.Library("192.168.2.11");
        await library.sideLoad("../channels/7_EpisodePickerScreen.zip", "rokudev", "aaaa");
    });

    it('Check if channel exist on the device', async function() { 
        this.timeout(5000);
        const apps = await library.getApps();
        const res = library.verifyIsChannelExist(apps, 'dev');
        expect(res).equal(true);
    });

    it('Verify is initial screen loaded', async function() { 
        this.timeout(10000);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'tag', 'value': 'GridView'}]});
        expect(res).to.equal(true);
    });

    it('Verify focused element on grid screen', async function() { 
        this.timeout(50000);
        await library.sendKey('down', 3);
        const element = await library.getFocusedElement();
        const children = getChildNodesByTag(element, "poster");
        if (children.length > 0) {
            const poster = children[0];
            const posterUri = library.getAttribute(poster, "uri");
            expect(posterUri).to.equal("https://blog.roku.com/developer/files/2016/10/twitch-poster-artwork.png")
        } else {
            throw new Error("Can't find poster")
        }
    });
    
    it('Verify is details screen loaded', async function() { 
        this.timeout(50000);
        await library.sendKeys(['up','select'], 3);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'tag', 'value': 'DetailsView'}]});
        expect(res).equal(true);
    });

    it('Verify focused element on details screen', async function() { 
        this.timeout(50000);
        const element = await library.getFocusedElement();
        const children = getChildNodesByTag(element, 'label');
        if (children.length > 0) {
            const label = children[0];
            const labelText = library.getAttribute(label, 'text');
            expect(labelText).to.equal("Episodes")
        } else {
            throw new Error("Can't find label")
        }
    });

    it('Verify is Categort list loaded', async function() { 
        this.timeout(50000);
        await library.sendKey('select', 3);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'tag', 'value': 'CategoryListView'}]});
        expect(res).equal(true);
    });

    it('Verify focused element on episodes screen (episodes list)', async function() { 
        this.timeout(50000);
        await library.sendKey('down');
        const element = await library.getFocusedElement();
        const children = getChildNodesByTag(element, 'label');
        if (children.length > 0) {
            const label = children[0];
            const labelText = library.getAttribute(label, 'text');
            expect(labelText).to.equal("Ideas Worth Spreading")
        } else {
            throw new Error("Can't find label")
        }
    });

    it('Verify focused element on episodes screen (seasons list)', async function() { 
        this.timeout(50000);
        await library.sendKey('left');
        const element = await library.getFocusedElement();
        const children = getChildNodesByTag(element, 'label');
        if (children.length > 0) {
            const label = children[0];
            const labelText = library.getAttribute(label, 'text');
            expect(labelText).to.equal("Season  1")
        } else {
            throw new Error("Can't find label")
        }
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

