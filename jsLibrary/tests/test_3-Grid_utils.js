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

describe('test_3-Grid_utils', () => {
    before(() => {
        library = new rokuLibrary.Library("192.168.2.11");
    });

    it('should launch the channel', async function() { 
        this.timeout(25000);
        await library.sideLoad("../channels/3_Grid.zip", "rokudev", "aaaa");
        await library.verifyIsChannelLoaded('dev');
    });

    it('Verify is initial screen loaded', async function() { 
        this.timeout(10000);
        const res = await library.verifyIsScreenLoaded({'elementData': [{'using': 'tag', 'value': 'GridView'}]});
        expect(res).to.equal(true);
    });
    
    it('Verify posters', async function() { 
        this.timeout(10000);
        const element = await library.getFocusedElement();
        const result = library.getChildNodes(element, [{"using": "tag", "value": "poster"}]);
        if (result.length > 0) {
            const uri = library.getAttribute(result[0], "uri");
            expect(uri).to.equal("https://roku-blog.s3.amazonaws.com/developer/files/2017/04/Roku-Recommends-thumbnail.png")
        } else {
            throw new Error('Can not find poster');
        }
    });

    it('Verify focused row title', async function() { 
        this.timeout(10000);
        const rowList = await library.getElement({"elementData" :[{"using": "tag", "value": "ZoomRowList"}]});
        const searchData = [{"using": "tag", "value": "RenderableNode"}, {"using": "attr", "attribute": "focused", "value": "true"}];
        const result = library.getChildNodes(rowList, searchData);
        if (result.length > 0) {
            const renderableNode = result[0];
            const labelSearchData = [{"using": "text", "value": "series"}];
            const labelList = library.getChildNodes(renderableNode, labelSearchData);
            expect(labelList.length).to.equal(1);
        } else {
            throw new Error('Can not find focused renderable node');
        }
    });

    after(async () => {
        await library.close();
        childProcess.kill();
    });
});

