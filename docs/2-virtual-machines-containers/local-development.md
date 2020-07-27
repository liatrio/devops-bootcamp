# Local Development

An important use case for virtual machines is keeping parity between development and production environments. Developing and testing in an environment identical to the production environment helps reduce "It works on my machine" type of problems that happen all too frequently when releasing software.

So far the tools we have looked at for building and running images are great for automation and managing complicated infrastructure. But for developers that want to be able to easily create, use and share environments they are not very convenient. If only we could create and start a virtual machine with just two commands...

#### Exercise 1: VMs the Easy Way

1. [Install Vagrant](https://www.vagrantup.com/intro/getting-started/install).

2. Create a new folder and make it your current directory.

3. Run `vagrant init centos/7`

4. Run `vagrant up`

5. Run `vagrant ssh`

  ... welcome to your new VM :)

6. Exit out of the VM and run `vagrant destroy` to clean up.

## Vagrant

> Vagrant is a tool for building and managing virtual machine environments in a single workflow -- [Introduction to Vagrant](https://www.vagrantup.com/intro/index)

Being able to to create a CentOS 7 VM in two commands is pretty cool but what if you want to customize your VM and share it with others. Imagine your working with a team of developers and you want the entire team to be able use a VM with the same libraries and tools installed.

#### Exercise 2: Creating a Base Box

In this exercise we will create a Vagrant Base Box using our VirtualBox Golden Image. Vagrant uses VirtualBox (or other providers) to run images. In order for Vagrant to be able to do this the image needs to be configured for Vagrant. Read [Creating a Base Box](https://www.vagrantup.com/docs/boxes/base) and [Creating a (VirtualBox) Base Box](https://www.vagrantup.com/docs/providers/virtualbox/boxes.html).

1. Clone your VirtualBox Golden Image and update it to meet the requires for a Vagrant VirtualBox Base Box.

2. Run `vagrant package ...` to create your base box.

3. Upload your base box to [VagrantUp](https://app.vagrantup.com/).

<center>

  ![](img2/config.svg ':size=125px')

</center>

# Deliverable

Also, discuss the following topics.
 - What could Vagrant be used for?
 - What are base boxes for?
 - What was the hardest part of this exercise?
