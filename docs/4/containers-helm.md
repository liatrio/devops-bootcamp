# Helm
![Helm](docs/img/helm-icon-color.png)

Helm is a package manager for Kubernetes. Helm's packaging format are called charts, which is a collection of files that describe related Kubernetes resources. Helm Charts are used to help install, manage, and upgrade the most complex Kubernetes applications.


## Exercise
1. Install Helm.

2. Create a new helm chart. Once you have created a helm chart, it will contain a templates folder, a charts folder, a values.yaml, and a chart.yaml. 

3. Modify the values.yaml file to reference the Docker image from the Kubernetes section. 

4. Change the container port to the port needed for our jenkins instance and also make sure that the app version is matching the tag of your instance as well.

5. Use the command `helm install` to install the chart you have configured.

?> Helm install has a --dry-run option which will output the potential configuration without deploying. Run this to see if your configuration is how you want it.

6. Use kubectl or k9s to look at the current pods which should contain the configured jenkins pod.


## Exercise 2

1. Create a new Helm chart.

2. Modify the values.yaml and chart.yaml to pull the Real World frontend and backend images from the previous section. Make the backend a dependency of the frontend in your values.yaml.

