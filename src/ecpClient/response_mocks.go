///////////////////////////////////////////////////////////////////////////
// Copyright 2019 Roku, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//////////////////////////////////////////////////////////////////////////

package ecpClient

var SuccessResponseMock = `{}`

var AppResponseMock =`<?xml version="1.0" encoding="UTF-8" ?>
<active-app>
    <app id="test" subtype="test" type="test" version="test">test</app>
</active-app>`

var AppsResponseMock = `<?xml version="1.0" encoding="UTF-8" ?>
<apps>
    <app id="test" type="test" version="test" subtype="test">test</app>
    <app id="test2" subtype="test2" type="test2" version="test2">test2</app>
</apps>`

var deviceInfoResponse =`<?xml version="1.0" encoding="UTF-8" ?>
<device-info>
    <udn>test</udn>
    <serial-number>test</serial-number>
    <device-id>test</device-id>
    <advertising-id>test</advertising-id>
    <vendor-name>test</vendor-name>
    <model-name>test</model-name>
    <model-number>test</model-number>
    <model-region>test</model-region>
    <is-tv>test</is-tv>
    <is-stick>test</is-stick>
    <supports-ethernet>test</supports-ethernet>
    <wifi-mac>test</wifi-mac>
    <wifi-driver>test</wifi-driver>
    <ethernet-mac>test</ethernet-mac>
    <network-type>test</network-type>
    <network-name>test</network-name>
    <friendly-device-name>test</friendly-device-name>
    <friendly-model-name>test</friendly-model-name>
    <default-device-name>test</default-device-name>
	<user-device-name>test</user-device-name>
	<user-device-location>test</user-device-location>
    <build-number>test</build-number>
    <software-version>test</software-version>
    <software-build>test</software-build>
    <secure-device>test</secure-device>
    <language>test</language>
    <country>test</country>
    <locale>test</locale>
    <time-zone-auto>test</time-zone-auto>
    <time-zone>test</time-zone>
    <time-zone-name>test</time-zone-name>
    <time-zone-tz>test</time-zone-tz>
    <time-zone-offset>test</time-zone-offset>
    <clock-format>test</clock-format>
    <uptime>test</uptime>
    <power-mode>test</power-mode>
    <supports-suspend>test</supports-suspend>
    <supports-find-remote>test</supports-find-remote>
    <find-remote-is-possible>test</find-remote-is-possible>
    <supports-audio-guide>test</supports-audio-guide>
    <supports-rva>test</supports-rva>
    <developer-enabled>test</developer-enabled>
    <keyed-developer-id>test</keyed-developer-id>
    <search-enabled>test</search-enabled>
    <search-channels-enabled>test</search-channels-enabled>
    <voice-search-enabled>test</voice-search-enabled>
    <notifications-enabled>test</notifications-enabled>
    <notifications-first-use>test</notifications-first-use>
    <supports-private-listening>test</supports-private-listening>
    <headphones-connected>test</headphones-connected>
    <supports-ecs-textedit>test</supports-ecs-textedit>
    <supports-ecs-microphone>test</supports-ecs-microphone>
    <supports-wake-on-wlan>test</supports-wake-on-wlan>
    <has-play-on-roku>test</has-play-on-roku>
    <has-mobile-screensaver>test</has-mobile-screensaver>
    <support-url>test</support-url>
    <grandcentral-version>test</grandcentral-version>
    <trc-version>test</trc-version>
    <trc-channel-version>test</trc-channel-version>
    <davinci-version>test</davinci-version>
</device-info>`

var UiResponseMock = `<?xml version="1.0" encoding="UTF-8" ?>
	<MainScene bounds="{0, 0, 1920, 1080}" children="0" extends="BaseScene" focusable="true">
		<Poster bounds="{0, 0, 1920, 1080}" index="0" loadStatus="3"/>
	</MainScene>
</app-ui>`