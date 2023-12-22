const rawQuizdown = `
---
shuffleAnswers: true
---

# If performed on a nonexistent file, all of the following commands will return a non-zero exit code except:

1. [x] touch
1. [ ] cat
1. [ ] mv
1. [ ] cp
1. [ ] rm

# If you run the following command on a nonexistent directory, what will happen?
\`\`\`bash
cd nonexistent-directory-name
\`\`\`

1. [x] The command will error out and do nothing
1. [ ] The directory will be created
1. [ ] The directory will be created AND you will navigate into it

# True or False: This command will yield the same result regardless of what your current working directory is.
\`\`\`bash
cd ~
\`\`\`

1. [x] True
1. [ ] False
  > **~** is the shorthand name for the current users home directory, and you can navigate to it from anywhere.

# Assume that `echo-test.txt` is a new file. What will the final output be?
\`\`\`bash
echo 'Hello' > echo-test.txt
echo 'World' > echo-test.txt
echo 'Hello' >> echo-test.txt
cat echo-test.txt
\`\`\`

1. [x] World<br>Hello
1. [ ] Hello<br>World
1. [ ] Hello
1. [ ] World
1. [ ] cat: nonexistent-file-name: No such file or directory
`;

export { rawQuizdown }