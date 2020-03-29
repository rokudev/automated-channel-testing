# Importing and testing the Postman collection

Roku's test automation software includes a Postman collection that enables you to test the Roku WebDriver API calls and write test suites. To import the Postman JSON collection and use it to test the Roku WebDriver calls, follow these steps:

1. [Download](https://www.postman.com/downloads/) Postman. 

2. Verify that the Roku WebDriver server is running (to start the WebDriver, run the **main** executable in the **/automated-channel-testing-master/src** folder). 

3. [Sideload](/docs/developer-program/getting-started/developer-setup.md#step-2-accessing-the-development-application-installer) the sample channel (**channel.zip**) in the **/automated-channel-testing-master/sample** folder.

4. In Postman, import the **/automated-channel-testing-master/sample/Postman/WebDriver_endpoints** Postman collection.

5. Create a new session. To do this, click the **POST create session** request, update the IP address to your Roku device, and then click **Send**.

6. Execute the requests in the Postman collection to test the Roku WebDriver.

   ![roku-webdriver-postman-collection](https://image.roku.com/ZHZscHItMTc2/roku-webdriver-postman-collection.png)

7. When you have finished testing, send the **DEL delete session** request to remove the session.
