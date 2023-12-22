const rawQuizdown = `
---
shuffleAnswers: true
---

# Given the above .ssh/config file, select all commands that will successfully ssh to the host. You can assume the public key file has already been sent to the host.

- [x] ssh my_host
- [x] ssh \`ec2-user@ec2-42-42-42-42.us-west-2.compute.amazonaws.com\` -i ~/.ssh/host_public_key.pub
- [ ] ssh \`ec2-user@ec2-42-42-42-42.us-west-2.compute.amazonaws.com\`
	> In this ssh command, no private key is specified, so you'll have no way of authenticating
- [ ] ssh ec2-user@my_host

# What are the main reasons for ssh key exchanges? Select all that apply:

- [x] They provide a way for automated actions to login successfully without having to enter a password
- [x] They eliminate the TOIL that comes with entering passwords repeatedly
- [x] They eliminate the need for a password, while also being more secure
- [ ] They eliminate the need for a password, but are less secure than using a password
	> Using key exhanges are *more* secure than passwords, because they are longer, more complex, and encrypted
- [ ] They allow multiple users to share the same authentication credentials, which enhances security
	> This is actually harmful for security. If one user's information leaks, it compromises others aswell

# You read above that the \`.ssh\` directory should have permissions set to 0700 and the \`.ssh/config\` file should have permissions set to 0600. Why should the one's and ten's place be set to 0?<br>
You may choose to explore more about permissions <a href="https://mason.gmu.edu/~montecin/UNIXpermiss.htm">here</a>

1. [x] The hundred's place resembles owner permissions, and groups and others shouldn't have access to these, so they are set to 0
1. [ ] The hundred's place resembles group permissions, and owners and others shouldn't have access to these, so they are set to 0
1. [ ] The hundred's place resembles other's (not owner or group) permissions, and owners and groups shouldn't have access to these, so they are set to 0

# What does the 0700 permission denote for the \`.ssh\` directory?<br>
You may choose to explore more about permissions <a href="https://mason.gmu.edu/~montecin/UNIXpermiss.htm">here</a>

1. [x] Read, Write, and Execute permissions
1. [ ] Read and Write permissions
1. [ ] Read and Execute permissions
1. [ ] Write and Execute permissions
1. [ ] Read permissions
1. [ ] Write permissions
1. [ ] Execute permissions

# What does the 0600 permission denote for the \`.ssh/config\` file?
1. [ ] Read, Write, and Execute permissions
1. [x] Read and Write permissions
1. [ ] Read and Execute permissions
1. [ ] Write and Execute permissions
1. [ ] Read permissions
1. [ ] Write permissions
1. [ ] Execute permissions
`;

export { rawQuizdown }