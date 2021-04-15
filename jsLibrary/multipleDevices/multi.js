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

const { spawn } = require('child_process');
const fs = require("fs");

function getCmdList(config) {
    const devices = config.devices;
    if (devices === undefined) {
        throw new Error("devices field is empty");
    }
    let cmds = [];
    for(const device in devices) {
        let deviceFields = devices[device];
        let cmd = {
            options: '',
            name: device 
        }
        for(const field in deviceFields) {
            cmd.options += `--${field}=${deviceFields[field]} `;
        }
        cmds.push(cmd);
    }
    return cmds;
}

function startTest(cmds, config) {
    let numberOfProcess = cmds.length;
    const {test, outputdir, server_path} = config;
    const webDriver = spawn(server_path);
    const isWin = process.platform === "win32";
    const executeCmd = isWin ? 'mocha.cmd' : 'mocha';
    cmds.forEach(({options, name}) => {
        let child = spawn(executeCmd, [`${test}`,'--reporter', 'mochawesome', '--reporter-options', `reportDir=${outputdir},reportFilename=${name}`, options]);
        child.on('error', (messasge) => {
            console.log(`${name} error: ${messasge}`);
        });
        child.stdout.on('data', (data) => {
            console.log(`${name} stdout:\n${data}`);
        });

        child.on('exit', () => {
            numberOfProcess--;
            if (numberOfProcess === 0) {
                webDriver.kill();
                createGeneralReport(config);
            }
        });
    });
}

function createGeneralReport(config) {
    let deviceList = [];
    const {devices, outputdir} = config;
    for(const device in devices) {
        deviceList.push({'name': device, 'src': `${device}.html`});
    }
    let template = fs.readFileSync('multipleDevices/template.html', 'utf8');
    template = template.replace('$content', JSON.stringify(deviceList));
    template = template.replace('$title', 'Report');
    fs.writeFile(`${outputdir}/multi_device_report.html`, template, (err) => {
        if (err) throw err;
        console.log(`Multi-device report saved to ${outputdir}/multi_device_report.html`);
    });
}

function checkRequiredFields(json) {
    const fields = ['devices', 'server_path', 'outputdir', 'test'];
    fields.forEach( field => {
       if (json[field] === undefined) {
          throw new Error(`Please, add field \'${field}\' to config`);
       }
    });
}

try {
    const args = process.argv;
    const configPath = args[2];
    if (configPath === undefined) {
        throw new Error('Path to config is required');
    }
    const config = fs.readFileSync(configPath);
    const jsonData = JSON.parse(config);
    checkRequiredFields(jsonData);
    const cmdList = getCmdList(jsonData);
    startTest(cmdList, jsonData);
} catch (e) {
    console.log(e);
}