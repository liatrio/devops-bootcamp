# Managing Infrastructure

In this section you will be taking what you've learned about virtual machines to build on top of your Golden Image and deploy servers to create a simple software delivery environment. We will cover the tools used in this section in greater detail later in the Bootcamp but for now here are some basic definitions to help you understand the environment will be creating. 

## Jenkins 
  Jenkins is a build automation server that manages and controls software delivery processes throughout the entire lifecycle. It orchestrates tasks including build, test, package, staging, deployment, static code analysis.

<center>

  ![](img2/jenkins.svg ':size=125px')

</center>

  - A leading open source continuous integration tool
  - Builds, tests, and deploys software continuously and monitors the execution and status of jobs
  - Trigger jobs manually, on a schedule, or on a code commit
  - Uses Java and Tomcat to do its work

## Artifactory
  Artifactory is a versioned artifact repository. It is used to store multiple versions of different packages and supports different repository types (Maven, npm, docker, etc). It is scalable and integrates with lots of other software delivery tools.

## Delivery Pipeline

<center>

  ![](img2/delivery-pipeline-stages.svg)

</center>

  Automated delivery pipelines accelerate software delivery and enable faster, safer deployments.
  An entire integrated toolchain is necessary to leverage the capabilities of an automated pipeline, orchestrated by a build server. Newly committed code that is pushed to the git server will send a Webhooks to the build servers to trigger tasks. General stages of a delivery pipeline include:
- Unit tests
- Code quality scanning
- Build artifacts (binaries, dist, containers)
- Publishing artifacts to an artifact repository
- Functional testing
- Deploying artifacts

# Exercise

Configure a Jenkins and Artifactory server and setup a Jenkins job to deploy a build artifact to Artifactory.

 1. Create a clone of your Golden Image from the previous section and [install Jenkins](https://wiki.jenkins.io/display/JENKINS/Installing+Jenkins+on+Red+Hat+distributions). 
 2. In GitHub create a Fork of the [spring-petclinic](https://github.com/spring-projects/spring-petclinic) project.
 3. Create a Maven job in Jenkins to build your Fork of spring-petclinic.
   ?> You may need to install some dependencies for the project to build successfully.
 4. Create another clone of your Golden Image and [install Artifactory OSS](https://www.jfrog.com/confluence/display/JFROG/Installing+Artifactory). Create a Maven repository during the Artifactory setup.
 5. Install the Artifactory Jenkins plugin.
 6. Add a Post-Build Action to deploy the build artifacts to your Artifactory server.
 7. Build the spring-petclinic job and confirm the artifacts are deployed to Artifactory.

### Extra Credit 
Working as a group configure your Jenkins server to deploy to someone elses Artifactory server using what you have learned about virtual networks.

?> If you are working on different networks you man need to use a tool like [Ngrok](https://ngrok.com/) to bridge communication between your servers over the Internet. 

<center>

  ![](img2/network.svg ':size=125px')


</center>

# Deliverables

|**Engineers**|**Consultants**|
|-------------|---------------|
| Experience using virtual machines to deploy software delivery tools | Research what Virtual Machines are and how they would be used at an Enterprise |
| Understand basics of a software delivery pipeline | Create a short slide deck explaining your findings |
| **Discuss:** After complete this exercise how would you improve the process to provision the Jenkins and Artifactory VMs? | Don't forget to use the Liatrio Presentations template |
| **Discuss:** How could this software delivery pipeline be improved? |  |
