# Local Development

An important use case for virtual machines is keeping parity between development and production environments. Developing and testing in an environment identical to the production environment helps reduce "It works on my machine" type of problems that happen all too frequently when releasing software.

So far the tools we have looked at for building and running images are great for automation and managing complicated infrastructure. But for developers that want to be able to easily create, use and share environments they are not very convenient. If only we could create and start a virtual machine with just two commands...

<center>

  ![](img2/config.svg ':size=125px')

</center>

#### Exercise 1: VMs the easy way

1. [Install Vagrant](https://www.vagrantup.com/intro/getting-started/install).

2. Create a new folder inside of it run `vagrant init centos/7`

3. Run `vagrant up`

    **Congratulations!!!** You have a new VM up and running!

4. Run `vagrant ssh` to access the new VM.

5. Exit out of the VM and run `vagrant destroy` to clean up.

## Vagrant

> Vagrant is a tool for building and managing virtual machine environments in a single workflow -- [Introduction to Vagrant](https://www.vagrantup.com/intro/index)

Being able to to create a CentOS 7 VM in two commands is pretty cool but what if you want to customize your VM and share it with others. Imagine your working with a team of developers and you want the entire team to be able use a VM with the same libraries and tools installed.

#### Exercise 2: Creating a Base Box

In this exercise we will create a Vagrant Base Box using our VirtualBox Golden Image. Vagrant uses VirtualBox (or other providers) to run images. In order for Vagrant to be able to do this the image needs to be configured for Vagrant. Read [Creating a Base Box](https://www.vagrantup.com/docs/boxes/base) and [Creating a (VirtualBox) Base Box](https://www.vagrantup.com/docs/providers/virtualbox/boxes.html).

1. Clone your VirtualBox Golden Image and update it to meet the requires for a Vagrant VirtualBox Base Box.

2. Run `vagrant package ...` to create your base box.

3. Upload your base box to [VagrantUp](https://app.vagrantup.com/).

## Packer & Vagrant

You may have noticed some similarities between Packer and Vagrant. They both make use of a variety of virtualization platforms. They both have provisioners to configure machine images. This might be surprising considering both tools are made by HashiCorp. To understand the differences between these tools it can be helpful to consider the problem they are trying to solve. Packer is primarily used to create machine images, often from scratch. Vagrant's use cases focus more on making machine images easier to use and manage. Another way to understand the relationship between these tools is knowing that many of the official Vagrant Base Boxes are actually created using Packer.

#### Exercise 3: Creating a Base Box using Packer

In this exercise you will use the Vagrant Post-Processor to build a Vagrant Box using the Packer template you created in the previous section.

1. Modify your [kickstart file](https://docs.centos.org/en-US/centos/install-guide/Kickstart2/) to automate the install of CentOS and meet Vagrant's requirements.

    - allow `vagrant` user to have password-less sudo
    - `UseDNS no` (optional)
    - update the CentOS packages and kernel

2. Use Packer provisioners to finish setting up the system.

    - install Vagrant [insecure key pair](https://github.com/hashicorp/vagrant/tree/master/keys) for `vagrant` user
    - install dependencies needed for VirtualBox Guest Additions
    - install Guest Additions

3. Add a [Vagrant post-processor](https://www.packer.io/docs/templates/post-processors) to your Packer build template. Configure it to output both a VirtualBox image and a Vagrant box.


# Deliverable

Also, discuss the following topics.
 - What could Vagrant be used for?
 - What are base boxes for?
 - What was the hardest part of this exercise?
