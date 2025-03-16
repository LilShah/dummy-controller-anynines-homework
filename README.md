# dummy-controller

## Overview

Dummy Controller is a simple example Kubernetes controller that creates an nginx pod for every Dummy custom resource created.

```yaml
apiVersion: interview.interview.com/v1alpha1
kind: Dummy
metadata:
  name: dummy1
  namespace: example
spec:
  message: "I'm just a dummy"
```

The message in the spec is used to populate a string in the status. Additionally, the status also contains a copy of the pod's current phase

```yaml
status:
  specEcho: "I'm just a dummy"
  podStatus: "Running"
```

## Requirements

The image for this operator is pushed on Docker Hub at [lilshah/dummy-controller](https://hub.docker.com/repository/docker/lilshah/dummy-controller/general). To run it, a kubernetes cluster is required. Mini kubernetes clusters can be used as well like [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/). Additionally, [kubectl](https://kubernetes.io/docs/tasks/tools/) is required to interface with the cluster.

## Running from repo

Since just running the given image isn't enough to run the entire operator, the best way to run it is via this repo. Running the following command will deploy the controller, RBAC, CRD and all other manifests into the cluster:

```sh
make deploy
```

Give it a few minutes to download and create the pod. You can check pod status by using:

```sh
kubectl get pods
```

Sample Dummy CRs are also provided within the repo at `config/samples`. They can be used to test the controller:

```sh
kubectl apply -f config/samples/dummy.yaml
```
