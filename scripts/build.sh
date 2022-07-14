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

TAG=$(date +INDEV-%Y%m%d-%H%M%S)
REGISTRY=$1

# Build images
docker build -f ./Frontend.Dockerfile -t ${REGISTRY}/space-agon-frontend:${TAG} .
docker build -f ./Dedicated.Dockerfile -t ${REGISTRY}/space-agon-dedicated:${TAG} .
docker build -f ./Director.Dockerfile -t ${REGISTRY}/space-agon-director:${TAG} .
docker build -f ./Mmf.Dockerfile -t ${REGISTRY}/space-agon-mmf:${TAG} .

# Push images
docker push ${REGISTRY}/space-agon-frontend:${TAG}
docker push ${REGISTRY}/space-agon-dedicated:${TAG}
docker push ${REGISTRY}/space-agon-director:${TAG}
docker push ${REGISTRY}/space-agon-mmf:${TAG}

# Replace image repository & tags
ESC_REGISTRY=$(echo ${REGISTRY} | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
ESC_TAG=$(echo ${TAG} | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
sed -E 's/image: (.*)\/([^\/]*):(.*)/image: '${ESC_REGISTRY}'\/\2:'${ESC_TAG}'/' deploy_template.yaml > deploy.yaml
