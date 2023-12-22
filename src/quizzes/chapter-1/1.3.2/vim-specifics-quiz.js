const rawQuizdown = `
---
shuffleQuestions: true
shuffleAnswers: true
---

# A file has the following contents:<br>
\`text tex|t text\`<br>
The '|' signifies where the cursor is at<br>
Vim is in normal mode<br><br>
What are the contents of the file after the following commands are entered?<br>
\`w\`<br>
\`a\`<br>
\`NEW\`

1. [x] text text tNEWext
1. [ ] text text NEWtext
	> Look above to recall what the \`w\` and \`a\` shortcuts do
1. [ ] text text textNEW
	> Look above to recall what the \`w\` and \`a\` shortcuts do
1. [ ] text textNEW text
	> Look above to recall what the \`w\` and \`a\` shortcuts do
1. [ ] text texNEWt text
	> Look above to recall what the \`w\` and \`a\` shortcuts do

# A file has the following contents:<br>
\`text\`<br>
Vim is in normal mode<br><br>
What are the contents of the file after the following commands are entered?<br>
\`yy\`<br>
\`p\`<br>
\`p\`

1. [x] text<br>text<br>text
1. [ ] text text text
	> Look above to recall what the \`yy\` and \`p\` shortcuts do
1. [ ] texttexttext
	> Look above to recall what the \`yy\` and \`p\` shortcuts do

# While in normal mode, which of the following commands will yield different results depending on where in the line your cursor is at?

- [x] a
- [x] w
- [x] e
- [x] b
- [ ] $
	> This will jump your cursor to the end of the line no matter where it's at
- [ ] 0
	> This will jump your cursor to the beginning of the line no matter where it's at
- [ ] o
	> This will insert a new line below the current line no matter where the cursor is at
- [ ] O
	> This will insert a new line above the current line no matter where the cursor is at
- [ ] A
	> This will jump your cursor to the end of the line (and enter insert mode) no matter where it's at
- [ ] yy
	> This will copy the current line no matter where your cursor is at
`;

export { rawQuizdown }