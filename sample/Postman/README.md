# WebDriver Postman collection

1. Side load [Roku Recommends](../../channels/Roku_Recommends.zip) sample channel. 
2. [WebDriver_endpoints.json](./WebDriver_endpoints.json) file should be imported as Postman collection.
3. Start instance of [Roku WebDriver.](../../README.md#Getting started)
4. Open "create session" request and populate `ip` field in request body with your Roku Device IP address.
5. Execute "create session" request to create new session.
6. Execute any other requests
7. Execute "delete session" request to remove session.


