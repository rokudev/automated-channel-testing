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
&{OverhangData}=  using=tag  value=Overhang
@{OverhangArray}=  &{OverhangData}
&{OVerhangParams}=  elementData=${OverhangArray}
&{GridParams}=  elementData=${GridArray}
&{DetailsData}=  using=tag  value=DetailsView
@{DetailsArray}=  &{DetailsData}
&{DetailsParams}=  elementData=${DetailsArray}
&{PosterData}=  using=tag  value=Poster
&{LabelData}=   using=text  value=Live Gaming
&{IndexData}=   using=attr  attribute=index  value=1
@{LabelArray}=  &{LabelData}  &{IndexData}
@{ParamArray}=  &{PosterData}

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/7_EpisodePickerScreen.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Verify focused element on grid screen
    Send key  Down  4
    &{focusedEl}=  get focusedElement
    @{Nodes}=  Get child nodes  ${focusedEl}  ${ParamArray}
    Log  ${Nodes}
    ${uri}=  Get attribute  @{Nodes}[${0}]  uri
    Run keyword if  '${uri}'!='https://blog.roku.com/developer/files/2016/10/twitch-poster-artwork.png'  Fail

Verify is details screen loaded
    Send key  Select  2
    Verify is screen loaded    ${DetailsParams}

Verify title on details screen
    &{Overhang}=  get element  ${OVerhangParams}
    @{Nodes}=  Get child nodes  ${Overhang}  ${LabelArray}
    Log  ${Nodes}
    ${count}=    Get length    ${Nodes}
    should be equal as numbers  ${count}  1