## Go Releaser

[Go Releaser](https://goreleaser.com/install/) is a tool for releasing and publishing Go Projects. It is meant to be used with CI Tools as a way to simplify the process of releasing the built Go Binaries into a GitHub/GitLab release. Go Releaser leverages a `.goreleaser.yml` file to customize what will be built and archived. 

## Exercise
1. Find an Open Source Go Project to fork and clone to your local system. 

?> This is a good list of [Trending Go Projects](https://github.com/trending/go)

2. Install Go Releaser onto your system.

3. Use Go Releaser locally to deploy the Go binaries to a GitHub release using the GitHub token with repo scope.

4. Now using your Jenkins instance from previous sections, write a Jenkinsfile that will use Go Releaser to deploy the Go binaries to the GitHub release of your chosen tag.

?> You may run into problems when making a release with Jenkins after creating one locally because Go Releaser creates folders that will give a Git dirty state. In the Jenkinsfile, run commands to delete those folders ahead of time.
