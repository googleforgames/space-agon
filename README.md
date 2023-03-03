The original work is [Laremere/space-agon](https://github.com/Laremere/space-agon).

# Space Agon

Space Agon is a integrated demo of [Agones](https://agones.dev/) and
[Open Match](https://open-match.dev/).

## Before Trying.

**Be aware of billing charges for running the cluster.**

Space Agon is intended to run on [Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine) and has been tested with the configured cluster size.   
Leaving the cluster running may incur your cost. You need to be responsible for the cost.  (See pricings of [GKE](https://cloud.google.com/kubernetes-engine/pricing), [Cloud Build](https://cloud.google.com/build/pricing) and [Artifact Registry](https://cloud.google.com/artifact-registry/pricing).) 

## Prerequisites

Create your [Google Cloud Project](https://cloud.google.com/).

Install tools in your dev environment:

- [gcloud](https://cloud.google.com/sdk/gcloud)
- [docker](https://www.docker.com/)
- [kubectl](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl#install_kubectl)
- [helm](https://helm.sh/)
- [envsubst](https://linux.die.net/man/1/envsubst)
- [skaffold](https://skaffold.dev/) (Optional)
- [minikube](https://minikube.sigs.k8s.io/docs/start/) (Optional)
- [hyperkit](https://github.com/moby/hyperkit) (Optional)

_[Google Cloud Shell](https://cloud.google.com/shell) has all tools you need._

## Create the Resources and Install Gaming OSS

### Deploy them to Google Cloud

```bash
# Set Your Project ID before you run
$ export PROJECT_ID=<your project ID>

$ export LOCATION=us-central1
$ export ZONE=$LOCATION-a

$ export REPOSITORY=space-agon

$ gcloud services enable artifactregistry.googleapis.com \
                        container.googleapis.com

$ gcloud config set project $PROJECT_ID

$ gcloud config set compute/zone $ZONE

# Create cluster (using default network)
# Set NETWORK=<your network>, if you want to select the network
$ make gcloud-test-cluster

# Create Artifact Registry Repository
$ gcloud artifacts repositories create $REPOSITORY \
    --repository-format=docker \
    --location=$LOCATION 

# Assign roles to default service account
$ gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:$(gcloud iam service-accounts list \
    --filter="displayName:Compute Engine default service account" \
    --format="value(email)") \
    --role roles/artifactregistry.reader

# Login Artifact Registry
$ gcloud auth configure-docker $LOCATION-docker.pkg.dev

# Install Agones
$ make agones-install

# Install Open Match
$ make openmatch-install
```

### Deploy them to local k8s cluster by minikube

```bash
# Start minikube
# ref: https://minikube.sigs.k8s.io/docs/commands/start/
$ minikube start --cpus="2" --memory="4096" --kubernetes-version=v1.23.14 --driv
er=hyperkit

# Install minimized Agones
$ make agones-install-local

# Install minimized Open Match
$ make openmatch-install-local
```

## Deploy applications

### Deploy them to Google Cloud

Make sure you installed docker to build and push images

```bash
# Build space-agon images
make build

# Apply space-agon images
make install
```

### Deploy them to local k8s cluster by minikube

```bash
# Build space-agon images for minikube cluster
make build-local

# Apply space-agon images for minikube cluster
make install
```

## View and Play

Get External IP from:

```bash
$ kubectl get service frontend
```

When you run space-agone in minikube, you should followings in another terminal:

```bash
$ minikube tunnel
```

Open `http://<external ip>/` in your favorite web browser.  You can use "find
match" to start searching for a match.

Repeat in a second web browser window to create a second player, the players
will be connected and can play each other.

## Access GameServer

View Running Game Servers:

```bash
$ kubectl get gameserver
```

Then use the connect to server option with the value `<ip>:<port>`.

## Clean Up

### Delete the deployment

```bash
$ make uninstall
```

### Uninstall Agones

```bash
# Deployed space-agone to Google Cloud
$ make agones-uninstall

# Deployed space-agone to minikube
$ make agones-uninstall-local
```

### Uninstall Open-Match

```bash
# Deployed space-agone to Google Cloud
make openmatch-uninstall

# Deployed space-agone to minikube
$ make openmatch-uninstall-local
```

### Delete your Google Cloud Project

```bash
$ gcloud projects delete $PROJECT_ID
```

## Develop Applications

In case testing your original match making logics, [`skaffold`](https://skaffold.dev/) can help you debug your applications. 

### Setup

1. [Create a space-agon k8s cluster.](#create-the-resources-and-install-gaming-oss)
1. [Install `skaffold`](https://skaffold.dev/docs/install/) if you haven't. 
1. Run `make skaffold-setup` on the project root to make a `skaffold.yaml`

Now you're ready to run `skaffold` commands.

### Debug

Once you create a `skaffold.yaml`, you can run `skaffold` commands.

You can check your own logic and debug.

```bash
# Build space-agon images with Cloud Build
$ skaffold build 

# Run Applicaitons in the space-agon cluster for debugging.
$ skaffold dev
```

Modifying applications during `skaffold dev` triggers Build and Deploy automatically. For more commands and details, visit [`skaffold`](https://skaffold.dev/). 

### Test the Cluster

When you would like to test the application, follow the steps below.

#### Google Cloud

1. [Install `skaffold`](https://skaffold.dev/docs/install/) if you haven't. 
1. [Create a space-agon k8s cluster.](#create-the-resources-and-install-gaming-oss)
1. Run `make skaffold-setup` on the project root to make a `skaffold.yaml`
1. Run below commands for integration test. 

```bash
# Run you space-agon applications
$ skaffold dev

# Open another terminal and
# Run Test command
$ make integration-test
```

#### minikube

1. [Create a space-agon k8s cluster via minikube.](#create-the-resources-and-install-gaming-oss)

```bash
# Connect to service in minikube
$ minikube tunnel

# Open another terminal and
# Run Test command
$ make integration-test
```

## LICENSE

This is [Apache 2.0 License](./LICENSE). 

## Note

This is not an officially supported Google product.