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
&{tagData}=  using=tag  value=RenderableNode
&{attrData}=  using=attr  attribute=focused  value=true
&{season1Data}=   using=attr  attribute=index  value=0 
&{season2Data}=   using=attr  attribute=index  value=1 
@{Season1Array}=  &{tagData}  &{attrData}  &{season1Data}
&{Season1Params}=  elementData=${Season1Array}
@{Season2Array}=  &{tagData}  &{attrData}  &{season2Data}
&{Season2Params}=  elementData=${Season2Array}
@{KEYS}=  Down  Down  Down  Down  Down  Down

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/7_EpisodePickerScreen.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Verify is details screen loaded
    Send key  Select  4
    Verify is screen loaded    ${DetailsParams}

Verify is Categort list started
    Send key  Select  4
    Verify is screen loaded    ${CategoryParams}

Verify is Season1 is focused
    Verify is screen loaded    ${Season1Params}

Verify is Season2 is focused
    Send keys  ${KEYS}
    Verify is screen loaded    ${Season2Params}