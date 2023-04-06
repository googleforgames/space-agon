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

export ENV=$1
export TAG=$2
export FRONTEND_IMG=$3
export DEDICATED_IMG=$4
export DIRECTOR_IMG=$5
export MMF_IMG=$6
export REGISTRY=$7
export PROJECT=$8
export LOCATION=$9

# Use docker engine in minikube
if [ ${ENV} = "test" ] ;then
    export REGISTRY=local
    eval $(minikube -p minikube docker-env)
fi

# Build images
docker build -f ./Frontend.Dockerfile -t "${REGISTRY}/${FRONTEND_IMG}:${TAG}" .
docker build -f ./Dedicated.Dockerfile -t "${REGISTRY}/${DEDICATED_IMG}:${TAG}" .
docker build -f ./Director.Dockerfile -t "${REGISTRY}/${DIRECTOR_IMG}:${TAG}" .
docker build -f ./Mmf.Dockerfile -t "${REGISTRY}/${MMF_IMG}:${TAG}" .

# Push images
if [ ${ENV} = "develop" ];then
    docker push ${REGISTRY}/${FRONTEND_IMG}:${TAG}
    docker push ${REGISTRY}/${DEDICATED_IMG}:${TAG}
    docker push ${REGISTRY}/${DIRECTOR_IMG}:${TAG}
    docker push ${REGISTRY}/${MMF_IMG}:${TAG}
fi

export REGEXP="s/[-/.:@]/_/g"
# Sanitized characters - / . : for skaffold
# https://skaffold.dev/docs/deployers/helm/#sanitizing-the-artifact-name-from-invalid-go-template-characters
export SANITIZED_REGISTRY=$(echo ${REGISTRY} | sed -e ${REGEXP})
export SANITIZED_FRONTEND=$(echo ${FRONTEND_IMG} | sed -e ${REGEXP})
export SANITIZED_DEDICATED=$(echo ${DEDICATED_IMG} | sed -e ${REGEXP})
export SANITIZED_DIRECTOR=$(echo ${DIRECTOR_IMG} | sed -e ${REGEXP})
export SANITIZED_MMF=$(echo ${MMF_IMG} | sed -e ${REGEXP})

# Create skaffold.yaml with environments
envsubst < templates/skaffold_template.yaml > skaffold.yaml