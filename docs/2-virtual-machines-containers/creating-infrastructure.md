## Jenkins 
  Jenkins is a build automation server that manages and controls software delivery processes throughout the entire lifecycle. It orchestrates tasks including build, test, package, staging, deployment, static code analysis.

<center>

  ![](img2/jenkins.png ':size=125px')

</center>

  - A leading open source continuous integration tool
  - Builds, tests, and deploys software continuously and monitors the execution and status of jobs
  - Trigger jobs manually, on a schedule, or on a code commit
  - Uses Java and Tomcat to do its work


## Delivery Pipeline
  ![](img2/delivery-pipeline-stages-1.png)

  Automated delivery pipelines accelerate software delivery and enable faster, safer deployments.
  An entire integrated toolchain is necessary to leverage the capabilities of an automated pipeline, orchestrated by a build server. Newly committed code that is pushed to the git server will send a Webhooks to the build servers to trigger tasks. General stages of a delivery pipeline include:
- Unit tests
- Code quality scanning
- Build artifacts (binaries, dist, containers)
- Publishing artifacts to an artifact repository
- Functional testing
- Deploying artifacts

# Exercise
## Local Deployment

 1. Using a clone of your Golden Image from the previous section, [install Jenkins](https://wiki.jenkins.io/display/JENKINS/Installing+Jenkins+on+Red+Hat+distributions).
 2. Using the same clone, install Nexus.
 3. Deploy [spring-petclinic](https://github.com/spring-projects/spring-petclinic) from Jenkins to Nexus.
<center>

  ![](img2/deploy.svg ':size=125px')
</center>

## Discussion
Discuss as a group the following topics.
 - What methods did you try?
 - Did you run into any issues?


## Remote Deployment

This is a group exercise. If you are just one person, that's okay; we believe in you! However, try to complete this as a group if possible.

 1. One person should clone their Golden Image again, and they should install Jenkins on it.
 2. Another person should clone their Golden Image again, and they should install Nexus on it.
 3. Each person should forward the correct ports on their Virtual Machine.
 4. Over a network, deploy spring-petclinic from Jenkins to Nexus.

<center>

  ![](img2/network.svg ':size=125px')


# Deliverables
</center>

<div class="grid2"><div class="col">
<center>

**Engineers**

</center>

- Successfully deploy spring-petclinic from Jenkins to Nexus over a network.
- What methods did you try?
- Did you run into any issues?
</div><div class="col">
<center>

**Consultants**

</center>

- Research what Virtual Machines are and how they would be used at an Enterprise
- Create a short slide deck explaining your findings
	- Don't forget to use the Liatrio Presentations template

</div></div>