# Delivery Pipeline
  ![](img2/delivery-pipeline-stages-1.png)

  Automated delivery pipelines accelerate software delivery and enable faster, safer deployments.
  An entire integrated toolchain is necessary to leverage the capabilities of an automated pipeline, orchestrated by a build server. Newly committed code that is pushed to the git server will send a Webhooks to the build servers to trigger tasks. General stages of a delivery pipeline include:
- Unit tests
- Code quality scanning
- Build artifacts (binaries, dist, containers)
- Publishing artifacts to an artifact repository
- Functional testing
- Deploying artifacts

### Artifact storage
  Artifacts generated in a build during continuous delivery pipeline needs to be stored in an artifact repository for future use. This provides one location for build and deploy artifacts like application binaries, containers, and manifest. By providing a centralized location secured and owned by the organization, the artifact repository acts as a single source for dependency management.

### Deployments
  Applications should be deployed using a repeatable process within the continuous delivery pipeline. Using GitOps principals, artifacts defined in a manifest file reflects what is in each environment, providing visibility and traceability for infrastructure and applications. 