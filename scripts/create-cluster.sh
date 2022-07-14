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

GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT=$1
GCP_CLUSTER_NODEPOOL_MACHINETYPE=$2
LOCATION=$3
NETWORK=$4

gcloud container clusters create space-agon \
  --cluster-version=1.22 \
  --tags=game-server \
  --scopes=gke-default \
  --network ${NETWORK} \
  --num-nodes=${GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT} \
  --no-enable-autoupgrade \
  --machine-type=${GCP_CLUSTER_NODEPOOL_MACHINETYPE} \
  --zone ${LOCATION}-a
        
gcloud compute firewall-rules create gke-game-server-firewall \
  --allow tcp:7000-8000 \
  --target-tags game-server \
  --description "Firewall to allow game server tcp traffic"
