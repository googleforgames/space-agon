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
export PREFIX="space-agon"

if [ $1 = "test" ] ;then
    export ENV="test"
    export REGISTRY=local
else
    export ENV="develop"
    export REGISTRY=$1
    export PROJECT=$2
    export LOCATION=$3
fi

# Use docker engine in minikube
if [ ${ENV} = "test" ] ;then
    eval $(minikube -p minikube docker-env)
fi

# Build images
docker build -f ./Frontend.Dockerfile -t ${REGISTRY}/${PREFIX}-frontend:${TAG} .
docker build -f ./Dedicated.Dockerfile -t ${REGISTRY}/${PREFIX}-dedicated:${TAG} .
docker build -f ./Director.Dockerfile -t ${REGISTRY}/${PREFIX}-director:${TAG} .
docker build -f ./Mmf.Dockerfile -t ${REGISTRY}/${PREFIX}-mmf:${TAG} .

# Push images
if [ ${ENV} = "develop" ];then
    docker push ${REGISTRY}/${PREFIX}-frontend:${TAG}
    docker push ${REGISTRY}/${PREFIX}-dedicated:${TAG}
    docker push ${REGISTRY}/${PREFIX}-director:${TAG}
    docker push ${REGISTRY}/${PREFIX}-mmf:${TAG}
fi

# Replace image repository & tags
if [ ${ENV} = "develop" ] ;then
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
	envsubst < templates/helm_template_values.yaml > install/helm/space-agon/values.yaml
fi

export REGEXP="s/[-/.:@]/_/g"
# Sanitized characters - / . : for skaffold
# https://skaffold.dev/docs/deployers/helm/#sanitizing-the-artifact-name-from-invalid-go-template-characters
export SANITIZED_REGISTRY=$(echo ${REGISTRY} | sed -e ${REGEXP})
export SANITIZED_PREFIX=$(echo ${PREFIX} | sed -e ${REGEXP})

# Create skaffold.yaml with environments
envsubst < templates/skaffold_template.yaml > skaffold.yaml