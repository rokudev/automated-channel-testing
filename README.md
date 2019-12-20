# Automated channel testing

Roku channel developers can use Roku's test automation software to write and execute test cases, including channel purchasing, performance, deep linking, and other certification-related testing. Roku provides custom Selenium-based WebDriver APIs for sending commands to launch channels, send keypresses, and check whether a specific screen has been loaded. Channels can use the WebDriver APIs to control a Roku device, while using a test framework or programming language to create, run, log, and record test cases. To make automated testing even easier, this software includes a library for creating Robot framework-compliant test cases.

> Roku's test automation tools require Roku OS 9.1 or higher. 

## Overview

Test cases can be written with the Roku Robot Framework Library (or another test framework) or a programming language such as Python, Java, or Go. The test scripts send command requests to the Roku WebDriver via the [JSON Wire Protocol](https://github.com/SeleniumHQ/selenium/wiki/JsonWireProtocol). 

The Roku WebDriver includes an HTTP server that receives the command requests and an [External Control Protocol (ECP)](https://developer.roku.com/docs/developer-program/debugging/external-control-api.md) client that translates them into ECP requests, which it then routes to the target device. The device returns the result of the command and XML data back to the Roku WebDriver, which then passes this information back to the test script as a JSON object (via the WebDriver client application).

## Getting started

To build, configure, and test the Roku WebDriver and Roku Robot Framework Library, follow these steps:  

1. [Download](https://golang.org/dl/) and install the Go programming language (the Roku WebDriver server is implemented as a Go application). 


2. Clone this repository or download it as a zip file.


3. Set the "GOPATH" environment variable to the path of the **automated-channel-testing-master** folder ($APP_PATH).


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

   a. Download and install python: https://www.python.org/downloads.

   b. Install the [**requests**](https://pypi.org/project/requests) HTTP library for Python, which enables the sample client application to send HTTP 1.1 requests:

        python -m pip install requests

   c. Sideload the sample channel (**channel.zip**) included in the **automated-channel-testing-master/sample** directory. 

   d. In line 19 of the sample WebDriver client application (**automated-channel-testing-master/sample/script/main.py**), change the IP address of the **web_driver** variable to the IP address of your Roku device. 

   e. Run the sample Web driver client application: 

        python <path>/automated-channel-testing-master/sample/script/main.py


9. Configure and test the Roku Robot Framework Library following these steps:

   a. Install the dependencies listed in the **/automated-channel-testing-master/RobotLibrary/requirements.txt** file:

        python -m pip install -r requirements.txt

   b. Sideload the sample Robot framework test  (**channel.zip**) included in the **automated-channel-testing-master/RobotLibrary** directory.

   c. Run the sample basic Robot test case (test cases must be run from the **RobotLibrary** folder to be completed successfully):

       cd RobotLibrary
       python -m robot.run --outputdir Results Tests/Basic_tests.robot
   
   > When running the tests and samples, you can provide the Roku device IP address and WebDriver server path as variables in the console as demonstrated in the following example: 

       python -m robot.run --outputdir Results --variable ip_address:192.168.1.94 --variable server_path:D:/projects/go/webDriver/src/main.exe  Tests/Basic_tests.robot

10. View the generated test case report and log, which are stored in the specified output directory.
