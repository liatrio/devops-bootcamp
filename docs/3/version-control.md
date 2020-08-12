# Version Control

As the number of people working on a project increases, it becomes more difficult to keep merge conflicts under control and release code often. Using a good version control system and strategy will allow teams to combat problems that come with working on a shared codebase. Different branching strategies can be utilized to help teams keep merge conflicts small and allow for frequent merges into the codebase. Version control also allows teams to keep backups of their entire codebase that can be rolled back if the code is broken

<center>

  ![](img3/git.svg ':size=125px')

</center>

## Branching Strategies

A good branching strategy should solve the problems of merge conflicts and growing release cycles while also being simple to understand and adopt. A great way to measure the effectiveness of a branching strategy is to look at the commit history of a project that uses it.

## Stale Branches

One of the traps that many branching strategies fall into is that they try to maintain several long-living branches. As branches age, they become more and more difficult to merge; it's important to merge branches into a single eternal branch often. Ideally, a team should be merging into master at least once per day.

## Continuous Integration and Fast Feedback

This concept of short-lived branches syncs with how we approach continuous integration and fast feedback. When all the code is checked in to a single branch, everyone's code gets built together in CI and everyone is aware of integration bugs that would otherwise be hiding.

Read this article titled [GitFlow considered harmful](http://endoflineblog.com/gitflow-considered-harmful), and also [the Follow-up](http://endoflineblog.com/follow-up-to-gitflow-considered-harmful).

## Security Considerations

Do NOT commit any private data to ANY repository. Even private repositories may
have public contributors, or may be made public at any time. Private keys,
secrets, passwords, etc. should all be stored securely and/or encrypted.

Git history will include a record of your change, and any sensitive information
will require rewriting of history to fully remove. Overwriting history is
necessary for security, but avoiding it in the first place is greatly
preferred.




