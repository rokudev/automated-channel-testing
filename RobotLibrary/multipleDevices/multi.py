########################################################################
# Copyright 2020 Roku, Inc.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.
########################################################################

import sys
import asyncio
import subprocess
import json
from string import Template
import codecs
if sys.platform == "win32":
    asyncio.set_event_loop_policy(asyncio.WindowsProactorEventLoopPolicy())

async def read_stream(stream, cb):
    while True:
        line = await stream.readline()
        if line:
            cb(line)
        else:
            break


async def stream_subprocess(cmd, stdout_cb, stderr_cb):
    try:
        process = await asyncio.create_subprocess_exec(
            *cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
        )

        await asyncio.wait(
            [
                read_stream(process.stdout, stdout_cb),
                read_stream(process.stderr, stderr_cb),
            ]
        )
        rc = await process.wait()
        return process.pid, rc
    except OSError as e:
        return e

async def prc(*aws):
    await  asyncio.gather(*aws)

def execute(*aws):
    asyncio.run(prc(*aws))

def printer(label):
    def pr(*args):
        print(label, *args)

    return pr

def runners(cmds):
    for cmd in cmds:
        out = printer(f"{cmd['name']}.stdout")
        err = printer(f"{cmd['name']}.stderr")
        yield stream_subprocess(cmd['command'], out, err)

def getCmdList(config):
    devices = config['devices']
    outputdir = config['outputdir']
    if devices == None:
        raise Exception("devices field is empty")
    
    for device in devices:
        cmd = {
            'command': [sys.executable, "-m", "robot.run"],
            'name': device
        }
        deviceFields = devices[device]
        for key in deviceFields:
            cmd['command'].extend(['--variable', f'{key}:{deviceFields[key]}'])
        cmd['command'].extend(['--outputdir', f'{outputdir}/{device}'])
        cmd['command'].append(config['test'])
        yield cmd

def createGeneralFile(config, fileName, title):
    devicesList = []
    outputdir = config['outputdir']
    for device in config['devices']:
        devicesList.append({"name": device, "src": f'{device}/{fileName}.html'})
    with  codecs.open("multipleDevices/template.html", 'r', 'utf-8') as template:
        res = Template(template.read()).safe_substitute(content=devicesList, title=title)
    with open(f"{outputdir}/multi_device_{fileName}.html", "w") as new_file:  
        new_file.write(res)

if __name__ == "__main__":
    try:
        config_path = sys.argv[1]
        with open(config_path) as myfile:
            data = json.load(myfile)
        if data['server_path'] == None:
            raise Exception("server_path field is required")
        else:
            process = subprocess.Popen(data['server_path'])
        cmdList = getCmdList(data)
        execute(*runners(cmdList))
        process.kill()
        createGeneralFile(data, "log", "Logs")
        createGeneralFile(data, "report", "Reports")
    except IndexError:
        print("Config path is required as first argument")
    except FileNotFoundError:
        print('This file doesn\'t exist')
    except Exception as e:
        print(e)
