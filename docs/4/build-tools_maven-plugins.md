# Maven Plugins

Plugins are a great way of enhancing tool function without much effort. The build number plugin does a lot and helps to add metadata to projects for logistics purposes.

?> Use the set up from the previous section to do this section.

 1. Download the latest version of spring-petclinic from your Nexus repository.
 2. Open up the WAR file.
 3. Find the joda-time hibernate library inside the WAR file.
 4. What version are you getting? Check with your lead to ensure this is the correct version.
 5. Clone joda-time-hibernate and add the [buildnumber-maven-plugin](http://www.mojohaus.org/buildnumber-maven-plugin/usage.html) so two values are added to the Maven metadata file.
   - SCM Revision Number and Build Timestamp
 6. Get the Joda-time-hibernate artifact and ensure that both values have been added correctly to the project.

<center>

  ![](img4/plugin.svg ':size=125px')

</center>

## Repositories

Read [Repository Manager Best Practices](https://help.sonatype.com/display/NXRM2/Repository+Manager+Best+Practices) from the Sonatype documentation.

It's important to see how components work together in real environments.
 - As a developer of spring-petclinic, you may not necessarily have joda-time and junit building on your computer, but you will have a dependency on them.
 - Imagine your virtual machine is actually a server within your company and not on your local machine.

Configure Maven on your laptop to use Nexus as the repository that grabs dependencies.
 1. Setup your Nexus as an additional repository to (the default) Maven Central.
 2. Setup your Nexus as the _only_ repository the computer will use and have Nexus _proxy_ everything.
 3. Try changing the GroupID of your clones to mimic proprietary software. You can use `com.devops-bootcamp.samples`.


# Deliverable

Discuss how rebuilding release versions twice in Jenkins can cause downstream issues, the benefits of plugins such as buildnumber, when a proxy should be used for dependencies, and the significance of the GroupID.
