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
Variables  ./../Library/variables.py 
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}
Library  Collections


*** Variables ***
${channel_code}  dev
&{DetailsData}=  using=tag  value=DetailsView
@{DetailsArray}=  &{DetailsData}
&{DetailsParam}=  elementData=@{DetailsArray}
@{KEYS}=  Down  Select

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/9_Bookmarks.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is details screen loaded
    Verify is screen loaded    ${DetailsParam}

Verify is playback started
    Send key  Select  2
    Verify is playback started

Bookmarks
    Sleep  12
    Send key  Back
    Verify is screen loaded    ${DetailsParam}
    Send keys  ${KEYS}  3
    Verify is playback started
    &{player}=  Get player info
    Run keyword if  ${player['Position']} < 10000  Fail
