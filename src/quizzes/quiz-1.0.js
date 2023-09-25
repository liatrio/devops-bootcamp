console.log("testing");

const config = {
  'locale': 'en', 
  'shuffleAnswers': true,
};


// Quiz content in quizdown format
const rawQuizdown = `
# This is the quiz for Chapter 1.0?

- [x] True
- [ ] False

# This is awesome!

- [x] True
- [ ] False
`;

export {config, rawQuizdown}