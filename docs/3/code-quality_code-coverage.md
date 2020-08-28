# Code Coverage:
	Code coverage is the percentage of code which is covered by automated tests. Code coverage measurement simply determines which statements in a body of code have been executed through a test run, and which statements have not.

## Exercise:
You will be using Mocha and Istanbul to test the code coverage in the provided code. 

[document](../../examples/codeQuality/javascript/simple.js ':include :type=code javascript')

Please perform the following:

    * Install Node
    * Install Mocha
    * Install Chai
    * Install Istanbul

Inside your package.json file under the “scripts” section, add a section called “coverage” and use the following command “nyc --reporter=text npm run test”
    * Note: Make sure that your “test” section is running the following command “mocha ./simpleTest.js