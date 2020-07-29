# Releases

Let's bring this all together and think about the process of releasing software.

Imagine PetClinic, Joda-Time, and Junit being maintained by 3 different teams working towards the same production release date.

One thing that is very important in software delivery is to know the exact version of all artifacts being released. Archive the versions with Maven.

Everything that will be released (including dependencies) must be release versions, not SNAPSHOT's.

Here are two ways of doing this.
 - Using the [maven release plugin](http://maven.apache.org/maven-release/maven-release-plugin/).
 - Manually updating `pom.xml` files and using Jenkins automatically building every code push to release.

<center>

  ![](img4/versions.svg ':size=125px')

</center>

## Reading

 - [Maven Release Plugin: The Final Nail in the Coffin](https://axelfontaine.com/blog/final-nail.html) by Axel Fontaine
 - [A New Way to Do Continuous Delivery with Maven and Jenkins Pipeline](https://www.cloudbees.com/blog/new-way-do-continuous-delivery-maven-and-jenkins-pipeline) by cloudbees.com

# Deliverable

Discuss what a release is and how they can be performed.
