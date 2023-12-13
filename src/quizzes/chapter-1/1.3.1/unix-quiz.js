const rawQuizdown = `
---
shuffleAnswers: true
---

# If performed on a nonexistent file, all of the following commands will fail except:

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

# If the following commands are run on a nonexistent file, what will the final output be?
\`\`\`bash
echo 'Hello' > nonexistent-file-name
echo 'World' > nonexistent-file-name
echo 'Hello' >> nonexistent-file-name
cat nonexistent-file-name
\`\`\`

1. [x] World<br>Hello
1. [ ] Hello<br>World
1. [ ] Hello
1. [ ] World
1. [ ] cat: nonexistent-file-name: No such file or directory
`;

export { rawQuizdown }