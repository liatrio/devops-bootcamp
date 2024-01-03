const rawQuizdown= `

---
shuffleQuestions: true
---

# In case you're worried about security risks stemming from dependency versioning, which of the security tools mentioned above would you use if you wanted to ensure your dependencies get safely updated?

1. [x] Dependabot
1. [ ] Code scanning
1. [ ] Secret Scanning

# SOS!!! Your SQL database just got dropped by maliscious data injection! Which of the tools mentioned above can you apply to help prevent this in the future?

1. [ ] Dependabot
1. [x] Code scanning
	> Good. Code Scanning can be added to your pipeline to track down security risks such as SQL injection vulnerabilities. So long, Johnny Drop Tables :)
1. [ ] Secret Scanning

# Most of us, at one point or another, have accidentally pushed a token to a GitHub repo. It's a simple mistake, but can easily get you into a serious mess. Which tool can you use as the ultimate lifeline to help prevent this?

1. [ ] Dependabot
1. [ ] Code scanning
1. [x] Secret Scanning
	> You'll thank yourself later if you turn this on :)
`;

export { rawQuizdown }