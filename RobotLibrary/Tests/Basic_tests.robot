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
Documentation  Basic smoke tests
Variables  ./../Library/variables.py 
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}
Library  Collections


*** Variables ***
${channel_code}  dev
&{DATA2}=  using=text  value=Barack Gates, Bill Obama
@{DATA2Array}=  &{DATA2}
&{Params2}=  elementData=${DATA2Array}
&{DATA3}=  using=text  value=Please enter your username
@{DATA3Array}=  &{DATA3}
&{Params3}=  elementData=${DATA3Array}
&{DATA4}=  using=text  value=Please enter your password
@{DATA4Array}=  &{DATA4}
&{Params4}=  elementData=${DATA4Array}
@{KEYS}=   down  down  down  down  select
&{DATA5}=  using=text  value=Authenticate to watch
@{DATA5Array}=  &{DATA5}
&{Params5}=  elementData=${DATA5Array}

*** Test Cases ***
Channel should be launched
    Side load  ../sample/channel.zip   rokudev   aaaa
    Verify is channel loaded    ${channel_code}

Check if details screen showed
    Send key  select  4
    Verify is screen loaded    ${Params2}

Check if playback started
    ${status}  ${value}=  Run Keyword And Ignore Error  Verify is screen loaded  ${Params5}  2 
    Run keyword if   "${status}"=="PASS"  Do auth
    ...  ELSE  Send key  select
    Verify is playback started  20  2

*** Keywords ***
Do auth
    Send key  select
    Verify is screen loaded   ${Params3}
    Send word  user
    Send keys  ${KEYS}
    Verify is screen loaded   ${Params4}
    Send word  pass
    Send keys  ${KEYS}
