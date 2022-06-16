The original work is [Laremere/space-agon](https://github.com/Laremere/space-agon).

# Space Agon

Space Agon is a demo of [Agones](https://agones.dev/) and
[Open Match](https://open-match.dev/). You can try integrations of Gaming OSS.

## Before Installing

**Warning**: Be aware of billing charges for running the cluster.  

Space Agon has been tested on this cluster size (nodes and machine types), but a small cluster may be sufficient for your use.    
Don't leave the cluster running when you're not using it if you're concerned about cost. See [pricing](https://cloud.google.com/kubernetes-engine/pricing) for more.

## Prerequisites

You need to install tools:

- [docker](https://www.docker.com/)
- [gcloud](https://cloud.google.com/sdk/gcloud)
- [kubectl](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl#install_kubectl)

[Google Cloud Shell](https://cloud.google.com/shell) has all tools you need.

## Create the Resources and Install Gaming OSS

```sh
LOCATION=us-central1
ZONE=$LOCATION-a
# Set Your Project ID before you run
PROJECT_ID=<your project id>
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

```sh
# build space-agon images
make build

# apply space-agon images
make install
```

## View and Play

Get External IP from:

```sh
kubectl get service frontend
```

Open `http://<external ip>/` in your favorite web browser.  You can use "find
match" to start searching for a match.

Repeat in a second web browser window to create a second player, the players
will be connected and can play each other.

## Additional Things to do

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

## LICENSE

This is [Apache 2.0 License](./LICENSE). 

## Note

This is not an officially supported Google product.