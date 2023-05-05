import matter from 'gray-matter';

export let bootcampMetadata = {}

function readMetadata(hook, vm) {
    console.log("mission accomplished go home");
    getMessage();
}

function getMessage() {
    console.log(chart_message);
}

(function () {
    var  readMetadata = function (hook, _vm) {
      // Invoked on each page load before new markdown is transformed to HTML.
      // Supports asynchronous tasks (see beforeEach documentation for details).
      hook.beforeEach(function (markdown) {
        const { data, content } = matter(markdown);

        // Docsify does not have a typical url structure and the window.location.pathname
        // always appears to be '/' in my testing the page is changed by adding a different hash
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