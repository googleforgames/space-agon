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
# Set --network <vpc> --subnetwork <subnet> if you want to select the network
gcloud container clusters create space-agon \
    --cluster-version=1.22 \
    --tags=game-server \
    --scopes=gke-default \
    --num-nodes=4 \
    --no-enable-autoupgrade \
    --machine-type=n1-standard-4

# Open Firewall for Agones
gcloud compute firewall-rules create gke-game-server-firewall \
    --allow tcp:7000-8000 \
    --target-tags game-server \
    --description "Firewall to allow game server tcp traffic"

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
kubectl create namespace agones-system
kubectl apply -f https://raw.githubusercontent.com/googleforgames/agones/release-1.23.0/install/yaml/install.yaml

# Install Open Match
kubectl create namespace open-match
kubectl apply -f https://open-match.dev/install/v1.3.0/yaml/01-open-match-core.yaml \
    -f https://open-match.dev/install/v1.3.0/yaml/06-open-match-override-configmap.yaml \
    -f https://open-match.dev/install/v1.3.0/yaml/07-open-match-default-evaluator.yaml \
    --namespace open-match
```

## Commands to deploy

Make sure you installed docker to build and push images

```sh
TAG=$(date +INDEV-%Y%m%d-%H%M%S) && \
REGISTRY=$LOCATION-docker.pkg.dev/$(gcloud config list --format 'value(core.project)')/$REPOSITORY && \
docker build . -f Frontend.Dockerfile -t $REGISTRY/space-agon-frontend:$TAG && \
docker build . -f Dedicated.Dockerfile -t $REGISTRY/space-agon-dedicated:$TAG && \
docker build . -f Director.Dockerfile -t $REGISTRY/space-agon-director:$TAG && \
docker build . -f Mmf.Dockerfile -t $REGISTRY/space-agon-mmf:$TAG && \
docker push $REGISTRY/space-agon-frontend:$TAG && \
docker push $REGISTRY/space-agon-dedicated:$TAG && \
docker push $REGISTRY/space-agon-director:$TAG && \
docker push $REGISTRY/space-agon-mmf:$TAG && \
ESC_REGISTRY=$(echo $REGISTRY | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
ESC_TAG=$(echo $TAG | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
sed -E 's/image: (.*)\/([^\/]*):(.*)/image: '$ESC_REGISTRY'\/\2:'$ESC_TAG'/' deploy_template.yaml > deploy.yaml && \
kubectl apply -f deploy.yaml
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

### Uninstall Agones

```sh
kubectl delete fleets --all --all-namespaces
kubectl delete gameservers --all --all-namespaces
kubectl delete -f https://raw.githubusercontent.com/googleforgames/agones/release-1.23.0/install/yaml/install.yaml
kubectl delete namespace agones-system
```

### Uninstall Open-Match

```sh
kubectl delete psp,clusterrole,clusterrolebinding --selector=release=open-match
kubectl delete namespace open-match
```

### Delete the deployment

```sh
kubectl delete -f deploy.yaml 
```

### Delete your Google Cloud Project

```sh
gcloud projects delete $PROJECT_ID
```

## LICENSE

This is [Apache 2.0 License](./LICENSE). 

## Note

This is not an officially supported Google product.