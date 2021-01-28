# Docker Compose

[Docker Compose](https://docs.docker.com/compose/) manages multiple Docker containers using a single `docker-compose.yml` configuration file. With it you can build and run an entire environment using just the `docker-compose up` command. Sound familiar to any other tools you've used lately? The configuration options correlate directly with arguments used with the Docker CLI which makes it an easy tool to start using as you're getting familiar with Docker. It lacks features needed for complicated environments but has some very useful use cases.

## Use Cases

- Managing local development environments or sharing among a team of developers.

- Sharing a demo of a containerized application.

Read the Docker Compose [getting started](https://docs.docker.com/compose/gettingstarted/) guide to familiarize yourself with writing a Docker Compose file and the basic commands.

## Exercise

This exercise will build on the earlier container section by creating a Docker Compose file to run your Jenkins and Artifactory containers.

1. Create a `docker-compose.yml` file with the following criteria:
    
    - Builds and runs your Jenkins and Artifactory images.
    
    - Opens ports to expose the UI.
    
    - Creates volumes to persist application data on your local filesystem.

?>Docker Compose [networking](https://docs.docker.com/compose/networking/) helps running services communicate by allowing each container to query other containers using the service name as the hostname.

2. Run `docker-compose up` to create your environment and ensure you can access the application UIs.

3. Configure a Jenkins job to build Spring Pet Clinic and deploy the artifact to Artifactory.

4. Run `docker-compose down` and `docker-compose up` to recreate the environment and check that the Jenkins configuration and Artifactory data are persisted.

