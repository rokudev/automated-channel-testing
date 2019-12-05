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
${pressDelay}  2000
${channel_code}  dev
&{SeriesDetails}=  using=text  value=No Content to play
@{SeriesDetailsArray}=  &{SeriesDetails}
&{SeriesDetailsParams}=  elementData=${SeriesDetailsArray}
&{Grid}=  using=tag  value=GridView
@{GridArray}=  &{Grid}
&{GridParams}=  elementData=${GridArray}
&{MovieDetails}=  using=text  value=Play
@{MovieDetailsArray}=  &{MovieDetails}
&{MovieDetailsParams}=  elementData=${MovieDetailsArray}

*** Test Cases ***
Check if channel exist on the device
    @{apps}=  Get apps
    Verify is channel exist    ${apps}  ${channel_code}

Verify is channel launched
    Launch the channel  ${channel_code}
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Verify series details screen button
    Send key  Select   3
    Verify is screen loaded    ${SeriesDetailsParams}  3  4
 
Verify movies details screen button
    Send key  Back  3
    Verify is screen loaded    ${GridParams}
    Send key  Down  3
    Send key  Select  3
    Verify is screen loaded    ${MovieDetailsParams}