# Maven Integration

This section is designed to give you an introduction into how Maven interacts with Jenkins and Nexus. Jenkins offers a plugin for Maven that helps seamlessly add the integration between the two tools.

 1. Fork spring-petclinic, joda-time hibernate, and junit.
 2. Use your local Jenkins + Nexus virtual machine and create jobs that build all three projects.
 3. Deploy the artifacts to Nexus.
   - Double check the location you are deploying to.
   - It may be a good idea to do `clean install` before `clean deploy`.
 4. Set up all three jobs to build on commit with Git polls.
   - First, configure Git polling to build the jenkins jobs.
   - Next, configure using a webhook and Ngrok.
 5. Mess with things, and see what happens.
   - Make some changes to code.
   - Change some of the versions.
 6. Change one of the versions to a release version.
 7. Validate the deployments to the Nexus Releases repository worked.
 8. Attempt to deploy that same release version again.
 9. Configure Jenkins so that when dependency artifacts build successfully, the dependent projects build as well.

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

Discuss with your group any difficulties you ran into, what worked well, and show off your completed deployment setup.
