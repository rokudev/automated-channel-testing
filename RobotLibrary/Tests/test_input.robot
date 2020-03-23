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

Verify is playback started
    Input deep linking data  ${channel_code}   ${content_id}  ${mediaType}
    Verify is playback started  25  2