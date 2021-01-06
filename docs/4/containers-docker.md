# Dockerhub

Another important use of creating releases is when releasing container images to Dockerhub. [Dockerhub](https://hub.docker.com/) is the default container registry for Docker. In short, this means that Dockerhub provides repositories to create, manage, and find container images. 

## Exercise

1. If you haven't already, create a [Dockerhub](https://hub.docker.com/) account, and create your first two repositories, name them realworld_backend and realworld_frontend.
2. Using two different `Dockerfiles` build two container images, one frontend [(React/Redux)](https://github.com/gothinkster/react-redux-realworld-example-app) and one backend [(Node/Express)](https://github.com/gothinkster/node-express-realworld-example-app) from the [Real World Repository](https://github.com/gothinkster/react-redux-realworld-example-app). Make sure to [tag](https://docs.docker.com/engine/reference/commandline/tag/) the image using semantic versioning.
  - Hint: You can tag a Docker image in the same step as the build.
3. Push both images to your Dockerhub repository.
4. Create a `Docker Compose` to be able to connect the frontend and the backend.
?> The Docker Compose will need to configure the MongoDB instance for the backend
5. Create a `Jenkinsfile` pipeline to automatically build and push your Docker images to your repository.
  - It is challenging to run Docker inside of Docker, so using a golden image will likely be easier.
  - You might need to use the `docker login` command in the Jenkinsfile so you can push to your Dockerhub repository.
  - Do not use the Docker Pipeline plugin.

## GitHub Container Registry

A popular alternative to Dockerhub is using GitHub Container Registry. GitHub Container Registry is similar to Dockerhub, but has a few differences. For one, since Dockerhub is the default container registry, no additional information is needed when you `pull` or `push` an image. Also, Dockerhub is more lenient with container size and usage, whereas GitHub Container Registry is typically better suited for organizations who want to manage access to container images in a centralized location. Lastly, GitHub Container Registry may be a better choice if you are using GitHub Actions.

One thing to note is that Dockerhub provides GitHub integration, where Dockerhub will build and publish an image after a commit to GitHub.

# Deliverable

Discuss the benefits of using a container registry, and how image releases work.
