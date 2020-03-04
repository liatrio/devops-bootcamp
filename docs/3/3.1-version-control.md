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

## Git

Git is one of the most popular version control softwares available, and the one used at
Liatrio. Git can be used with only local repositories, but it can also be
used alongside remote git repositories, such as ones hosted by GitHub,
Bitbucket, and GitLab. Git is fully distributed, and every repository,
including local repositories contain full history and tracking abilities.

### Common Git Commands

| Command        | Description                                                                                                    |
|----------------|----------------------------------------------------------------------------------------------------------------|
|   `git init`   | Initializes a new local git repository                                                                         |
|   `git clone`  | Creates a local git repository copied from a remote                                                            |
|    `git add`   | Stages new local changes to be committed                                                                       |
|  `git commit`  | Commits staged changes into git history                                                                        |
|   `git push`   | Pushes local commits to a remote repository                                                                    |
|   `git pull`   | Pulls new commits from a remote repository and merges them with a local copy                                   |
|  `git branch`  | Provides information about different development branches                                                      |
| `git checkout` | Changes the state of HEAD to a different branch or commit                                                      |
|  `git status`  | Displays current information about the state of your git repository                                            |
|    `git log`   | Displays git's commit history                                                                                  |
|   `git reset`  | Resets git history to the state of a previous commit, permanently changing history                             |
|  `git rebase`  | Appends a git branch onto another branch or commit, permanently changing history |

## Subversion (SVN)
Although git is one of the most widely used version control solutions, there are others that are worth being aware of. Subversion (SVN) is a version control tool that works similarly to git, but has some key differences that make it a good option in certain scenarios. With git, the user keeps a local copy of the entire project repository and makes commits locally before pushing changes to a remote repository. SVN allows users to check out individual files that need to be worked on and then commit the changes directly back to the server. This makes SVN a good option when projects get too large to store on a local machine. SVN is also considered easier to use as git can become pretty complex with its wide variety of commands. SVN also has different branching and security systems that are worth considering when choosing a version control solution. Read [Git vs. SVN](https://www.perforce.com/blog/vcs/git-vs-svn-what-difference) to highlight the major differences between git and SVN.

While the Bootcamp tends to focus on git as a version control solution, it is important to be aware of the pros and cons of other version control tools such as SVN.

### Common SVN Commands
| Command           | Description                                                                              |
|-------------------|------------------------------------------------------------------------------------------|
| `svnadmin create` | Creates a new SVN repo                                                                   |
| `svn import`      | Adds local files to the remote repo                                                      |
| `svn checkout`    | Checks out (Copies) a file locally from the remote repo                                  |
| `svn commit`      | Send changes from local files to the remote repo                                         |
| `svn status`      | Prints the status of local files and directories as compared to those on the remote repo |
| `svn revert`      | Revert local copies of files to reflect the versions in the remote repo                  |
| `svn merge`       | Apply the differences between two sources                                                |
| `svn update`      | Bring changes from the remote repo into the local filesystem                             |



# Deliverable

Locate an open-source project on GitHub and determine the branching strategy being used. Examine how difficult it would be to identify a bug introduced in that project's commit history. Discuss as a group.
