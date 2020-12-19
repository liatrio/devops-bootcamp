# Kubernetes

![](img2/kubernetes.svg ':size=723x702 :class=icon')

There are many different solutions for managing containers but Kubernetes has become one of the most popular and widely used platforms. It is part of a large ecosystem of tools and projects organized by the [Cloud Native Computing Foundation (CNCF)](https://www.cncf.io/) which drives a lot of the development and support for containerized applications. So what is Kubernetes?

> Kubernetes is a portable, extensible open-source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. Read more at [kubernetes.io: What is Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/).

If you are still having a difficult time understanding what Kubernetes is that is OK. Kubernetes is complicated and we are not going to try to explain all of it in this section. All that you really need to know right now is that you can use it to run Docker containers.


# Architecture

### Cluster

A Kubernetes **cluster** is a set of node machines for running containerized applications. The Kubernetes Master is a collection of three processes that run on a single node in your cluster, which is designated as the master node. Those processes are: kube-apiserver, kube-controller-manager and kube-scheduler.

![](img2/delivery-k8s-cluster.svg ':class=img-center')

### Node
The **nodes** in a cluster are the machines (VMs, physical servers, etc) that run the applications and workflows. The Kubernetes master controls each node.

![](img2/delivery-k8s-node.svg ':class=img-center')

# Design
Kubernetes is declarative by design. When you want to change something in the cluster, you simply create/update/delete the resources from etcd via the API server. Various controllers then respond to this change by converging the system to comply with the newly declared state. To best understand, check out this drawing from [Julia Evans](https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/):

![](img2/scenes-from-kubernetes-page1.svg ':size=600px :class=img-center')

?>**Declarative:** you declare what you want the system to do - a desired state - and the system will align towards that state. In Kubernetes, you create an API object to represent what you want the system to do. And all the components in the system work to drive towards that state, until the object is deleted.

## Pods
A **Pod** is the basic execution unit of a Kubernetes application – the smallest and simplest unit in the Kubernetes object model that you create or deploy to run containers.

![](img2/delivery-k8s-pods.svg ':class=img-center')

## Services
A **Service** is an abstraction which defines a logical set of Pods and a policy by which to access them. As pods are created, scaled and destroyed over an application's lifetime the node they are running on and their IP addresses will change. A Service provides a consistent way to communicate with a set of pods regardless of how many pods there are or where they are running.

![](img2/delivery-k8s-service.svg ':class=img-center')

## Deployments

Deployments abstract away the low level details of managing Pods. Behind the scenes, **Deployments** rely on **ReplicaSets** to manage starting, stopping, scaling, and restarting the Pods if they happen to go down for some reason. Below is an example illustration of a deployment.

![](img2/delivery-k8s-deploy.gif ':class=img-center')

# Exercises

### Exercise 1: Hello Minikube

Minikube is a single node Kubernetes cluster that can be run on a laptop or VM to do local development or just experiment with Kubernetes.

For this exercise follow the instructions in the [Hello Minikube](https://kubernetes.io/docs/tutorials/hello-minikube/#) tutorial to get some experience creating basic resources in Kubernetes. You can use Katacode for the exercise or [install Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) locally.

After completing this exercise you should be familiar with the basics of creating Kubernetes resources and using `kubectl` to inspect them. Be sure to spend some time experimenting with `kubectl` commands and exploring the Kubernetes cluster.

### Exercise 2: Deployments and Services

Docker for Desktop also includes a single node Kubernetes cluster similar to Minikube. This is convenient since we already have Docker for Desktop installed but keep in mind it is only available on Mac OS and Windows.

In this exercise you will write YAML files to define Deployment and Service resources to deploy your Jenkins and Artifactory containers to the Docker for Desktop cluster. Note that this is different from how your created resources in the Hello Minikube exercise. Here you will be defining the entire resource spec instead of creating it with a minimal configuration using command line arguments. 

Kubernetes resource definitions can be very complicated documents which have a very strict structure. You will need to dig into the Kubernetes documentation and fine other examples to be able to create them correctly. You will also need to use commands like `kubectl get`, `kubectl describe` and `kubectl logs` to be able to troubleshoot the resources you create.

After completing this exercise you should be familiar with Kubernetes Deployment, Replica Set, Pod and Service resources, how those resources relate to each other as well as using `kubectl` to create and inspect the resources. 

1. Start up the Docker for Desktop Kubernetes cluster by follow the instructions at [docs.docker.com](https://docs.docker.com/docker-for-mac/#kubernetes).

2. Create a file called `deployment-jenkins.yaml` and define a [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) to run your Jenkins container. Run `kubectl apply -f deployment-jenkins.yaml` and check that the Deployment was created and that the Pod is running.

  ?> As long as the container `imagePullPolicy` is set to `IfNotPresent` (default) or `Never` the Kubernetes cluster will try to find the container image in your local Docker images. You must make sure the image is built and the tag matches the what is in the Deployment spec. Read more about how Pods use [container images](https://kubernetes.io/docs/concepts/containers/images/).

3. Use `kubectl port-forward POD_NAME LOCAL_PORT:REMOTE_PORT` and open the Jenkins UI.

4. Create a file called `service-jenkins.yaml` and define a [Service](https://kubernetes.io/docs/concepts/services-networking/service/) with a selector for your Jenkins pod. Run `kubectl apply -f service-jenkins.yaml` and check that the Service was created.

5. Repeat steps 2 - 4 and create a Deployment and Service for Artifactory.

6. Configure a Jenkins job to build Spring Pet Clinic and deploy the artifact to Artifactory.

  ?> For Jenkins to communicate with Artifactory it should use the Artifactory Service you created but it still needs to know the network address to use. Kubernetes manages its own internal DNS service to make this easy. Read more about [DNS for Services and Pods](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/) to figure out what the Service network name is. 

!> Note that we did not create any kind of persistent storage for your Pods so if they are restarted you will loose your Jenkins jobs and any artifacts stored in Artifactory. Creating persistent volumes in Kubernetes is a more advanced topic we will leave for later.

!> Running a Kubernetes cluster on your laptop can use a lot of memory and CPU time. If you installed Minikube locally you should make sure to run `minikube stop` and select *Docker Desktop* -> *Kubernetes* -> *Disable local cluster* to stop the clusters to save your battery life.
