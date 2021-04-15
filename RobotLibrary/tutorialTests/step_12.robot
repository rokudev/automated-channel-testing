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
&{AdData}=  using=text  value=Ad 1 of 1
@{AdArray}=  &{AdData}
&{AdParam}=  elementData=@{AdArray}
&{LabelData}=  using=tag  value=LabelList
&{PlayData}=  using=text  value=Play
@{PlayArray}=  &{PlayData}
@{LabelArray}=  &{LabelData}
&{PlayParam}=  elementData=@{PlayArray}  parentData=@{LabelArray}
${content_id}=  12
${mediaType}=  movie

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
    Send keys  ${keys}  3
    Verify is screen loaded    ${DetailsParam}

Get description element and verify color
    &{descriptionElement}=  Get element   ${TextParam}
    ${color}=  Get attribute  ${descriptionElement}  color
    Run Keyword If  "${validColor}"!="${color}"  Fail

Verify is Ad playback started
   Send key  Select
   Verify is playback started  20  3
   Verify is screen loaded  ${AdParam}

Verify is correct Details screen opened after Back key
   Send key  Back  3
   Verify is screen loaded  ${PlayParam}

Verify deep linking
    Mark timer
    Launch the channel   ${channel_code}  ${content_id}  ${mediaType}
    Verify is channel loaded  ${channel_code}  
    Verify is playback started  20  3

Verify launch time
     ${time_in}=  Get timer
     Run Keyword If  ${time_in} > 40000  Fail