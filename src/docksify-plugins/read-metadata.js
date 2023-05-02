import matter from 'gray-matter';

var chart_message ="foobar";
var bootcampMetadata = {}

function readMetadata(hook, vm) {
    console.log("mission accomplished go home");
    getMessage();
}

function getMessage() {
    console.log(chart_message);
}

(function () {
    var  readMetadata = function (hook, vm) {
      // Invoked on each page load before new markdown is transformed to HTML.
      // Supports asynchronous tasks (see beforeEach documentation for details).
      hook.beforeEach(function (markdown) {

        // yamlFront is browser bundle I would like to replace with the npm package gray-matter.
        // Need to figure out how to incorporate webpack and update our deploy process.
        // Could stay for the time being as it will just be used to strip out the front matter.
        const { data, content } = matter(markdown);
        //const result = yamlFront.loadFront(markdown);

        // Docsify does not have a typical url structure and the window.location.pathname always appears to be '/' in my testing
        // the page is changed by adding a different hash
        if (window.location.hash === '#/') {
          bootcampMetadata = data;
        }
        // Strip out the front-matter
        return content;
      });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(readMetadata, $docsify.plugins || []);
  })();