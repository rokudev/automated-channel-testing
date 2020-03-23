########################################################################
# Copyright 2020 Roku, Inc.
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
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}
Library  Collections


*** Variables ***
${channel_code}  dev
&{GridData}=  using=tag  value=GridView
@{GridArray}=  &{GridData}
&{GridParams}=  elementData=${GridArray}
&{PosterData}=   using=attr  attribute=name  value=poster
@{PosterArray}=  &{PosterData}
&{PosterParams}=  elementData=${PosterArray}
${poster1}=  https://roku-blog.s3.amazonaws.com/developer/files/2017/04/Roku-Recommends-thumbnail.png
${poster2}=  https://blog.roku.com/developer/files/2016/10/twitch-poster-artwork.png

*** Test Cases ***
Verify is channel launched
    Launch the channel  ${channel_code}
    Verify is channel loaded    ${channel_code}    
    
Verify is initial screen loaded
    Verify is screen loaded    ${GridParams}

Verify posters
    @{elements}=  Get elements   ${PosterParams}  4
    ${realPoster1}=  Get attribute  ${elements[0]}  uri
    Run Keyword If  "${realPoster1}" != "${poster1}"   Fail
     ${realPoster2}=  Get attribute  ${elements[1]}  uri
    Run Keyword If  "${realPoster2}" != "${poster2}"   Fail