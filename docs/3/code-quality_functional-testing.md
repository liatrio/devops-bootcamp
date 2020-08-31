# Functional Testing
While unit testing tests small pieces (or units) of code, functional testing combines multiple units to test particular functionality of the code. Unlike unit testing, functional testing typically involves testing functionality that uses project dependencies, such as a connection to a database. This means that functional testing often takes longer and is more in depth than unit testing. 

## How to do functional testing
Functional tests are written from a user’s perspective. That is, you should be testing that the system performs how you (and the user) expect. Since functional testing is highly dependent on the use cases for the application, how you write functional tests will vary widely. 

## Example
One popular tool for functional testing website and web applications is Selenium. Selenium is a portable framework for browser automation. Selenium allows you to write declarative tests that programmatically interact with a website, checking functionality. This can be extremely helpful automating tests, allowing the same tests to be performed using all popular browsers. 

An example of a use-case for Selenium would be testing a web application’s login form. Where a unit test might check that the email field only accepts valid email addresses, a functional test using Selenium might check that logging in with a valid email and password results in the user being logged in.

## Exercise:

We will use Selenium and Python to run a couple functional tests on the bootcamp website. 

1. Install a Selenium webdriver of your choice
    * Firefox - `brew install geckodriver`
    * Chrome - `brew cask install chromedriver`
    * Safari - built in
2. Install the [Selenium Python bindings](https://pypi.org/project/selenium/)

    Read through the [Selenium Getting Started page](https://selenium-python.readthedocs.io/getting-started.html) to get a brief introduction on how to use Selenium.

3. Use the starter code provided to test
    * The link in the bottom right corner of the homepage takes you to [the bootcamp introduction](https://devops-bootcamp.liatr.io/#/1-introduction/overview)
    * The sidebar toggle button properly hides the sidebar

[document](https://raw.githubusercontent.com/PaulDHenson/devops-bootcamp/master/examples/codeQuality/selenium-frame.py ':include :type=code python')

While these are fairly trivial functional tests, they serve to show in general how they differ from unit tests.



