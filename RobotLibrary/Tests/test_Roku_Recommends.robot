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
&{DATA2}=  using=text  value=No Content to play
@{Params2}=  &{DATA2}
&{Grid}=  using=tag  value=GridView
@{GridArray}=  &{Grid}
&{GridParams}=  elementData=${GridArray}
&{Row}=  using=tag  value=Row
@{RowArray}=  &{Row}
&{RowParams}=   elementData=${RowArray}


*** Test Cases ***
Verify is channel launched
    Side load  ../channels/Roku_Recommends.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Should be 3 Rows on Grid initially
    @{elements}=  Get elements  ${RowParams}  2
    ${len}=    Get length    ${elements}
    should be equal as numbers   ${len}  3

Should be 4 Rows on Grid after Down key press
    Send key  Down
    @{elements}=  Get elements  ${RowParams}  2
    ${len}=    Get length    ${elements}
    should be equal as numbers   ${len}  4