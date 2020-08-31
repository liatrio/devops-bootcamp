# Automated Testing
Automated Testing is a crucial step to ensure code quality because developed tests are guaranteed to run once new source code is committed. Jenkins is a tool that can help us leverage our tests in a build step. With the Testing stage in Jenkins, code quality can be ensured by automatically running the tests that have been written. If tests fail, the build can be marked as unstable and deployment will not occur.

## Exercise
1. Locate the Jenkins Dockerfile created in the container section of the bootcamp.
2. Edit this Dockerfile to add a Go installation to the image.
3. Create a Freestyle Project that will point to your completed Go Unit Testing exercise from your fork of the Bootcamp.
4. Add a build step for testing your exercise.
5. Add a build step to build the exercise.
6. Build the Jenkins job and confirm the job has run the tests and the build was successful.
	
## Multi-Stage Dockerfiles
When having a containerized application, it is a good idea to have a multi-stage Dockerfile to reduce the size and space taken up by the resulting image. In a multi-stage Dockerfile you can implement testing into its own stage. With this strategy, all of the dependencies needed for testing are left in that stage and the runtime stage will only be populated with the necessary build artifacts.

# Deliverable

Discuss the positives of having your testing built into an automation tool such as Jenkins

Discuss the benefits of having a testing stage in a Dockerfile
