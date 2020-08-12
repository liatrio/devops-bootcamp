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

