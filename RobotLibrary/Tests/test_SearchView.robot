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
&{SearchViewData}=  using=tag  value=SearchView
@{SearchViewArray}=  &{SearchViewData}
&{SearchViewParams}=  elementData=${SearchViewArray}
&{PlaceholderTextData}=  using=text  value=Enter search term
@{PlaceholderArray}=  &{PlaceholderTextData}
&{PlaceholderParams}=  elementData=${PlaceholderArray}
&{RowData}=  using=tag  value=Row
@{RowDataArray}=  &{RowData}
&{RowParams}=  elementData=${RowDataArray}
&{ParentTagData}=  using=tag  value=TextEditBox
@{ParentArray}=  &{ParentTagData}
&{ElementTagData}=  using=tag  value=Label
@{ElementArray}=  &{ElementTagData}
&{LabelParams}=  elementData=${ElementArray}  parentData=${ParentArray}
${input}=  hello@1 r~

*** Test Cases ***
Verify is channel launched
    Side load  ../channels/SearchView.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}    

Verify is search screen loaded
    Verify is screen loaded    ${SearchViewParams}

Verify search input
    Verify is screen loaded    ${PlaceholderParams}
    Send word  ${input}
    &{label}=  Get Element   ${LabelParams}
    ${text}=  Get attribute  ${label}  text
    Run Keyword If  "${text}" != "${input}"   Fail