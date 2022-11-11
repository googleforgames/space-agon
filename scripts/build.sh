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

TAG=$(git rev-parse --short HEAD)

if [ $1 = "test" ] ;then
    ENV="test"
    REGISTRY=local
else
    ENV="develop"
    REGISTRY=$1
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
    REGISTRY=${REGISTRY} TAG=${TAG} REPLICAS_FRONTEND=1 REPLICAS_DEDICATED=1 REQUEST_MEMORY=100Mi REQUEST_CPU=100m \
    LIMITS_MEMORY=100Mi LIMITS_CPU=100m BUFFER_SIZE=1 MIN_REPLICAS=0 MAX_REPLICAS=1 REPLICAS_MMF=1 \
	envsubst < deploy_template.yaml > deploy.yaml
else
    REGISTRY=${REGISTRY} TAG=${TAG} REPLICAS_FRONTEND=2 REPLICAS_DEDICATED=2 REQUEST_MEMORY=200Mi REQUEST_CPU=500m \
    LIMITS_MEMORY=200Mi LIMITS_CPU=500m BUFFER_SIZE=2 MIN_REPLICAS=0 MAX_REPLICAS=50 REPLICAS_MMF=2 \
	envsubst < deploy_template.yaml > deploy.yaml
fi