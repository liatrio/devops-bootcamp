# Merge Conflicts

1. Clone your new repository to your local machine.
2. Add a file to the repository, and push your changes.
3. Clone the repository again in a different location.
4. In both locations edit the file in different ways and commit your changes.
5. Push your changes from the first repository, and then pull in the second
repository.

We'll now created a problem, git doesn't know which version of the file is the
correct version to keep, generating a merge conflict. When viewing the file,
you'll see both versions of the conflicting section surrounded by `<<<` and
`>>>`. The two versions are separated in the center by `===`.

Each conflict is resolved in a different way, sometimes one version of the code
may be more correct, and you should keep that version. In some cases you may
want to keep parts of both code changes. Lastly you might want to both sets of
changes. In any case, the `<<<`, `>>>`, and `===` are simply markers to denote
the end of each section of changes. Generally when resolving a merge conflict
those lines should be deleted.

When you're done resolving the conflict in which ever way you prefer, commit
the updated file, and push them to GitHub.

# Pull Requests

1. In your repository make a new branch and switch to it with
`git checkout -b my_new_branch`.
2. Make some changes to files on this branch.
3. Push your changes to the remote branch with `git push origin my_new_branch`.

Now you're ready to submit a pull request. A pull request, or PR, is a request
for others to pull in your work into their own repository. Pull requests aren't
a built in feature of git, rather they are a functionality of GitHub.

If you navigate to your repository on GitHub, and your pull request tab, you
can create a new pull request. Choose to compare your new branch with a base
of master. Once you submit the pull request you'll have the option to request
reviews from your peers and have an area for discussions about the branch.

Once the pull request is acceptable, it can be merged into master. In some
cases different options can be made to restrict merging, such as requiring
peer review, or unit testing.

# Commit Reversals

1. Add another file to your git repository, commit, and push it.
2. Add some "sensitive content" to the file, such as a line "SENSITIVE INFO".
3. Commit and push your file.
4. Remove the sensitive data, and commit and push the changes.

We have a problem here. Even though we removed the sensitive data from the
file, git's history still has a record of those changes. We'll have to
overwrite the git history, while also working to preserve the commits that
didn't include any private data.

## Option 1: git reset
Using a git reset, we can get rid of any commits after a specific point. As long as we don't use reset --hard the files themselves will not be changed, only the git history will be impacted. This is a good option if the trouble commit is near or at the current head of our git repository. To undo a commit with reset:

1. Use `git log` to find the hash of the commit just before the credentials
were added.
2. `git reset 1a2b3c4` where 1a2b3c4 are the first 7 characters of the hash
you just found.
3. Edit any files necessary to remove the credentials.
4. Commit the sanitized files.
5. Push the changes using the `-f` flag.

This is a simple and fairly easy approach, however it has the side effect of
combining all commits after the error into one. The actual changes will be
preserved, but each individual commit message will be lost.

## Option 2: git rebase
Using a git rebase, we can remove only a single commit, which is useful for a commit that is not near the head, or for keeping important commit messages. To rebase:
1. Get the hash value of the commit containing credentials.
2. Start an interactive rebase on that commit `git rebase -i 1a2b3c4^` where
1a2b3c4 are the first 7 characters of your hash and the last character is a '^'.
3. Change 'pick' to 'edit' for the commit you want to change, and then exit the file
4. The rebase will stop when the 'edit' point is reached. Sanitize the commit
by removing the credentials wherever necessary.
5. Add the change and run 'git commit --amend' (Do not change anything in the amend).
6. Continue the rebase: `git rebase --continue`.
7. Push the changes to remote using the `-f` flag.

This is a more involved approach, but allows all commits and history to be preserved.

These approaches both change git history, which will require a force push to propagate to the remote repository. Make sure you can force push, or ask an administrator to do it for you if necessary. Rewriting git history is a destructive process, and should only be carried out when necessary for security.

# Deliverable

Discuss what merge conflicts are and strategies to avoid them.

Discuss the benefits and downsides of using pull requests to merge code within
a project.

Discuss the problems that can be caused by overwriting git history, and
ways to mitigate the problems caused by it.
