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
    //   // Invoked one time when docsify script is initialized
    //   hook.init(function () {
    //     // ...
    //   });

    //   // Invoked one time when the docsify instance has mounted on the DOM
    //   hook.mounted(function () {
    //     // ...
    //   });

      // Invoked on each page load before new markdown is transformed to HTML.
      // Supports asynchronous tasks (see beforeEach documentation for details).
      hook.beforeEach(function (markdown) {
        // yamlFront is browser bundle I would like to replace with the npm package gray-matter.
        // Need to figure out how to incorporate webpack and update our deploy process.
        // Could stay for the time being as it will just be used to strip out the front matter.
        var result = yamlFront.loadFront(markdown);
        // Strip out the front-matter
        return result.__content;
      });

    //   // Invoked on each page load after new markdown has been transformed to HTML.
    //   // Supports asynchronous tasks (see afterEach documentation for details).
    //   hook.afterEach(function (html) {
    //     // ...
    //     return html;
    //   });

    //   // Invoked on each page load after new HTML has been appended to the DOM
    //   hook.doneEach(function () {
    //     // ...
    //   });

    //   // Invoked one time after rendering the initial page
    //   hook.ready(function () {
    //     // ...
    //   });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(readMetadata, $docsify.plugins || []);
  })();