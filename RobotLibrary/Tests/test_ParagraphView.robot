########################################################################
# Copyright 2019 Roku, Inc.
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
*** Settings ***
Documentation  Test 2
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}
Library  Collections


*** Variables ***
${ip_address}  192.168.1.94
${server_path}  D:/projects/go/webDriver/src/main.exe
${timeout}  20000
${pressDelay}  3000
${channel_code}  dev
&{ParagraphViewData}=  using=tag  value=ParagraphView
@{ParagraphViewArray}=  &{ParagraphViewData}
&{ParagraphViewParams}=  elementData=${ParagraphViewArray}
&{HeaderColorData}=  using=attr  attribute=color  value=#22ffffff
&{HeaderTextData}=  using=text  value=Header Text
@{HeaderArray}=  &{HeaderColorData}  &{HeaderTextData}
&{HeaderParams}=  elementData=${HeaderArray}
&{CodeIndexData}=  using=attr  attribute=index  value=4
&{CodeColorData}=  using=attr  attribute=color  value=#ffff22ff
@{CodeArray}=  &{CodeColorData}  &{CodeIndexData}
&{CodeParams}=  elementData=${CodeArray} 

*** Test Cases ***
Verify is channel launched
    Launch the channel  ${channel_code}
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${ParagraphViewParams}

Verify header color
    Verify is screen loaded    ${HeaderParams}

Verify reload linking code
    &{element}=  Get element   ${CodeParams}
    ${code}=  Get attribute  ${element}  text
    Send key  Select
    &{element2}=  Get element   ${CodeParams}  
    ${newCode}=  Get attribute  ${element2}  text
    Run Keyword If  "${code}" == "${newCode}"   Fail