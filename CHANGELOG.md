# **Roku Automated Channel Testing**

## v.2.2.0 (April 2021)

### Bug Fixes

 * Fixed submitting '@' symbol over "Send word" keyword.
 * Fixed "Get player info" keyword failing when called before playback starts
 * Adjusted sample tests for Robot and JS libraries (added **Sideload** command to automate sideloading of sample channels used for Robot and JavaScript sample tests).

### Features

 * Added pre-built WebDrivers for iOS, Linux, and Windows. 
 * Added option for installing Python version of Robot library as a local Python package.

## v.2.1.0 (July 2020)

### Features

- Added "Get child nodes" method for Robot and JavaScript libraries.
- Updated WebDriver **/element/active** endpoint to consistently return correct element.
- Updated WebDriver **/elements** endpoint to returns correct elements when multiple locators are specified.
- Added and updated sample tests for Robot and JavaScript libraries.

## v.2.0.0 (March 2020)

### Features

- JavaScript client.
- Channel side loading.
- Input deep linking.
- Timers.
- Multiple device support for Robot and JavaScript libraries.

## v.1.0.0 (December 2019)

### Features

- Initial Release
- Initial version of a Roku WebDriver.
- Key press simulation.
- Grab UI elements.
- Current app metadata.