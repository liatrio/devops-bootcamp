# Containers

>&ldquo;Containers and virtual machines have similar resource isolation and allocation benefits, but function differently because containers virtualize the operating system instead of hardware. Containers are more portable and efficient.&rdquo;

![](img2/containers_light.svg ':size=150x150 :class=light-mode-icon')
![](img2/containers_dark.svg ':size=150x150 :class=dark-mode-icon')

Containers allow you to easily package an application's code, configurations, and dependencies into easy to use building blocks. Containers can help ensure that applications deploy quickly, reliably, and consistently regardless of deployment environment. This enables the following benefits:

- **Developer Productivity.** Containers enable engineers to run newly developed applications and services in the same infrastructure that their applications will run on in production. This creates a faster feedback loop for engineers about application changes.

- **Environmental Consistency & Versioning.** Container images are a point in time capture of an application's code, configuration, and dependencies. These images are able to be versioned per build, and each artifact is considered immutable. Deployments are represented as code, tagged, and versioned for auditability and clarity.

- **Operational Efficiency.** Being able to build, test, ship, and run the exact same container idempotently through all stages of the Software Delivery Life Cycle makes delivering a high quality, reliable application considerably easier. Continuous Integration servers can run the same container artifact through multiple stages, including integration tests, static code scanning, load testing, and functional testing, all to make sure the artifact passes known quality gates.

- **Resource Density.** Containers maximize the hardware by allowing multiple processes to run on a single piece of hardware, or a logical cluster of hardware. Efficiency is a result of the isolation and allocation techniques that managed container services use. This allows for better projected utilization costs, as well as capacity planning for existing and future scaling for traffic needs, and experimentation without costly environment spin-up.

## Containers vs Virtual Machines

![A diagram showing progression from traditional server deployments to virtual machines to container-based deployments](img2/delivery-containers-evolution.svg ':class=img-center')

The picture above represents the progression that has happened over the last 15 years in the industry for deploying applications. In the beginning, apps were deployed to their own physical hosts, each with their own operating system.

Then, Virtual Machines (VMs) were introduced; that allowed for slightly higher process density per physical host, and some logical grouping of applications to be done.

In the third progression, containers are introduced into the flow. Containers are a smaller form of a VM that only contains the application code and it's necessary dependencies to run. This progression of technology has led to greater flexibility and scalability.

## Containers in Production
Some people have the mistaken idea that containers is only for development and testing, but there are thousands of companies running in production with containers on Kubernetes and OpenShift. The image is meant to be immutable, meaning it will be identical no matter where it is deployed, whether it is in development or production hardware. Every iteration of a container will deploy the same way on top of the orchestrator.

