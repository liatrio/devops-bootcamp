import * as matter from 'gray-matter';

(function () {
  var renderQuizdown = function (hook, _vm) {
    let rawQuizdown;
    hook.beforeEach(function (markdown) {
      const { data, content } = matter(markdown);
      rawQuizdown = data.quizdown
      return content;
    });

    hook.doneEach(function() {
      const config = {
        'locale': 'en', 
        'shuffleAnswers': true,
      };
      let node = document.createElement('div');
      node.id = 'quizdown';
      document.body.appendChild(node);
      createApp(rawQuizdown, node, config);
    });
  };

  // Add plugin to docsify's plugin array
  $docsify = $docsify || {};
  $docsify.plugins = [].concat(renderQuizdown, $docsify.plugins || []);
})();