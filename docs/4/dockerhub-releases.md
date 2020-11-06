# Dockerhub

Another important use of creating releases is when releasing container images to Dockerhub. [Dockerhub](https://hub.docker.com/) is the default container registry for Docker. In short, this means that Dockerhub provides repositories to create, manage, and find container images. 

## Exercise

1. If you haven't already, create a [Dockerhub](https://hub.docker.com/) account, and create your first repository.
2. Using a `Dockerfile`, build a container image. Make sure to [tag](https://docs.docker.com/engine/reference/commandline/tag/) the image using semantic versioning.
  - Hint: You can tag a Docker image in the same step as the build.
3. Push this image to your Dockerhub repository.
4. Create a `Jenkinsfile` pipeline to automatically build and push your Docker image to your repository.
  - It is challenging to run Docker inside of Docker, so using a golden image will likely be easier.
  - You might need to use the `docker login` command in the Jenkinsfile so you can push to your Dockerhub repository.
  - Do not use the Docker Pipeline plugin.

## GitHub Container Registry

A popular alternative to Dockerhub is using GitHub Container Registry. GitHub Container Registry is similar to Dockerhub, but has a few differences. For one, since Dockerhub is the default container registry, no additional information is needed when you `pull` or `push` an image. Also, Dockerhub is more lenient with container size and usage, whereas GitHub Container Registry is typically better suited for organizations who want to manage access to container images in a centralized location. Lastly, GitHub Container Registry may be a better choice if you are using GitHub Actions.

One thing to note is that Dockerhub provides GitHub integration, where Dockerhub will build and publish an image after a commit to GitHub.

# Deliverable

Discuss the benefits of using a container registry, and how image releases work.