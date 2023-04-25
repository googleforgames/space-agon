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

# DO NOT RUN this script directly.
# Run `make skaffold-setup` command after creating the cluster.
# This script setup skaffold.yaml

export PROJECT_ID=$1
export REGISTRY=$2
export FRONTEND_IMG=$3
export DEDICATED_IMG=$4
export DIRECTOR_IMG=$5
export MMF_IMG=$6
export LOCATION=$7

gcloud config set project ${PROJECT_ID}
gcloud services enable cloudbuild.googleapis.com

gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member=serviceAccount:$(gcloud projects describe ${PROJECT_ID} \
    --format="value(projectNumber)")@cloudbuild.gserviceaccount.com \
    --role="roles/storage.objectViewer"

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