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
docs/3-cloud-computing/3.0-overview.md:
  category: Cloud Computing
  estReadingMinutes: 30
docs/3-cloud-computing/3.1.1-aws.md:
  category: Cloud Computing
  estReadingMinutes: 60
docs/3-cloud-computing/3.1.2-azure.md:
  category: Cloud Computing
  estReadingMinutes: 30
docs/3-cloud-computing/3.2.1-s3-cloudfront.md:
  category: Cloud Computing
  estReadingMinutes: 20
  exercises:
    - name: Hello S3
      description: Host a static website in s3 publicly
      estMinutes: 150
      technologies:
        - AWS S3
        - AWS
    - name: Hello Cloudfront
      description: Create a CloudFront distribution for your s3 website
      estMinutes: 60
      technologies:
        - AWS S3
        - AWS
        - CloudFront
docs/3-cloud-computing/3.2.2-ec2.md:
  category: Cloud Computing
  estReadingMinutes: 10
  exercises:
    - name: Hello EC2
      description: >-
        Create EC2 VMs and configure one as a Jenkins server and the other as a
        Jenkins agent registered to the server.
      estMinutes: 240
      technologies:
        - AWS
        - EC2
        - Jenkins
docs/3-cloud-computing/3.2.3-auto-scaling.md:
  category: Cloud Computing
  estReadingMinutes: 20
  exercises:
    - name: Create EC2 instance with User Data
      description: Create an EC2 instance running Spring PetClinic.
      estMinutes: 240
      technologies:
        - AWS
        - EC2
    - name: Setup a Network Load Balancer
      description: >-
        Create a Load Balancer in AWS and register 2 EC2 configured to host a
        web application via User Data.
      estMinutes: 60
      technologies:
        - AWS
        - EC2
        - AWS ELB
    - name: Launch Templates
      description: >-
        Create a Launch Template that will create EC2 instances configured like
        the vm you made in exercise 1
      estMinutes: 60
      technologies:
        - AWS
        - EC2
        - AWS Launch Templates
    - name: Auto Scaling Groups
      description: >-
        Create an AutoScaling group leveraging the Launch Template you made in
        exercise 3
      estMinutes: 60
      technologies:
        - AWS
        - EC2
        - AWS Launch Templates
        - AWS Auto Scaling Groups
        - AWS ELB
docs/3-cloud-computing/3.2.4-aws-packer.md:
  category: Cloud Computing
  estReadingMinutes: 20
  exercises:
    - name: Creating a custom AMI
      description: >-
        Create an AMI with Packer that creates a provisioned machine like the
        one you made in exercise 1 from 3.2.3. Create a Launch Template and Auto
        Scaling group to leverage your new AMI via the AWS cli
      estMinutes: 180
      technologies:
        - AWS
        - EC2
        - AWS Launch Templates
        - AWS Auto Scaling Groups
        - Packer
docs/3-cloud-computing/3.2.5-lambda.md:
  category: Cloud Computing
  estReadingMinutes: 20
  exercises:
    - name: Creating a Lambda function with a trigger
      description: >-
        Create a Lambda function that sets up an s3 bucket and a DynamoDB,
        downloads a csv file, and then uses a python script to parse the csv and
        push all the data into DynamoDB.
      estMinutes: 480
      technologies:
        - AWS
        - AWS Lambda
        - AWS S3
        - AWS DynamoDB
docs/3-cloud-computing/3.2.6-ecs.md:
  category: Cloud Computing
  estReadingMinutes: 20
  exercises:
    - name: Creating an EKS cluster
      description: Deploy the SockShop Demo on EKS
      estMinutes: 480
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
    - name: Adding an autoscaler to your cluster
      description: Adding an autoscaler to your cluster
      estMinutes: 240
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
    - name: Create a ClusterIP service
      description: Create a ClusterIP service
      estMinutes: 60
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
    - name: Create a NodePort Service
      description: Create a NodePort Service
      estMinutes: 60
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
    - name: Create a LoadBalancer Service
      description: Create a LoadBalancer Service
      estMinutes: 30
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
    - name: Add an Ingress controller and your cluster
      description: Add an Ingress controller and your cluster
      estMinutes: 60
      technologies:
        - AWS
        - AWS EKS
        - Kubernetes
docs/3-cloud-computing/3.3.1-storage-accounts.md:
  category: Cloud Computing
  estReadingMinutes: 10
  exercises:
    - name: Static website on Azure Blob Storage
      description: >-
        Create a basic Angular application and deploy it to an Azure Blob
        Storage with an Azure CDN.
      estMinutes: 240
      technologies:
        - Azure
        - Azure Blob Storage
        - Azure CDN
        - Angular
docs/3-cloud-computing/3.3.2-virtual-machines.md:
  category: Cloud Computing
  estReadingMinutes: 10
  exercises:
    - name: Create Azure VM with NGINX via cloud-init
      description: >-
        Create an Azure VM and configure it as an NGINX web server via
        cloud-init file.
      estMinutes: 360
      technologies:
        - Azure
        - Azure VM
        - NGINX
        - cloud-init
docs/3-cloud-computing/3.3.3-vmss.md:
  category: Cloud Computing
  estReadingMinutes: 15
  exercises:
    - name: Create VMSS and serve content from Azure Storage Account
      description: >-
        Create a VMSS and serve content from an Azure Storage Account. Do this
        via the cli, deploying a simple node web app and provision the VM's with
        cloud-init.
      estMinutes: 240
      technologies:
        - Azure
        - Azure VMSS
        - Azure Storage Account
        - Node.js
        - cloud-init
