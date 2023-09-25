// quizzes.js

const config = {
  'locale': 'en', 
  'shuffleAnswers': true,
};

// Quiz content in quizdown format
const rawQuizdown = `
# This is awesome!

- [x] True
- [ ] False

# This is awesome!

- [x] True
- [ ] False
`;

export { config, rawQuizdown };

const config2 = {
  'locale': 'en', 
  'shuffleAnswers': true,
};

// Quiz content in quizdown format
const rawQuizdown2 = `
# This isn't awesome!

- [x] True
- [ ] False

# This isn't awesome!

- [x] True
- [ ] False
`;

export { config2, rawQuizdown2 };