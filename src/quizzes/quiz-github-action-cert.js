const rawQuizdown = `
# Which statement is correct regarding passing permissions to reusable workflows?

> https://docs.github.com/en/actions/using-workflows/reusing-workflows#access-and-permissions

1. [x] The \`GITHUB_TOKEN\` permissions passed from the caller workflow can be only downgraded by the called workflow.
1. [ ] The \`GITHUB_TOKEN\` permissions passed from the caller workflow can be only elevated by the called workflow.
1. [ ] The \`GITHUB_TOKEN\` permissions passed from the caller workflow can be both downgraded and elevated by the called workflow.
1. [ ] The \`GITHUB_TOKEN\` permissions passed from the caller workflow can be neither downgraded or elevated by the called workflow.

# What are the different permission levels You can assign to \`GITHUB_TOKEN\` in the \`permissions\` block?

> https://docs.github.com/en/actions/using-jobs/assigning-permissions-to-jobs

1. [x] none, write, read
1. [ ] read, write, delete
1. [ ] read, write

`;
export { rawQuizdown }