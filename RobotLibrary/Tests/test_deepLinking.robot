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
    Verify is channel loaded  ${channel_code}

Verify is playback started
    Verify is playback started  25  2