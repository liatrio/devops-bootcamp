---
docs/1-introduction/1.1-devops-defined.md:
  category: Fundamentals
  estReadingMinutes: 30
docs/1-introduction/1.2-liatrio-and-devops.md:
  category: Fundamentals
  estReadingMinutes: 5
docs/1-introduction/1.2.1-dojo.md:
  category: Fundamentals
  estReadingMinutes: 5
docs/1-introduction/1.3.1-unix.md:
  category: Fundamentals
  estReadingMinutes: 30
docs/1-introduction/1.3.2-vim.md:
  category: Fundamentals
  estReadingMinutes: 15
  technologies:
    - Vim
docs/1-introduction/1.3.3-installations.md:
  category: Fundamentals
  estReadingMinutes: 10
docs/1-introduction/1.3.4-passwords-and-keys.md:
  category: Fundamentals
  estReadingMinutes: 25
  technologies:
    - ssh
docs/1-introduction/1.3.5-networking.md:
  category: Fundamentals
  estReadingMinutes: 25
docs/1-introduction/1.4-external-resources.md:
  category: Fundamentals
  estReadingMinutes: 50
docs/2-virtual-machines-containers/2.1-golden-images.md:
  category: Virtualization
  estReadingMinutes: 15
  exercises:
    - name: hello VirtualBox
      description: >-
        Configure a VM with VirtualBox. Install some software and configure it
        for SSH
      estMinutes: 90
      technologies:
        - VirtualBox
        - CentOS
    - name: hello Packer
      description: >-
        Use Packer to create an OVF configured identically to the VM you
        configured in exercise 1
      estMinutes: 210
      technologies:
        - VirtualBox
        - Packer
        - CentOS
docs/2-virtual-machines-containers/2.2-local-development.md:
  category: Virtualization
  estReadingMinutes: 15
  exercises:
    - name: hello Vagrant
      description: Install Vagrant spin up CentOS 7 box
      estMinutes: 60
      technologies:
        - Vagrant
        - CentOS
    - name: Vagrant base box
      description: >-
        Create a ‘Base Box’ that adheres to Vagrants ‘Base Box’ standards.
        Upload this base box to VagrantUp
      estMinutes: 180
      technologies:
        - Vagrant
        - CentOS
    - name: Packer to Vagrant
      description: >-
        Use Packer to create an image compatible with Vagrants Base Box
        requirements. This requires modification of CentOS’s ‘kickstart’ file
        (understand CentOS kickstart DSL), leverage a Packer provisioner to add
        ssh key, install VirtualBox dependencies and tools, leverage Packer
        post-processor to output an OVF and boot a vagrant box
      estMinutes: 300
      technologies:
        - Vagrant
        - Packer
        - CentOS
docs/2-virtual-machines-containers/2.3-managing-infrastructure.md:
  category: Virtualization
  estReadingMinutes: 10
  exercises:
    - name: Jenkins and Artifactory on VMs
      description: >-
        Create VMs from your golden image. On one VM install Jenkins and
        configure workflow to build Spring’s PetClinic. On the other VM install
        and configure Artifactory. Configure these VMs so they can communicate
        and create another workflow that pushes the build artifact from the
        Jenkins server to Artifactory
      estMinutes: 390
      technologies:
        - Vagrant
        - CentOS
        - Jenkins
        - Artifactory
docs/2-virtual-machines-containers/2.4-containers.md:
  category: Containerization
  estReadingMinutes: 20
  exercises:
    - name: Hello Containers
      description: Complete Docker's 'Hello World'
      estMinutes: 30
      technologies:
        - Docker
    - name: Jenkins to Artifactory containerized
      description: >-
        Containerize Jenkins and Artifactory and build a pipeline that pushes a
        PetClinic build artifact from one container to the other.
      estMinutes: 150
      technologies:
        - Docker
        - Jenkins
        - Artifactory
docs/2-virtual-machines-containers/2.5.1-docker-compose.md:
  category: Container Orchestration
  estReadingMinutes: 5
  exercises:
    - name: Docker Compose Jenkins and Artifactory
      description: >-
        Create a compose file that spins up your Jenkins and Artifactory
        containers, exposing ports, and using volumes for persistent storage
      estMinutes: 120
      technologies:
        - Docker
        - Docker Compose
        - Jenkins
        - Artifactory
