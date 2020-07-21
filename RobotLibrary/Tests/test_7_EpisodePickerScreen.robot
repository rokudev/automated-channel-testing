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
&{GridData}=  using=tag  value=GridView
@{GridArray}=  &{GridData}
&{GridParams}=  elementData=${GridArray}
&{DetailsData}=  using=tag  value=DetailsView
@{DetailsArray}=  &{DetailsData}
&{DetailsParams}=  elementData=${DetailsArray}
&{EncardData}=  using=text  value=Play again
@{EncardArray}=  &{EncardData}
&{EncardParams}=  elementData=${EncardArray}
&{CategoryData}=  using=tag  value=CategoryListView
@{CategoryArray}=  &{CategoryData}
&{CategoryParams}=  elementData=${CategoryArray}
@{GridKeys}=  Up  Select

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/7_EpisodePickerScreen.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Verify focused element on grid screen
    Send key  Down  4
    &{focusedEl}=  get focusedElement
    @{Nodes}=  Get From Dictionary	${focusedEl}  Nodes 
    ${uri}=  Get attribute  @{Nodes}[${1}]  uri
    Run keyword if  '${uri}'!='https://blog.roku.com/developer/files/2016/10/twitch-poster-artwork.png'  Fail

Verify is details screen loaded
    Send keys  ${GridKeys}  4
    Verify is screen loaded    ${DetailsParams}

Verify focused element on details screen
    &{focusedEl}=  get focusedElement
    @{Nodes}=  Get From Dictionary	${focusedEl}  Nodes 
    @{Nodes}=  Get From Dictionary  @{Nodes}[${0}]  Nodes
    ${text}=  Get Attribute  @{Nodes}[${2}]  text
    Run keyword if  '${text}'!='Episodes'  Fail

Verify is Categort list started
    Send key  Select  4
    Verify is screen loaded    ${CategoryParams}

Verify focused element on episodes screen (episodes list)
    Send key  Down
    &{focusedEl}=  get focusedElement
    @{Nodes}=  Get From Dictionary	${focusedEl}  Nodes 
    @{Nodes}=  Get From Dictionary  @{Nodes}[${1}]  Nodes
    ${uri}=  Get Attribute  @{Nodes}[${0}]  uri
    Run keyword if  '${uri}'!='https://blog.roku.com/developer/files/2016/10/ted-poster-artwork.png'  Fail

Verify focused element on episodes screen (seasons list)
    Send key  Left
    &{focusedEl}=  get focusedElement
    @{Nodes}=  Get From Dictionary	${focusedEl}  Nodes 
    @{Nodes}=  Get From Dictionary  @{Nodes}[${0}]  Nodes
    ${text}=  Get Attribute  @{Nodes}[${2}]  text
    Run keyword if  '${text}'!='Season${SPACE*2}1'  Fail