docs/3-cloud-computing/3.3.4-az-packer.md:
  category: Cloud Computing
  estReadingMinutes: 5
  exercises:
    - name: Create Shared Image Gallery with custom Azure VM image
      description: >-
        Create a Shared Image Gallery and put a custom Azure VM image in it.
        Then redo the exercise in 3.3.2 with the custom image.
      estMinutes: 360
      technologies:
        - Azure
        - Packer
        - Shared Image Gallery
        - Azure VM
docs/3-cloud-computing/3.3.5-aci.md:
  category: Cloud Computing
  estReadingMinutes: 10
  exercises:
    - name: 'Deploy Node app using Docker, ACR, and ACI'
      description: >-
        Create a docker image for your Node application. Push your image into
        ACR. Then deploy your image via ACI.
      estMinutes: 300
      technologies:
        - Azure
        - Azure Container Instances (ACI)
        - Azure Container Registry (ACR)
        - Docker
        - Node.js
docs/3-cloud-computing/3.3.6-aks.md:
  category: Cloud Computing
  estReadingMinutes: 10
  exercises:
    - name: Deploy Node app to AKS using ACR image
      description: >-
        Create an AKS cluster and deploy the Node application via the image
        uploaded to ACR made in the previous section.
      estMinutes: 240
      technologies:
        - Azure
        - Azure Kubernetes Service (AKS)
        - Azure Container Registry (ACR)
        - Node.js
docs/4-software-development-practices/4.1-overview.md:
  category: Agile Development
  estReadingMinutes: 90
docs/4-software-development-practices/4.2-jira.md:
  category: Agile Development
  estReadingMinutes: 30
docs/4-software-development-practices/4.3.1-branching-merging.md:
  category: Version Control
  estReadingMinutes: 15
docs/4-software-development-practices/4.3.2-git.md:
  category: Version Control
  estReadingMinutes: 25
  exercises:
    - name: Git Basics
      description: Learn the basics of Git.
      estMinutes: 30
      technologies:
        - Git
    - name: Branching and Merging
      description: Practice branching and merging with Git.
      estMinutes: 30
      technologies:
        - Git
    - name: Merge Conflicts
      description: Learn how to resolve merge conflicts in Git.
      estMinutes: 30
      technologies:
        - Git
    - name: Removing Secrets
      description: Learn how to remove secrets from Git history.
      estMinutes: 30
      technologies:
        - Git
docs/4-software-development-practices/4.3.3-github.md:
  category: Version Control
  estReadingMinutes: 25
  exercises:
    - name: Fork and Merge
      description: Practice forking and merging on GitHub
      estMinutes: 30
      technologies:
        - Git
        - GitHub
docs/4-software-development-practices/4.4-pairprogramming.md:
  category: Agile Development
  estReadingMinutes: 20
  exercises:
    - name: Pair Programing
      description: >-
        Using 'Live Share' or some equivillant try pair programming a 'Hello
        World' app in the language of your choice
      estMinutes: 30
      technologies:
        - VSCode
docs/4-software-development-practices/4.5.1-unit-testing.md:
  category: Software Quality
  estReadingMinutes: 20
  exercises:
    - name: Solve a problem in Go with TDD
      description: Solve a given problem in Go using test-driven development (TDD).
      estMinutes: 120
      technologies:
        - Go
        - TDD
docs/4-software-development-practices/4.5.2-functional-testing.md:
  category: Software Quality
  estReadingMinutes: 20
  exercises:
    - name: Create functional tests with Selenium in Python
      description: Create functional tests for the bootcamp using Selenium in Python.
      estMinutes: 120
      technologies:
        - Selenium
        - Python
docs/4-software-development-practices/4.5.3-code-styling.md:
  category: Software Quality
  estReadingMinutes: 30
docs/4-software-development-practices/4.5.4-code-coverage.md:
  category: Software Quality
  estReadingMinutes: 20
  exercises:
    - name: Achieve 100% code coverage
      description: >-
        Install node dependencies for code coverage (mocha, chai, nyc) and get
        100% code coverage for a given module and tests.
      estMinutes: 60
      technologies:
        - Node.js
        - Mocha
        - Chai
        - NYC
docs/4-software-development-practices/4.5.5-test-automation.md:
  category: Software Quality
  estReadingMinutes: 5
  exercises:
    - name: Create a Jenkins Pipeline to test Go project
      description: >-
        Create a Jenkins server via an existing Dockerfile, add Go as a
        dependency to the Image. On the Jenkins instance point to your fork of
        the Bootcamp with Go Unit Tests. Create a pipeline that builds and tests
        your fork.
      estMinutes: 180
      technologies:
        - Jenkins
        - Docker
        - Go
docs/4-software-development-practices/4.5.6-sonarqube.md:
  category: Software Quality
  estReadingMinutes: 10
  exercises:
    - name: Setup SonarQube and Jenkins Integration
      description: >-
        Create a SonarQube server and add a Jenkins job to run SonarQube in our
        build pipeline.
      estMinutes: 180
      technologies:
        - SonarQube
        - Jenkins
docs/4-software-development-practices/4.6-hello-devops.md:
  category: Agile Development
  estReadingMinutes: 5
  exercises:
    - name: Build an application
      description: >-
        Using what you have leanred thus far create a production ready
        application
      estMinutes: 3000
---
<!-- DO NOT EDIT ABOVE THIS COMMENT. -->
<!-- MANAGED BY ./husky.sh/front-matter-condenser.js -->

<!-- <canvas id="wordcloud-canvas"></canvas> -->

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

<canvas id="category-doughnut-canvas"></canvas>

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
