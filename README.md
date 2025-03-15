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

The message in the spec is used only to populate a string in the status. Additionally, the status also contains a copy of the pod's current phase

```yaml
status:
  specEcho: "I'm just a dummy"
  podStatus: "Running"
```

## Installing Kind

To run the image, you will need a Kubernetes cluster (for example, one created with [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/)).
You can use [Homebrew](https://brew.sh/) (mac, linux) or [Chocolatey](https://chocolatey.org/) (windows) to install Kind. Homebrew is installed via this command:

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

To install Chocolatey on windows, you must ensure `Get-ExecutionPolicy` is not `Restricted`. Use `Bypass` to bypass the policy to get things installed or `AllSigned` for quite a bit more security.

In an elevated Powershell environment, run `Get-ExecutionPolicy`. If it returns `Restricted`, then run `Set-ExecutionPolicy AllSigned` or `Set-ExecutionPolicy Bypass -Scope Process`.

Now run the following command:

```sh
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

Chocolatey should be ready to use now.

To install Kind, run the relevant command:

MacOS/Linux: `brew install kind`

Windows: `choco install kind`

You can now create and run a local cluster with `kind create cluster`

## Test the controller

The simplest method to install the controller for testing is to clone the [dumy-controller repo](https://github.com/anynines/tmp-homework-ms) and from within it, run the following command:

```sh
make deploy
```

Give it a few minutes to download and create the pod. You can check pod status by using:

```sh
kubectl get pods -n tmp-homework-ms-system
```

A sample Dummy CR is also provided within the repo:

```sh
kubectl apply -f config/samples/interview_v1alpha1_dummy.yaml
```

You will need to install [kubectl](https://kubernetes.io/docs/tasks/tools/) to run the above commands.

<!-- # dummy-controller
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/dummy-controller:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/dummy-controller:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/dummy-controller:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/dummy-controller/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
 -->
