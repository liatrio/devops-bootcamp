# Version Control

<center>

  ![](img3/git.svg ':size=125px')

</center>

A Version Control System (also referred to as Source Code Management) is the foundation for Continuous Integration. It allows teams to iterate quicker and release more frequently by being able to work on code changes in isolation and then merge in their changes in a way that does not interfere with other work. As teams make smaller changes more frequently the ability to track changes, ensure they meet quality standards and do not conflict becomes critical. However just using a VCS does not guarantee a team will perform better. It is important that teams also develop a source code management strategy that makes integrating changes to their codebase easier, faster and less error prone.

## Security Considerations

Do NOT commit any private data to ANY repository. Even private repositories may
have public contributors, or may be made public at any time. Private keys,
secrets, passwords, etc. should all be stored securely and/or encrypted.

Git history will include a record of your change, and any sensitive information
will require rewriting of history to fully remove. Overwriting history is
necessary for security, but avoiding it in the first place is greatly
preferred.

