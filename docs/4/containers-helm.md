# Helm

![](img4/helm-icon.svg ':size=125px')

Helm is a package manager for Kubernetes. Helm's packaging format are called charts, which is a collection of files that describe related Kubernetes resources. Helm Charts are used to help install, manage, and upgrade the most complex Kubernetes applications. 

We can manually create the files that Helm will use to make a chart. The required files that Helm will work with are `values.yaml` and `Chart.yaml`. The chart.yaml will dictate the general description of the chart and the values.yaml is the place where you can define the values that are used in the templates.

## Exercise 1

1. Install Helm.

2. Create a templates folder that will hold the Kubernetes resource templates. Create a `configmap.yaml` in that folder which will be processed by Helm to create a ConfigMap Kubernetes resource.

?> Good idea to find the documentation on creating a ConfigMap resource

3. Create a `values.yaml` and a `chart.yaml`

4. Fill out these two files with the necessary amount of parameters needed for a basic Helm chart.

?> [Helm Getting Started Documentation](https://helm.sh/docs/chart_template_guide/getting_started/) is a good place to start looking at to fill out those files. Do not start off with using `helm create` because that provides more than is necessary.

5. Run `helm install configmap . ` to apply the Helm chart to the Kubernetes cluster. 

6. Run `kubectl get configmap NAME -o yaml` to see the yaml that is applied from the Helm install step.

7. Run `helm template . ` to show the entire output that Helm is creating for you.

8. Go to `values.yaml` and change the amount of replicas to 2. Run `helm upgrade configmap .` to apply the changes, you can run `kubectl get pods ` to see that there are 2 configmap pods running.

## Exercise 2

1. If you don't already have a fork of the bootcamp, do that now. This exercise will build off of a Helm chart found in the `examples/ch4/helm/RealWorld` folder.

2. Modify the values.yaml to pull the Backend image that you pushed to Dockerhub from the previous exercise. 

3. Install the Helm Chart and use kubectl or k9s to check the status of the Pod created by Helm.

4. This pod will run into errors because the backend is reliant on MongoDB, Docker Compose set up MongoDB as a service. We need to do the same in Helm now.

5. Set MongoDB as a dependency for the Backend Helm Chart. The official Bitnami MongoDB Helm Chart is found [here](https://bitnami.com/stack/mongodb/helm) Make sure to run `helm dependencies update ` to download the mongoDB helm chart.

?> The MongoDB from Bitnami will likely be setup differently than how you set it up in Docker Compose from the previous section. Make sure to read through the [documentation](https://github.com/bitnami/charts/tree/master/bitnami/mongodb) to figure out how to pass in parameters to MongoDB.

6. The `deployment-backend.yaml` will need to be modified to include the correct MANGODB_URI.

7. Install the Helm Chart again once the deployment is up to date to handle the mongoDB dependency. You can use `kubectl logs` command to check the status of deployments. Once your pods are running, you can move onto Exercise 3.


## Exercise 3

1. Create `deployment-frontend.yaml` and `service-frontend.yaml` based on the `deployment-backend.yaml` and the `service-backend.yaml` created previously.

?> Some of the settings will be different for the frontend

2. Run `helm install NAME . ` to apply the Helm Chart to Kubernetes with the new configuration for the frontend.

?> Take note of what is output when you run helm install here

3. Once you see the frontend in a local browser, you can continue to the next exercise.

## Exercise 4

1. Create a new helm chart with `helm create`. Once you have created a helm chart, it will contain a templates folder, a charts folder, a values.yaml, and a chart.yaml. 

2. Modify the values.yaml file to reference the Docker image from the Kubernetes section. 

3. Change the container port to the port needed for our jenkins instance and also make sure that the app version is matching the tag of your instance as well.

4. Use the command `helm install` to install the chart you have configured.

?> Helm install has a --dry-run option which will output the potential configuration without deploying. Run this to see if your configuration is how you want it.

5. Use kubectl or k9s to look at the current pods which should contain the configured jenkins pod.


# Delieverables

