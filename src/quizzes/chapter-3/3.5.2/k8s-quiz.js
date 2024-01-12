const rawQuizdown = `
---
shuffleQuestions: true
shuffleAnswers: true
---

# Fill in the blanks: ____ is more suitable for simple, local development, while ____ is more suitable for complex, production environments

1. [x] Docker Compose, Kubernetes
1. [ ] Kubernetes, Docker Compose
1. [ ] Kubernetes, Docker
1. [ ] Docker, Kubernetes

# Order the following Kubernetes resources into the correct scopes (highest scope to lowest scope)

1. Cluster
2. Node
3. Pod

# True or False: Making an invalid command to the Kubernetes API can run the risk of breaking resources in your cluster

1. [ ] True
	> Not quite. Refer back to the [Design](https://devops-bootcamp.liatr.io/#/3-virtual-machines-containers/3.5.2-kubernetes?id=design) section above to get a firmer grasp on how Kubernetes handles changes to the cluster
1. [x] False
	> Excellent! The great thing about the Kubernetes engine is that *proposed* changes (your commands) must be validated by an API server before any changes are made to etcd

# Fill in the blanks: While Services and Deployments both are used to group pods, ____ are used for the starting/stopping/scaling of pods and ____ are used for the linking/networking of pods

1. [x] Deployments, Services
1. [ ] Services, Deployments

`;

export { rawQuizdown }