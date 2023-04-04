#!/bin/bash

# Copyright 2022 Google LLC All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

export TAG=$(git rev-parse --short HEAD)

if [ $1 = "test" ] ;then
    export ENV="test"
    export REGISTRY=local
else
    export ENV="develop"
    export REGISTRY=$1
fi

# Use docker engine in minikube
if [ ${ENV} = "test" ] ;then
    eval $(minikube -p minikube docker-env)
fi

# Build images
docker build -f ./Frontend.Dockerfile -t ${REGISTRY}/space-agon-frontend:${TAG} .
docker build -f ./Dedicated.Dockerfile -t ${REGISTRY}/space-agon-dedicated:${TAG} .
docker build -f ./Director.Dockerfile -t ${REGISTRY}/space-agon-director:${TAG} .
docker build -f ./Mmf.Dockerfile -t ${REGISTRY}/space-agon-mmf:${TAG} .

# Push images
if [ ${ENV} = "develop" ];then
    docker push ${REGISTRY}/space-agon-frontend:${TAG}
    docker push ${REGISTRY}/space-agon-dedicated:${TAG}
    docker push ${REGISTRY}/space-agon-director:${TAG}
    docker push ${REGISTRY}/space-agon-mmf:${TAG}
fi

# Replace image repository & tags
if [ ${ENV} = "test" ] ;then
    REGISTRY=${REGISTRY} 
    TAG=${TAG} 
    export REPLICAS_DEDICATED=1 
    export REPLICAS_FRONTEND=1 
    export REPLICAS_MMF=1 
    export REQUEST_MEMORY=100Mi 
    export REQUEST_CPU=100m
    export LIMITS_MEMORY=100Mi 
    export LIMITS_CPU=100m 
    export BUFFER_SIZE=1 
    export MIN_REPLICAS=0 
    export MAX_REPLICAS=1 
	envsubst < deploy_template.yaml > deploy.yaml
	envsubst < templates/helm_template_values.yaml > install/helm/space-agon/values.yaml
	envsubst < templates/skaffold_template.local.yaml > skaffold.yaml
else
    REGISTRY=${REGISTRY} 
    TAG=${TAG} 
    export REPLICAS_DEDICATED=2 
    export REPLICAS_FRONTEND=2 
    export REPLICAS_MMF=2 
    export REQUEST_MEMORY=200Mi 
    export REQUEST_CPU=500m 
    export LIMITS_MEMORY=200Mi 
    export LIMITS_CPU=500m 
    export BUFFER_SIZE=2 
    export MIN_REPLICAS=0 
    export MAX_REPLICAS=50 
	envsubst < deploy_template.yaml > deploy.yaml
	envsubst < templates/helm_template_values.yaml > install/helm/space-agon/values.yaml
	envsubst < templates/skaffold_template.yaml > skaffold.yaml
fi