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
        const result = yamlFront.loadFront(markdown);

        // If we are loading the root page we need to store the master record
        if (window.location.pathname === '/') {
          const { __content, ...metaData } = result;
          bootcampMetadata = metaData;
        }
        // Strip out the front-matter
        return result.__content;
      });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(readMetadata, $docsify.plugins || []);
  })();