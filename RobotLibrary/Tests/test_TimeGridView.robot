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
&{TimeGridViewData}=  using=tag  value=TimeGridView
@{TimeGridViewArray}=  &{TimeGridViewData}
&{TimeGridViewParams}=  elementData=${TimeGridViewArray}
&{ParentTagData}=  using=tag  value=ChannelRow
&{ParentIndexData}=  using=attr  attribute=index  value=1
@{ParentArray}=  &{ParentTagData}  &{ParentIndexData}
&{ElementTagData}=  using=tag  value=Label
@{ElementArray}=  &{ElementTagData}
&{LabelParams}=  elementData=${ElementArray}  parentData=${ParentArray}
${LabelValue}=  KTVK-SD 3.2

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/TimeGridView.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is search screen loaded
    Verify is screen loaded    ${TimeGridViewParams}

Verify Label text
    &{element}=  GetElement  ${LabelParams}
    ${value}=  Get attribute  ${element}  text
    Run Keyword If  "${value}" != "${LabelValue}"  Fail
    