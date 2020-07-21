*** Settings ***
Documentation  Deep Linking
Variables  ./../Library/variables.py 
Library  ./../Library/RobotLibrary.py  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}
Library  Collections


*** Variables ***
${channel_code}  dev
${content_id}=  12
${mediaType}=  movie

*** Test Cases ***
Side load
    Side load  ../channels/Roku_Recommends.zip   rokudev   aaaa
    Send key  Home

Verify is channel launched
    Launch the channel   ${channel_code}  ${content_id}  ${mediaType}
    Mark timer
    Verify is channel loaded  ${channel_code}

Verify is playback started (Deep linking)
    Verify is playback started

Verify is playback started quickly
     ${time_dl}=  Get timer
     Run Keyword If  ${time_dl} > 12000  Fail

Verify is playback started (input)
    Send key  Stop
    Sleep  2
    Mark timer
    Input deep linking data  ${channel_code}  ${content_id}  ${mediaType}
    Verify is playback started  25  2

Verify is playback(input) started quickly
     ${time_in}=  Get timer
     Run Keyword If  ${time_in} > 10000  Fail