docs/2-virtual-machines-containers/2.5.2-kubernetes.md:
  category: Container Orchestration
  estReadingMinutes: 15
  exercises:
    - name: Hello Minikube
      description: Complete the 'Hello Minikube' tutorial
      estMinutes: 60
      technologies:
        - Kubernetes
    - name: Kind cluster Jenkins and Artifactory
      description: Create a Kind cluster running your Jenkins and Artifactory containers
      estMinutes: 600
      technologies:
        - Kubernetes
        - Jenkins
        - Artifactory
---
<!-- DO NOT EDIT ABOVE THIS COMMENT. -->
<!-- MANAGED BY ./husky.sh/front-matter-condenser.js -->

<canvas id="chart-canvas"></canvas>

<center>

  ![Liatrio Logomark](../img/favicon.svg ':size=150x150 :class=logo :alt= Liatrio image')

# Liatrio's DevOps Bootcamp

</center>

## Introduction

Welcome! You must be here because you're interested in DevOps. Don't worry, we'll get to that in a moment! First, let's briefly go over some details.

"1.0" can be referred to as "Chapter One," while "1.2" can be referred to as "Chapter One, Section Two."

## Target Audience

This DevOps Bootcamp is used as an introduction to DevOps for Liatrio's apprentices. Therefore, there are deliverables at the end of each section that help guide the apprentices' training. However, anyone can go through the DevOps Bootcamp. Feel free to ignore the deliverables, etc. This bootcamp also assumes the bootcamper is using a Mac for the deliverables.

## Disclaimer

This bootcamp alone will by no means make anyone an expert on DevOps; that's where real, in-the-field experience engineering and shadowing come in. However, this bootcamp is a great start for anyone entering the world of DevOps! It is recommended that anyone going through this bootcamp have a mentor to fill in the gaps of information and answer questions, as well. Also, many of the exercises proposed in this bootcamp have no solution listed. This is because it is up to the bootcamper to come up with the solution on their own. So, if you're wondering where a solution to an exercise is - there is likely none; that's up to you!

## Learning Goals

 1. Introduction to DevOps
 2. Virtual Machines and Containers
 3. Cloud Computing
 4. Development Practices
 5. Release Management
 6. Automation and Orchestration
 7. Kubernetes

## Expectations

 1. Share what you learn. Knowledge is best retained when you can teach it back yourself.
 2. If you're participating in the bootcamp with others, lift them up instead of speeding past them.
 3. Your learning should be in a demonstrable state at any given time.
 4. Slow and steady wins the race. If you don't fully understand something, speak up for help and spend more time on it.
 5. Be transparent in your progress.
 6. Have fun!

> "It's better to over communicate than to under communicate."

## Contributing

This bootcamp was adapted from an apprentice training program, and some mistakes may have slipped through the cracks. If you see any errors, outdated methods, or citation errors please submit a pull request.

## Local Development

Below are a number of ways to develop locally, choose whichever you are most comfortable with.

### Using Package.json **Recommended**

Install Docsify locally (in current directory)

1. Run `npm install`
2. Run `npm start`
3. Open <http://localhost:3000>

### Global Install

Install Docsify Globally in your system

1. Install [docsify](https://docsify.js.org/#/quickstart)
2. Navigate to local library of onboarding
3. Run `docsify serve .`
4. Open <http://localhost:3000>

### Docker

Use Docker to build and serve the content, but remember to rebuild the Docker image to review changes

#### Build and Run Docker Container

1. Execute `docker build . -t devops-bootcamp` from the project's root directory to generate a container image
2. Run `docker run -d -p 3000:3000 --name devops-bootcamp devops-bootcamp` to run a detached Docker container
3. Open <http://localhost:3000>

#### Docker Compose

1. Run `docker compose up -d`
2. Open <http://localhost:3000>

### More Information on Contributing

- Images should be placed under the root `img` folder and referred to using HTML `<img>` tags
- H3 header (`###`) should be the default header within a page
- H2 header (`##`) will appear in the navigation as the page's table of contents

[Contributors](contributors.md)

![classroom image](../img/class_light.svg ':size=100x100 :class=light-mode-icon :alt= classroom image; light mode')
![classroom image](../img/class_dark.svg ':size=100x100 :class=dark-mode-icon :alt= classroom image; dark mode')
