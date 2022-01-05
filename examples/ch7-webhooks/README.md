# 7.5 Admission Webhooks

## Webhooks in Kubernetes

Kubernetes offers admission webhooks to validate and mutate resources before they are provisioned and stored in etcd. In this section we will examine the two types of admission webhooks: validating and mutating. 

## Validating and Mutating Webhooks 
Take a scenario where you are a developer looking to create a new pod on your cluster. For this example let's use a simple nginx pod definition. 

```
apiVersion: v1
kind: Pod
metadata: 
    name: nginx-example
spec:
    containers: 
        - image: nginx:latest
          name: nginx-example
```

How does this pod actually get created? This pod is created when the pod definition is applied- resulting in the kubelet sending an api request to the kube-apiserver, where the resource is authenticated, enabling the pod to persist in etcd. However, what if there are problems with this pod? What if a cluster admin wants to require that all pods created in a namespace must have specific labels, or that they need resource defaults associated with them? As our nginx pod stands now we would not want to actually create this pod in our org's cluster. So how can we enforce policy to catch these errors and others? 

![](img7/api-lifecycle.svg ':class=img-center')

Kubernetes gives us a chance to interact with a resource before it is persisted and provisioned. We can achieve this by using validating and mutating webhooks. A mutating webhook allows us to make changes to a resource before it is persisted in etcd. For example, if we pushed up a pod without resource defaults, our mutating webhook would catch this and inject a resource block (requests and limits for memory and cpu) into our pod. A validating webhook does not let us edit resources. Instead it gives us the ability to enforce policy over a cluster. For instance, we could have a requirement in our validating webhook requiring all namespaced pods have the label `foo:bar` in their metadata, and if they don't we reject the resource. 

Mutating and validating webhooks are used in conjunction to enforce policy. The mutating webhook happens first so we can inject any missing fields that the validating webhook requires before actually getting to the validating webhook step. These are both custom webhook configurations. As cluster admin we would write our policy code in whatever language we prefer (Golang and Python are the most widely used).  

### Setup

This project is for setting up a basic Kubernetes validating Admission
Controller using python.

1. Fork the webhooks [repository](https://github.com/liatrio/bootcamp-webhooks).

2. Deploy a KIND cluster with Admission Controller enabled

   ```bash
   kind create cluster --config kind.yaml
   ```

3. Install cert-manager

   ```
   # Add the Jetstack Helm repository
   helm repo add jetstack https://charts.jetstack.io

   # Update your local Helm chart repository cache
   helm repo update

   # Install the cert-manager Helm chart (including cert-manager CRDs)
   helm install cert-manager \
   --namespace cert-manager \
   --create-namespace \
   --set installCRDs=true \
   --version v1.6.1 \
   jetstack/cert-manager
   ```

4. Create the validation namespace, the root CA, and self signed certificate by applying the `certs.yaml` file.

5. Get the base64 value of the ca.crt file in the secret
    ``` 
    CA=`kubectl -n validation get secret validation-ca-tls -o jsonpath='{.data.ca\.crt}'`
    ```

### Validating Webhooks
1. Add a rule to `validating-webhook.yaml` for the creation of pods. Ensure the apiGroups are "apps" and "".

2. Configure a validating webhook in `app/webhooks.py`. A reference is provided with `examples/warden-validating.py`.

#### Criteria:
- Only check pods in the `busybox` namespace.
- If any container's image isn't `busybox:latest`, reject the pod.

3. Create a namespace called `busybox`.

4. Build the container using the Dockerfile within the directory. Push the image to your image repository.

5. Load the image into KIND.

6. Update the `warden-k8s.yaml` to point to your new image and apply.

7. Apply the `validating-webhook.yaml` file to deploy the validation configuration to the Kubernetes API server.
    ```
    cat validating-webhook.yaml | sed "s/      caBundle: .*/      caBundle: ${CA}/" | kubectl apply -f -
    ```
    
8. Test your validating webhook
Tests 2, 4, and 5 should pass.
    ```kubectl apply -f test-pods```

9. Delete the pods that passed.

### Mutating Webhooks

1. Copy the same rule from `validating-webhook.yaml` to `mutating-webhook.yaml` for the creation of pods.

2. Configure a mutating webhook in `app/webhooks.py`. A reference is provided with `examples/warden-mutating.py`.

#### Criteria:
- Only check pods in the `busybox` namespace.
- If any container's image isn't using a `busybox` image, reject the pod. 
- If the pod is using a version of `busybox` other than `latest`, mutate the container to change the image to `busybox:latest`.

3. Build the container using the Dockerfile within the directory. Push the image to your image repository.

4. Load the image into KIND. Ensure your `warden` pod in the `validating` namespace is using the most recent version of your image (verify using the kubectl `describe` command). Delete and recreate your validating webhook and `warden-k8s.yaml` if necessary.

7. Apply the `mutating-webhook.yaml` file to deploy the validation configuration to the Kubernetes API server.
    ```
    cat mutating-webhook.yaml | sed "s/      caBundle: .*/      caBundle: ${CA}/" | kubectl apply -f -
    ```

8. Test your validating webhook
Tests 1, 2, 4, 5, and 7 should pass.
    <br>```kubectl apply -f test-pods```
