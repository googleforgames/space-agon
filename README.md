The original work is [Laremere/space-agon](https://github.com/Laremere/space-agon).

# Space Agon

Space Agon is a integrated demo of [Agones](https://agones.dev/) and
[Open Match](https://open-match.dev/).

## Before Trying.

**Be aware of billing charges for running the cluster.**

Space Agon is intended to run on [Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine) and has been tested with the configured cluster size .   
Leaving the cluster running may incur your cost. You need to take responsibility for it.  (See pricings of [GKE](https://cloud.google.com/kubernetes-engine/pricing), [Cloud Build](https://cloud.google.com/build/pricing) and [Artifact Registry](https://cloud.google.com/artifact-registry/pricing).) 

## Prerequisites

Create your [Google Cloud Project](https://cloud.google.com/) to test and deploy.

You need to install tools in your dev environment:

- [gcloud](https://cloud.google.com/sdk/gcloud)
- [docker](https://www.docker.com/)
- [kubectl](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl#install_kubectl)
- [skaffold](https://skaffold.dev/) (Optional)

_[Google Cloud Shell](https://cloud.google.com/shell) has all tools you need._

## Create the Resources and Install Gaming OSS

```bash
# Set Your Project ID before you run
PROJECT_ID=<your project ID>

LOCATION=us-central1
ZONE=$LOCATION-a

REPOSITORY=space-agon

gcloud services enable artifactregistry.googleapis.com \
                        container.googleapis.com

gcloud config set project $PROJECT_ID

gcloud config set compute/zone $ZONE

# Create cluster (using default network)
# Set NETWORK=<your network>, if you want to select the network
make gcloud-test-cluster

# Create Artifact Registry Repository
gcloud artifacts repositories create $REPOSITORY \
    --repository-format=docker \
    --location=$LOCATION 

# Assign roles to default service account
gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:$(gcloud iam service-accounts list \
    --filter="displayName:Compute Engine default service account" \
    --format="value(email)") \
    --role roles/artifactregistry.reader

# Login Artifact Registry
gcloud auth configure-docker $LOCATION-docker.pkg.dev

# Install Agones
make agones-install

# Install Open Match
make openmatch-install
```

## Commands to deploy

Make sure you installed docker to build and push images

```bash
# Build space-agon images
make build

# Apply space-agon images
make install
```

## View and Play

Get External IP from:

```bash
kubectl get service frontend
```

Open `http://<external ip>/` in your favorite web browser.  You can use "find
match" to start searching for a match.

Repeat in a second web browser window to create a second player, the players
will be connected and can play each other.

## Testing Space-Agon

View Running Game Servers:

```sh
kubectl get gameserver
```
Then use the connect to server option with the value `<ip>:<port>`.

## Clean Up

### Delete the deployment

```sh
make uninstall
```

### Uninstall Agones

```sh
make agones-uninstall
```

### Uninstall Open-Match

```sh
make openmatch-uninstall
```

### Delete your Google Cloud Project

```sh
gcloud projects delete $PROJECT_ID
```

## Development

In case testing your original match making logics, [`skaffold`](https://skaffold.dev/) can help you to debug in the space-agon cluster. 

### Setup

1. Create a space-agon k8s cluster.
1. [Install `skaffold`](https://skaffold.dev/docs/install/) if you haven't installed. 
1. Run `make skaffold-setup` on the project root to create a `skaffold.yaml`

Now you're ready to run `skaffold` commands.

### Debug

Once you create a `skaffold.yaml`, you can run `skaffold` commands.

```bash
# Build space-agon images with Cloud Build
skaffold build 

# Run Applicaitons in the space-agon cluster for debugging.
skaffold dev
```

Changing files triggers Build and Deploy on the fly. For more commands and details, visit [`skaffold`](https://skaffold.dev/). 

## LICENSE

This is [Apache 2.0 License](./LICENSE). 

## Note

This is not an officially supported Google product.