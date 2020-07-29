# Jenkins 
  Jenkins is a build automation server that manages and controls software delivery processes throughout the entire lifecycle. It orchestrates tasks including build, test, package, staging, deployment, static code analysis.

<center>

  ![](img2/jenkins.png ':size=125px')

</center>

  - A leading open source continuous integration tool
  - Builds, tests, and deploys software continuously and monitors the execution and status of jobs
  - Trigger jobs manually, on a schedule, or on a code commit
  - Uses Java and Tomcat to do its work


Jenkins is at the heart of the pipeline, providing capabilities to achieve Continuous Integration and Continuous Delivery, Continuous Feedback.

### Pipelines
Jenkins is at the heart of the pipeline, providing capabilities to achieve continuous integration, continuous delivery, and continuous Feedback. The definition of a [Jenkins pipeline](https://www.jenkins.io/doc/book/pipeline/jenkinsfile/) is written into a text file Jenkinsfile in the project's source control repository, defining each step in the pipeline in various stages. This script is executed automatically from new commits to source code through webhooks.

```Jenkinsfile
pipeline {
    stages {
        stage('Build') {
            steps {
                // commands to build app
            }
        }
        stage('Deploy') {
            steps {
                // commands to deploy app
            }
        }
    }
}
```
### Shared Libraries
As pipelines are adopted for more and more projects in an organization, common patterns are likely to emerge. Oftentimes it is useful to share parts of pipeline code between various projects to reduce redundancies and keep code simple. For this reason, code can be leveraged from shared libraries which can be defined in external source control repositories and loaded into existing pipelines.

``` library 'shared-library' ```

### Plugins
Plugins are the primary means of enhancing the functionality of a Jenkins environment to suit user-specific needs. There are over a 1500 different plugins which can be installed on a Jenkins master and to integrate various build tools, cloud providers, analysis tools, and more.

### Agents: Distributed Architecture
A Jenkins master can operate by itself both managing the build environment and executing the builds with its own executors and resources. However, organiations will quickly realize the resource limits and need to dertime a [scaling strategy](https://www.jenkins.io/doc/book/architecting-for-scale/). Jenkins supports the "master/agent" mode, where the workload of building projects are delegated to multiple agents.

*[Jenkins documentation](https://www.jenkins.io/doc/)*