# Maven Releases

Let's bring this all together and think about the process of releasing software.

Imagine PetClinic, Joda-Time, and Junit being maintained by 3 different teams working towards the same production release date.

One thing that is very important in software delivery is to know the exact version of all artifacts being released. Archive the versions with Maven.

Everything that will be released (including dependencies) must be release versions, not SNAPSHOT's.

Here are two ways of doing this.
 - Using the [maven release plugin](http://maven.apache.org/maven-release/maven-release-plugin/).
 - Manually updating `pom.xml` files and using Jenkins automatically building every code push to release.

## Reading

 - [Maven Release Plugin: The Final Nail in the Coffin](https://axelfontaine.com/blog/final-nail.html) by Axel Fontaine
 - [A New Way to Do Continuous Delivery with Maven and Jenkins Pipeline](https://www.cloudbees.com/blog/new-way-do-continuous-delivery-maven-and-jenkins-pipeline) by cloudbees.com

# Maven Integration

This section is designed to give you an introduction into how Maven interacts with Jenkins and Artifactory. Jenkins offers a plugin for Maven that helps seamlessly add the integration between the two tools.

 1. Fork spring-petclinic, joda-time hibernate, and junit.
 2. Use your local Jenkins + Artifactory virtual machine and create jobs that build all three projects.
 3. Deploy the artifacts to Artifactory.
   - Double check the location you are deploying to.
   - It may be a good idea to do `clean install` before `clean deploy`.
 4. Set up all three jobs to build on commit with Git polls.
   - First, configure Git polling to build the jenkins jobs.
   - Next, configure using a webhook and Ngrok.
 5. Mess with things, and see what happens.
   - Make some changes to code.
   - Change some of the versions.
 6. Change one of the versions to a release version.
 7. Validate the deployments to the Artifactory Releases repository worked.
 8. Attempt to deploy that same release version again.
 9. Configure Jenkins so that when dependency artifacts build successfully, the dependent projects build as well.

## Confirming Version in Artifactory
 1. Download the latest version of spring-petclinic from your Artifactory repository.
 2. Open up the WAR file.
 3. Find the joda-time hibernate library inside the WAR file.
 4. What version are you getting? Check with your lead to ensure this is the correct version.
 5. Clone joda-time-hibernate and verify the [buildnumber-maven-plugin](http://www.mojohaus.org/buildnumber-maven-plugin/usage.html) is included so two values are added to the Maven metadata file.
   - SCM Revision Number and Build Timestamp  

<center>

  ![](img4/river.svg ':size=125px')

</center>

## Tips

 - Make sure to be creating Maven jobs in Jenkins.
 - Is Jenkins using the correct `settings.xml`? Are you sure?
 - What does the ID in the Distribution Management section align with?
 - Have you created a release version twice?
 - Don't rush. Think about what is going on behind the scenes.



# Deliverable

Discuss what a release is and how they can be performed. 

Discuss how rebuilding release versions twice in Jenkins can cause downstream issues and the benefits of plugins such as buildnumber.

Discuss with your group any difficulties you ran into, what worked well, and show off your completed deployment setup. 

