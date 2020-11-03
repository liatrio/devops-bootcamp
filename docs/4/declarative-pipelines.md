# Declarative Pipelines

A Declarative Pipeline is a method for automating the process of setting up a pipeline in jenkins. To do this we will be using a Jenkinsfile. Jenkinsfiles contain the definition of a Jenkins Pipeline in steps and is checked into source control to be used by jenkins. 

### Example Jenkinsfile
```
pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
```

``` agent any ```
- Execute this Pipeline or any of its stages, on any available agent

``` stages ```
- Lists the stages that will be executed in this pipeline, for this example there are 3:
	- Build
	- Test
	- Deploy

``` stage('Build') ```
- Defines a stage 'Build'.

``` steps ```
- Performs some steps related to the 'Build' stage.

``` stage('Test') ```
- Defines a stage 'Test'.

``` steps ```
- Performs some steps related to the 'Test' stage.

``` stage('Deploy') ```
- Defines a stage 'Deploy'.

``` steps ```
- Performs some steps related to the 'Deploy' stage.

> Note: Stage name can be anything, and there can be any number of stages. Link to Jenkins Documentation on Pipelines: https://jenkins.io/doc/book/pipeline/

## Exercise

1. Using your Jenkins and Artifactory instances from the previous section, implement a Jenkinsfile to automate your pipeline. 
    - For the three projects spring-petclinic, joda-time hibernate, and junit your projects should build, test, and deploy each to Artifactory.
    - When deploying to Artifactory, do not use plugins and instead use shell commands and the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API).
2. Setup each Pipeline to build and test every branch, but only deploy from the master branch. 


# Deliverable

Describe why a declarative pipeline is used with Jenkins and Artifactory.
