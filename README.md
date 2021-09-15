# Automated channel testing

Roku channel developers can use Roku's test automation software to write and execute test cases, including channel purchasing, performance, deep linking, and other certification-related testing. Roku provides custom Selenium-based WebDriver APIs for sending commands to launch channels, send keypresses, and check whether a specific screen has been loaded. Channels can use the WebDriver APIs to control a Roku device, while using a test framework or programming language to create, run, log, and record test cases. To make automated testing even easier, this software includes a library for creating Robot framework-compliant test cases.

## Overview

Test cases can be written with the Roku Robot Framework Library (or another test framework) or a programming language such as Python, JavaScript, or Go. The test scripts send command requests to the Roku WebDriver via the [JSON Wire Protocol](https://github.com/SeleniumHQ/selenium/wiki/JsonWireProtocol). 

The Roku WebDriver includes an HTTP server that receives the command requests and an [External Control Protocol (ECP)](https://developer.roku.com/docs/developer-program/debugging/external-control-api.md) client that translates them into ECP requests, which it then routes to the target device. The device returns the result of the command and XML data back to the Roku WebDriver, which then passes this information back to the test script as a JSON object (via the WebDriver client application).

## Getting started

To build, configure, and test the Roku WebDriver and Roku Robot Framework Library, follow these steps:

**Important note:** You can use already built version of Roku webDriver for Windows, Linux or Mac(see bin folder) instead of steps 1-7

1. [Download](https://golang.org/dl/) and install the Go programming language (the Roku WebDriver server is implemented as a Go application). 


2. Clone the [Roku automated channel testing repository](https://gitlab.eng.roku.com/developer_web_tools/roku-automated-channel-testing) or download it as a zip file.


3. Set the "GOPATH" environment variable to the path of the **roku-automated-channel-testing** folder ($APP_PATH).


4. Install the following dependencies ([mux](https://github.com/gorilla/mux/blob/master/README.md) is a URL router and dispatcher; [logrus](https://github.com/sirupsen/logrus/blob/master/README.md) is a structured logger):

        cd <path>/roku-automated-channel-testing-develop/src
        go get github.com/gorilla/mux
        go get github.com/sirupsen/logrus

5. Build the Roku WebDriver project:

        go build main.go

6.  Run the **webDriver** executable to start the Roku WebDriver server. Default port is 9000 but you can run webdriver on a specific port:

        ./RokuWebDriver_mac 9001


7. Test the Roku WebDriver server following these steps:

   a. Install the [**assert**](https://godoc.org/github.com/stretchr/testify/assert) package, which provides testing tools to be used with Go applications.

        go get github.com/stretchr/testify/assert

   b. Test the ECP client:

        go test ecpClient

   c. Test the HTTP server (the host is "localhost"; the port used is 9000):

        go test httpServer
        
### Steps for Robot library

8. Run Roku's Python-based sample WebDriver client application following these steps: 

   a. Download and install python: https://www.python.org/downloads.
   Important: python version must be 3.7 or higher. Also, this version must be set as default version of python.
   
   b. Download and install pip: https://pypi.org/project/pip

   c. Install the [**requests**](https://pypi.org/project/requests) HTTP library for Python, which enables the sample client application to send HTTP 1.1 requests:

        python -m pip install requests

   d. Sideload the sample channel (**channel.zip**) included in the **roku-auotmated-channel-testing/sample** directory. 

   e. Run the sample Web driver client application: 

        python <path>/roku-automated-channel-testing-develop/sample/script/main.py  <device_ip>


9. Configure and test the Roku Robot Framework Library following these steps:

   a. Install the dependencies listed in the **/roku-automated-channel-testing-develop/RobotLibrary/requirements.txt** file:

        python -m pip install -r requirements.txt

   b. Sideload the sample Robot framework test  (**channel.zip**) included in the **roku-auotmated-channel-testing/RobotLibrary** directory.

   c. Visit RobotLibrary\Library\variables.py file and update it with IP address of device and path to webDriver

   d. Run the sample basic Robot test case (test cases must be run from the **RobotLibrary** folder to be completed successfully):

       cd RobotLibrary
       python -m robot.run --outputdir Results Tests/Basic_tests.robot
   
   > When running the tests and samples, you can provide the IP address and WebDriver server path as variables in the console as demonstrated in the following example: 

       python -m robot.run --outputdir Results --variable ip_address:192.168.1.94 --variable server_path:D:/projects/go/webDriver/src/main.exe  Tests/Basic_tests.robot
10. View the generated test case report and log, which are stored in the specified output directory.

### Install Robot library as package
Requirements: python( version 3.7 or higher), pip

       python pip install <local_path>/roku-automated-channel-testing/RobotLibrary

After that you will be able to import RobotLibrary to robot test file by the following way:

       Library  Library.RobotLibrary  ${ip_address}  ${timeout}  ${pressDelay}  ${server_path}

   
### Steps for JavaScript library

8. Configure and test the Roku JavaScript Library following these steps:
 
   a. Download and install nodeJs: https://nodejs.org/en/

   b. Download and install yarn: https://classic.yarnpkg.com/en/docs/install
   
   c. Install the dependencies listed in the **/roku-automated-channel-testing-develop/jsLibrary/package.json** file:

        yarn install
        
   d. Run the sample basic test case:
   
      1.) Using yarn:

        yarn test
        
    In this case the following line should be added to the package.json:

        "scripts": {
            "test": "mocha \"tests/test_basic.js\""
        },
    
      2.) Using mocha:

        mocha  tests/test_basic.js
        
    In this case mocha must be installed globally:

        yarn global add mocha
        
## Multiple device support

Multiple device support provides opportunity to run specified test file(s) on a few devices in parallel. Both RobotLibrary and JsLibrary support multiple device test run. 

Important: channel side loading("sideLoad" method in jsLibrary and "Side load" keyword in RobotLibrary) can cause an error in multiple device case so be careful when use it in multiple device tests.

### JavaScript multiple device support

1.) Install mocha and mochawesome globally:

        yarn global add mocha
        
        yarn global add mochawesome

2.) Run the following console cmd to start script:

        node multipleDevices/multi.js   <path_to_config.json>


config.json contains details needed for multi.js script run.  It has the following structure:

    { "devices": {
        "Austin": {
            "ip_address": <string>,
            "timeout":  <number>,
            "pressDelay": <number>
        },
       "Cooper": {
            "ip_address": <string>,
            "timeout":  <number>
        }
    },
    "server_path": <string>,
    "test": <string>,
    "outputdir": <string> 
    }



Fields:

1.) devices(required): Contains "key-value" pairs where:

    key - name of device(can be what ever user want);

    value - object with fields  which value should be used in tests for this device.

2.) server_path": path to webDriver launch file. multi.js will start one  instance of webDriver for all test executions.

3.) "test': path to javeScript test which use want to run on specified devices.

4.)  "outputdir": path to folder where report/log files will be stored.



### Robot Framework multiple device support

Run the following console command:

        cd  RobotLibrary
        <path_to_python>/python     multipleDevices/multi.py   <path_to_config.json>


config.json contains details needed for multi.py script run.  It has the following structure:

    {
    "devices": {
        "Austin": {
            "ip_address": <string>,
            "timeout":  <number>,
            "pressDelay": <number>
        },
       "Cooper": {
            "ip_address": <string>,
            "timeout":  <number>
        }
    },
    "server_path": <string>,
    "test": <string>,
    "outputdir": <string> 
    }




Fields:

1.) devices(required): Contains "key-value" pairs where:

    key - name of device(can be what ever user want) ;

    value - object with fields  which value should be used in tests for this device. this fields will override fields in test.

2.) server_path": path to webDriver launch file. multi.py will start one  instance of webDriver for all test executions. 

3.) "test': path to robot test which use want to run on specified devices.

4.)  "outputdir": path to folder where report/log files will be stored.




       
