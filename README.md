# Automated channel testing

Roku channel developers can use Roku's test automation software to write and execute test cases, including channel purchasing, performance, deep linking, and other certification-related testing. Roku provides custom [Selenium](https://selenium.dev)-based [WebDriver APIs](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/web-driver.md) for sending commands to launch channels, send keypresses, and check whether SceneGraph components are present on the screen. Channels can use the WebDriver APIs to control a Roku device, while using a test framework or programming language to create, run, log, and record test cases. To make automated testing even easier, Roku provides [Robot](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/robot-framework-library.md) and [JavaScript](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/javascript-library.md) libraries, which support running tests on multiple devices at the same time. 

Executing test automation allows channels to run state-driven UI testing for a number of scenarios. For example, channels can create a test case that installs a channel and launches it with a specific contentID and mediaType to verify that deep linking works. Authenticated channels can execute more complex test cases such as launching a channel, trying to play content before authenticating the user, entering valid/invalid credentials, and then trying to play content again. 

All test cases can be run simultaneously on multiple Roku devices. This is useful for testing channel performance across different models with varying RAM and CPU. It is especially important for certification testing, which requires channels to meet [performance criteria](https://developer.roku.com/docs/developer-program/certification/certification.md#3-performance) that varies for different device types.

Implementing automated testing speeds up channel development by reducing the number of manual UI tests that need to be run for simple to complex test cases.

> Roku's test automation tools require Roku OS 9.1 or higher.

> To test production channels with the Roku test automation tools, you must [package the channel](https://developer.roku.com/docs/developer-program/publishing/packaging-channels.md#rekeying) on your Roku device using the same Roku developer account linked to the production version of the channel.

## Overview

Test cases can be written with the [Roku Robot Framework Library](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/robot-framework-library.md), [Roku JavaScript library](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/javascript-library.md), another test framework, or a programming language such as Python, JavaScript, or Go. The test scripts send command requests to the [Roku WebDriver](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/web-driver.md) via the [JSON Wire Protocol](https://github.com/SeleniumHQ/selenium/wiki/JsonWireProtocol). 

The Roku WebDriver includes an HTTP server that receives the command requests and an [External Control Protocol (ECP)](https://developer.roku.com/docs/developer-program/debugging/external-control-api.md) client that translates them into ECP requests, which it then routes to the target device. 

The device returns the result of the command and XML data back to the Roku WebDriver, which then passes this information back to the test script as a JSON object (via the WebDriver client application).

## Channel UI testing

The Roku WebDriver includes a set of [APIs](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/web-driver.md#roku-webdriver-apis) that enable developers to run state-driven UI tests. For example, the RokuWebDriver has an [**element** API](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/web-driver.md#POST-v1/session/:sessionId/element) to check if a specific SceneGraph component is present on the screen in order to  determine whether a specific screen has been loaded. In addition, the Roku Robot Framework library has [keywords](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/robot-framework-library.md#keywords) that are mapped to the Roku WebDriver APIs so that developers can execute channel UI-based test cases with the Robot Framework. 

## Getting started

To build, configure, and test the [Roku WebDriver](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/web-driver.md), follow these steps:  

1. [Download](https://golang.org/dl/) and install the Go programming language (the Roku WebDriver server is implemented as a Go application). 


2. Clone this repository or download it as a zip file.


3. Set the "GOPATH" environment variable to the path of the **automated-channel-testing** folder ($APP_PATH).


4. Install the following dependencies ([mux](https://github.com/gorilla/mux/blob/master/README.md) is a URL router and dispatcher; [logrus](https://github.com/sirupsen/logrus/blob/master/README.md) is a structured logger):

        cd <path>/automated-channel-testing-master/src
        go get github.com/gorilla/mux
        go get github.com/sirupsen/logrus


5. Build the Roku WebDriver project:

        go build main.go


6.  Run the **main** executable in the **/automated-channel-testing-master/src** folder to start the Roku WebDriver server. 


7. Test the Roku WebDriver server following these steps:

   a. Install the [**assert**](https://godoc.org/github.com/stretchr/testify/assert) package, which provides testing tools to be used with Go applications.

        go get github.com/stretchr/testify/assert

   b. Test the ECP client:

        go test ecpClient

   c. Test the HTTP server (the host is "localhost"; the port used is 9000):

        go test httpServer


8. Run Roku's Python-based sample WebDriver client application following these steps: 

   a. [Download](https://www.python.org/downloads) and install Python 3.7 (or higher). Set the version you install as the default version of Python on your computer.

   b. Download and install the Python package installer ([pip](https://pypi.org/project/pip)).

   c. Install the [**requests**](https://pypi.org/project/requests) HTTP library for Python, which enables the sample client application to send HTTP 1.1 requests:

       python -m pip install requests

   d.  [Sideload](https://developer.roku.com/docs/developer-program/getting-started/developer-setup.md#step-2-accessing-the-development-application-installer) the sample channel (**channel.zip**) included in the **automated-channel-testing-master/sample** directory. 

   e. Run the sample Web driver client application. Include the IP address of your Roku device as an argument. 

        python <path>/automated-channel-testing-master/sample/script/main.py <device-ip-address>

### Installing and testing the Robot Framework Library 

To install the [Roku Robot Framework Library](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/robot-framework-library.md) and test it on one or more devices, follow these steps:  

1. Install the dependencies listed in the **/automated-channel-testing-master/RobotLibrary/requirements.txt** file:

        python -m pip install -r /automated-channel-testing-master/RobotLibrary/requirements.txt

2. Update line 41 of the **/automated-channel-testing-master/RobotLibrary/Tests/Basic_tests.robot** file with the password of your Roku device.

3. Run the sample basic Robot test case on a single device. When running the Robot tests and samples, you must run them from the **RobotLibrary** folder. You must also provide the Roku device IP address and WebDriver server path as variables in the console as demonstrated in the following example:

        cd RobotLibrary
        python -m robot.run --outputdir Results --variable ip_address:192.168.1.94 --variable server_path:D:/projects/go/webDriver/src/main.exe  Tests/Basic_tests.robot

     > Alternatively, you can the hard code the Roku device IP address and WebDriver server path variables in the **roku-automated-channel-testing-develop/RobotLibrary/Library/variables.py** file. 

4. View the generated test case report, which is stored in the specified output directory (**/automated-channel-testing-master/RobotLibrary/Results** by default).

5. Run the sample basic Robot test case on multiple devices following these steps:

   a. Update the JSON configuration file (**config.json**) in the **automated-channel-testing-master/RobotLibrary/multipleDevices** directory, which contains the Roku devices to be used for testing, the Web driver server path, test case, and the output directory.

   Each Roku device is an object that has an arbitrary name and a key-value pair with the device's IP address. Key-value pairs may also be provided for the timeout and keypress delay to be used for the test on that device (these override the global test values specified in the Robot test case).

   The syntax of the **config.json** file is as follows:

        {
            "devices": {
                "Device 1 name": {
                    "ip_address": <string>,
                    "timeout":  <number>,
                    "pressDelay": <number>
                },
                "Device 2 name": {
                    "ip_address": <string>,
                    "timeout":  <number>
                }
            },
            "server_path": <string>,
            "test": <string>,
            "outputdir": <string> 
        }

   The following example demonstrates how to write the **config.json** file: 

        {
            "devices": {
                "Amarillo": {
                    "ip_address": "192.168.1.64",
                    "timeout":  20000,
                    "pressDelay": 2000
                },
                "Littlefield": {
                    "ip_address": 192.168.1.16,
                    "timeout":  25000,
                    "pressDelay": 1000
                }
            },
            "server_path": "/automated-channel-testing-develop/src/main,
            "test": "Tests/Basic_tests_multi_device.robot",
            "outputdir": "Results"
        }

   b.  [Sideload](https://developer.roku.com/docs/developer-program/getting-started/developer-setup.md#step-2-accessing-the-development-application-installer) the sample channel (**channel.zip**) in the **/automated-channel-testing-master/sample** folder.

   c. Run the following console command:

        cd RobotLibrary
        python multipleDevices/multi.py multipleDevices/config.json

   d. View the generated test case report and log for each device, which are stored in the specified output directory (**/automated-channel-testing-master/RobotLibrary/Results** by default).

### Installing and testing the Roku JavaScript Library 

To install the [Roku JavaScript  Library](https://developer.roku.com/docs/developer-program/dev-tools/automated-channel-testing/javascript-library.md) and test it on one or more devices, follow these steps: 

1. Download and install the [node.js](https://nodejs.org/en/) JavaScript runtime environment.

2. Download and install the [Yarn](https://classic.yarnpkg.com/en/docs/install) JavaScript package manager.

3. Install the dependencies listed in the **/automated-channel-testing-master/jsLibrary/package.json** file:

        yarn install

4. To use the [Mocha](https://mochajs.org/) JavaScript test framework and run tests on multiple devices, globally install Mocha and [Mochawesome](https://www.npmjs.com/package/mochawesome): 

        yarn global add mocha
        yarn global add mochawesome

5. Update the **/automated-channel-testing-master/jsLibrary/tests/test_basic.js** file with the following: 

   a. In line 20, update the WebDriver server path.

   b. In line 27, update the IP address to your Roku device. 

   c. In line 28, update the password. 

6. Run the sample basic JavaScript test case on a single device. When running the JavaScript tests and samples, you must run them from the **jsLibrary** folder

        yarn tests/test_basic.js

   To run the test using Mocha and report the test results with Mochawesome, enter the following command:

         mocha tests/test_basic.js --reporter mochawesome

7. View the generated test case report, which is stored in the **mochawesome-report** directory.

8. Run the sample basic JavaScript test case on multiple devices following these steps:

   a. Update the JSON configuration file (**config.json**) in the **/automated-channel-testing-master/jsLibrary/multipleDevices** directory, which contains the Roku devices to be used for testing, the Web driver server path, test case, and the output directory. 

   Each Roku device is an object that has an arbitrary name and a key-value pair with the device's IP address. Key-value pairs may also be provided for the timeout and keypress delay to be used for the test on that device (these override the global test values specified in the Robot test case).

   The syntax of the **config.json** file is as follows:

        {
            "devices": {
                "Device 1 name": {
                    "ip_address": <string>,
                    "timeout":  <number>,
                    "pressDelay": <number>
                },
                "Device 2 name": {
                    "ip_address": <string>,
                    "timeout":  <number>
                }
            },
            "server_path": <string>,
            "test": <string>,
            "outputdir": <string> 
        }

   The following example demonstrates how to write the **config.json** file: 

        {
            "devices": {
                "Amarillo": {
                    "ip_address": "192.168.1.64",
                    "timeout":  20000,
                    "pressDelay": 2000
                },
                "Littlefield": {
                    "ip_address": 192.168.1.16,
                    "timeout":  25000,
                    "pressDelay": 1000
                }
            },
            "server_path": "/roku-automated-channel-testing-develop/src/main.exe,
            "test": "multipleDevices/multiple_devices_test_basics.js",
            "outputdir": "Results"
        }

   b.  [Sideload](https://developer.roku.com/docs/developer-program/getting-started/developer-setup.md#step-2-accessing-the-development-application-installer) the sample channel (**channel.zip**) in the **/automated-channel-testing-master/sample** folder.

   c. Run the following console command:

        node multipleDevices/multi.js  config.json

   d. View the generated test case report and log for each device, which are stored in the specified output directory (**/automated-channel-testing-master/jsLibrary/Results** by default).
