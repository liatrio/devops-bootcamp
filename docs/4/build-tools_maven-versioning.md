# Versioning



## Exercise

 1. Create a new application using the archetype.
```
mvn archetype:generate -DarchetypeArtifactId=maven-archetype-webapp
```
 2. Add Jetty to `pom.xml`.
```
<build>
        <plugins>
            <plugin>
                <groupId>org.mortbay.jetty</groupId>
                <artifactId>maven-jetty-plugin</artifactId>
                <version>6.1.10</version>
                <configuration>
                    <scanIntervalSeconds>10</scanIntervalSeconds>
                    <connectors>
                        <connector implementation="org.mortbay.jetty.nio.SelectChannelConnector">
                            <port>8080</port>
                            <maxIdleTime>60000</maxIdleTime>
                        </connector>
                    </connectors>
                </configuration>
            </plugin>
        </plugins>
</build>
```
 3. Run Jetty.
```
mvn jetty:run
```
 4. Investigate the Maven Internal (local) repo. What does it contain? What is its purpose?
 5. What changes occur in the project directory and the local repo when running `mvn package` and when running `mvn install`?
 6. Modify dependencies/versions using the POM by opening `pom.xml` and changing the versions of dependencies to get different versions to be used in the project. Locate them in the Maven repo.
 7. Change spring-petclinic's group ID, artifact ID and version.
 8. Alter one of spring-petclinic's open-source dependencies, such as jodatime-hibernate.
   - Change the jodatime dependency in spring-petclinic to use `9.2.1-SNAPSHOT`.
   - After building, the build should fail. Why?
 9. Clone jodatime and create a new version of the plugin.
   - Change the version to be the missing version from before.
   - Build the "new" version which is stored in the local Maven repo.
 10. Run spring-petclinic build and verify it was successful.






# Deliverable

Discuss with your group example scenarios and what type of new version they would need and when a new major version is necessary.