Read more from [Docker: What is a Container](https://www.docker.com/resources/what-container) to get a better understanding of containers and how they compare to virtual hosts.


# Docker

>&ldquo;Docker is a platform to create, run and manage containers and container resources (images, networks and data volumes). It consists of a server, a REST API and a CLI client to interact with the API.&rdquo;

Docker is often used synonymously with containers but is actually a platform for using containers. Similar to how Java has development and runtime environments created by different organizations, there are many [other platforms for containers](https://jfrog.com/knowledge-base/6-alternatives-to-docker-all-in-one-solutions-and-standalone-container-tools/) however Docker is by far the most widely adopted container platform.

![](img2/cloud_docker_light.svg ':size=100x100 :class=light-mode-icon')
![](img2/cloud_docker_dark.svg ':size=100x100 :class=dark-mode-icon')

## Images and Containers
>&ldquo;An image is an executable package that includes everything needed to run an application--the code, a runtime, libraries, environment variables, and configuration files.&rdquo; - [Docker Concepts](https://docs.docker.com/get-started/#docker-concepts)

>&ldquo;A container is a runtime instance of an image--what the image becomes in memory when executed (that is, an image with state, or a user process).&rdquo; - [Docker Concepts](https://docs.docker.com/get-started/#docker-concepts)

Images are the building blocks of Docker. They define what is in a container. They can be used by other images to extended extend their functionality. Images can also be shared across hosts.

# Running Containers

Install [Docker](https://docs.docker.com/install) and run the [hello world container](https://hub.docker.com/_/hello-world) `docker run hello-world`

This deceptively simple command actually does a lot which is explained by the output of the container.

If the Docker image hello-world does not exist in the local image repository it downloads it from Docker Hub and adds it to the local repository. You can use `docker image ls` to list the locally installed images. You will see the hello-world image in this list. You can download or update an image directly with `docker pull <IMAGE NAME>`.

Once the hello-world image is available it creates and runs a container from the image. Images define the default command that is run when the container is started. When the command exits the container stops. You can list running and stopped containers with `docker container ls -a`.

Note that it has automatically generated a name for the container since we did not specify one. What happens to the images and containers if you run `docker run hello-world` again?

### Cleaning Up

Remove the stopped hello world containers with `docker rm CONTAINER_ID|NAME`.

Remove the hello world image with `docker image rm IMAGE_ID|NAME`.

?>Docker Desktop on OSX limits the amount of resources containers can use by default. If you experience performance issues with resource intensive containers make sure to increase the resources available. `Docker Desktop icon` -> `Preferences` -> `Advanced`

# Building Images

Docker images are built by defining a set of steps to create the image in a [Dockerfile](https://docs.docker.com/engine/reference/builder/). Each step begins with a single word instruction, which are uppercase by convention, followed by arguments all on the same line. The image is then built by running `docker build PATH_TO_DOCKERFILE`.

## Base Images
A base image is the image that is used to create all of your container images. It is referenced using the `FROM` instruction in a Dockerfile. A proper implementation of base images improves maintainability of your images. Your base image can be an [official Docker image](https://docs.docker.com/docker-hub/official_images/), such as Centos, or you can modify an official Docker image to suit your needs, or you can create your own base image from [scratch](https://hub.docker.com/_/scratch). When building images, there are a few things to consider in regards to the base image from which to build from.

### Use official images
Where possible, use [official images](https://hub.docker.com/search?q=&type=image&image_filter=official) from [Docker Hub](https://hub.docker.com/) rather than installing tools manually. For example, the following Dockerfile is installing Java.

```docker
## Don't do this
FROM centos:centos6
RUN yum install java-1.7.0-openjdk-devel -y

## Better
FROM openjdk
```

### Tag specific base image versions 
If you omit the tag on the `FROM` line, you will end up with the latest tag. This can result in inconsistencies in the image that is built as the latest tag can change between times when the image is built. Therefore, be sure to always include a tag to pin to a specific upstream image.

```docker
## Don't do this
FROM openjdk

## Better
FROM openjdk:7
```

### Use specialized images
Container images, especially base images, should be specialized for a specific purpose and not contain any unnecessary packages or tools. Even small tools that you might normally consider essential or innocuous such as text editors, debugging tools or network diagnostic tools should be avoided. Adding unnecessary tools not only increases the size of the image it increases the attack surface of container which could introduce a security vulnerability.

## Layers
A Docker image consists of read-only layers each of which represents a Dockerfile instruction. The layers are stacked and each one is a delta of the changes from the previous layer. Care must be taken when designing a Dockerfile to optimize both the time it takes to build the image as well as the size of the image that is created.

?>Layers reduce the time it takes to build the image through efficient use of image layers.

### Order of instructions
Docker will cache layers from previous builds to decrease the time required to build images. However, once Docker detects that a layer needs to be rebuilt, then all layers after it must also be rebuilt. Consider the following Dockerfile.

```docker
FROM ubuntu
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk
COPY . /app
```

### Dependency layer
Downloading dependencies for an application can be an expensive operation. You can leverage Docker layer caching to avoid this in image build time by having a separate layer that resolves the dependencies for the application. In the following example, we have a layer to `RUN gradle build` before we add the source code to the image. This results in a layer with all our application dependencies downloaded.

```docker
COPY build.gradle gradle.properties settings.gradle
RUN gradle build
```

## Multi-stage builds
One of the most challenging things about building images is keeping the image size down. Each instruction in the Dockerfile adds a layer to the image, and you need to remember to clean up any artifacts you don’t need before moving on to the next layer. With multi-stage builds, you use multiple `FROM` statements in your Dockerfile. Each `FROM` instruction can use a different base, and each of them begins a new stage of the build. You can selectively copy artifacts from one stage to another, leaving behind everything you don’t want in the final image.

### Build from source
In order to ensure repeatability of image builds, you should put as much of the build process as possible inside the Dockerfile. For example, if you run Gradle to build a jar from your workspace and then have the Dockerfile copy in the jar file, you run the risk of other developers not having the same tools on their workspace. In order to mitigate that, install and run Gradle from within the Dockerfile:

```docker
COPY build.gradle gradle.properties settings.gradle
RUN gradle build
COPY . /app
RUN gradle clean build release
```

### Separate build stage
The concern with running the build within your Dockerfile is the additional space consumed by the image for all the build tools. This can be addressed through the use of a separate stage for the build.

```docker
## Builder stage
FROM openjdk:7 as builder
COPY build.gradle gradle.properties settings.gradle
RUN gradle clean build release
COPY . /app

## Runner stage
FROM ubuntu as runner
COPY --from=builder /app/dist/web-app-1.0.1.BUILD.war app/webapps/web-app.war
```

## Processes
The final step in the `Dockerfile` is to run the application using a process command. The recommendation is to separate areas of concern by using **one process per container**.

- **Simplicity** - Running multiple process within a container requires additional scripting within the container to coordinate the startup of each process.
- **Reliability** - When all process for a container exit, the container orchestrator can automatically restart the container. If there is more than one process, and it exits, it is up to you to manage the restart of the process within the container rather than allowing the orchestrator to restart it.
- **Scalability** - If one process requires additional capacity and triggers an auto scaling event, then all process within the container are also scaled.

The recommendation is the exec form of the `ENTRYPOINT` instruction to run the one process:

```docker
ENTRYPOINT ["java", "-jar", "/app.jar"]
```

## Dockerhub

[Dockerhub](https://hub.docker.com/) is an online docker image registry where users can upload their ready-built images, or pull down images uploaded by the docker community. Similarly to atlas for Vagrant boxes, dockerhub is driven by its users. If there isn't an image available with the features you want, you can make one and upload it yourself.

# Best Practices

The following best practices are part of a philosophy for containers which make them scale horizontally easier, be more reusable and deploy quicker.

 - Ephemeral Containers: A container should, as much as possible, be able to be stopped and destroyed, then rebuilt and replaced with an absolute minimum set up and configuration.
 - Decouple Applications: A container should have only one concern.

See [Docker: Dockerfile Best Practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) for more best practices regarding performance and optimization.

# Exercise

Create a Jenkins and Artifactory Docker image similar to the virtual machines you created using VirtialBox.

1. Create a `Dockerfile` in a new folder

2. Add the `FROM` instruction to use [Alpine Linux](https://hub.docker.com/_/alpine) as the base image.

3. Install Jenkins and any necessary dependencies.

?> You will use the `RUN` instruction for most of these steps but you should also familiarize yourself with other Dockerfile instructions.

4. `EXPOSE` the Jenkins web UI port.

5. Add a non privileged user.

6. Add an `ENTRYPOINT` instruction to run Jenkins as the non privileged user.

7. Build the image and run the container. Verify you can access the Jenkins web UI.

?> You will need to play with some of the `docker run` arguments so you can access the UI locally.

Repeat the steps for Artifactory.

# Deliverable

Understand the basic concepts of containers.

Be familiar with the Dockerfile instructions used to build an image.

Understand the best practices for Docker images and containers.

Discuss the differences between running `vagrant init ubuntu/bionic64 && vagrant up && vagrant ssh` and `docker run -it ubuntu:bionic bash`

Discuss the advantages / disadvantages of using containers vs. virtual machines.

Discuss some possible challenges of using Containers in an enterprise production environment.
