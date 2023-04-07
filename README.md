<center>

  ![Liatrio Logomark](img/favicon.svg ':size=150x150 :class=logo :alt= Liatrio image')

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

![classroom image](img/class_light.svg ':size=100x100 :class=light-mode-icon :alt= classroom image; light mode')
![classroom image](img/class_dark.svg ':size=100x100 :class=dark-mode-icon :alt= classroom image; dark mode')