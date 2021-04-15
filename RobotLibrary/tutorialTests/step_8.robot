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

*** Settings ***
Documentation  Test 2
Variables  ./../Library/variables.py 
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}
Library  Collections


*** Variables ***
${channel_code}  dev
@{keys}  Right  Select
&{Grid}=  using=tag  value=GridView
@{GridArray}=  &{Grid}
&{GridParams}=  elementData=${GridArray}
${Title}=   Roku Recommends
&{DetailsData}=  using=tag  value=DetailsView
@{DetailsArray}=  &{DetailsData}
&{DetailsParam}=  elementData=@{DetailsArray}
&{TextData}=  using=text  value=The Paula Deen Channel
@{TextArray}=  &{TextData}
&{TextParam}=  elementData=@{TextArray}
${validColor}=  "#ddddddff"

*** Test Cases ***
launch the channel and verify
    Launch the channel  ${channel_code}  
    Verify is channel loaded  ${channel_code}

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Get info about current channel and verify Title
    &{app}=  Get current channel info
    Run Keyword If  "${Title}" != "${app['Title']}"  Fail

Open details screen
    Send keys  ${keys}
    Verify is screen loaded    ${DetailsParam}

Get description element and verify color
    &{descriptionElement}=  Get element   ${TextParam}
    ${color}=  Get attribute  ${descriptionElement}  color
    Run Keyword If  "${validColor}"!="${color}"  Fail
