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

#  __     __         _       _     _
#  \ \   / /_ _ _ __(_) __ _| |__ | | ___ ___
#   \ \ / / _` | '__| |/ _` | '_ \| |/ _ \ __|
#    \ V / (_| | |  | | (_| | |_) | |  __\__ \
#     \_/ \__,_|_|  |_|\__,_|_.__/|_|\___|___/
#

PROJECT=$(shell gcloud config list --format 'value(core.project)')
LOCATION=us-central1
REPOSITORY=space-agon
REGISTRY=${LOCATION}-docker.pkg.dev/${PROJECT}/${REPOSITORY}

#   _____                    _
#  |_   _|_ _ _ __ __ _  ___| |_ ___
#    | |/ _` | '__/ _` |/ _ \ __/ __|
#    | | (_| | | | (_| |  __/ |_\__ \
#    |_|\__,_|_|  \__, |\___|\__|___/
#                 |___/

# help output
.PHONY: help
help:
	@echo ""
	@echo "Build Docker images"
	@echo "    make build"
	@echo ""
	@echo "Create GKE Cluster"
	@echo "    make gcloud-test-cluster"
	@echo ""
	@echo "Install Agones"
	@echo "    make agones-install"
	@echo ""
	@echo "Install Open Match"
	@echo "    make openmatch-install"
	@echo ""
	@echo "Install Space Agon"
	@echo "    make install"
	@echo ""
	@echo "Uninstall Agones"
	@echo "    make agones-uninstall"
	@echo ""
	@echo "Uninstall Open Match"
	@echo "    make openmatch-uninstall"
	@echo ""
	@echo "Uninstall Space Agon"
	@echo "    make uninstall"
	@echo ""
	@echo "Setup a Skaffold file for debugging !!RUN AFTER CREATING YOUR CLUSTER!!"
	@echo "    make skaffold-setup"
	@echo ""
	@echo "Run E2E test. "
	@echo "    make e2e-test"
	@echo ""

# build space-agon docker images
.PHONY: build
build:
	./scripts/build.sh ${REGISTRY}

# create gke cluster
.PHONY: gcloud-test-cluster
gcloud-test-cluster: GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT ?= 4
gcloud-test-cluster: GCP_CLUSTER_NODEPOOL_MACHINETYPE ?= e2-standard-4
gcloud-test-cluster: NETWORK ?= default
gcloud-test-cluster:
	./scripts/create-cluster.sh ${GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT} ${GCP_CLUSTER_NODEPOOL_MACHINETYPE} ${LOCATION} ${NETWORK}

# install agones
.PHONY: agones-install
agones-install:
	kubectl create namespace agones-system
	kubectl apply -f https://raw.githubusercontent.com/googleforgames/agones/release-1.23.0/install/yaml/install.yaml

# uninstall agones and agones resources
.PHONY: agones-uninstall
agones-uninstall:
	kubectl delete fleets --all --all-namespaces
	kubectl delete gameservers --all --all-namespaces
	kubectl delete -f https://raw.githubusercontent.com/googleforgames/agones/release-1.23.0/install/yaml/install.yaml
	kubectl delete namespace agones-system

# install open-match
.PHONY: openmatch-install
openmatch-install:
	kubectl create namespace open-match
	kubectl apply -f https://open-match.dev/install/v1.4.0/yaml/01-open-match-core.yaml \
		-f https://open-match.dev/install/v1.4.0/yaml/06-open-match-override-configmap.yaml \
		-f https://open-match.dev/install/v1.4.0/yaml/07-open-match-default-evaluator.yaml \
		--namespace open-match

# uninstall open-match
.PHONY: openmatch-uninstall
openmatch-uninstall:
	kubectl delete psp,clusterrole,clusterrolebinding --selector=release=open-match
	kubectl delete namespace open-match

.PHONY: skaffold-setup
skaffold-setup:
	./scripts/setup-skaffold.sh ${PROJECT} ${REGISTRY}

# install space-agon itself
.PHONY: install
install:
	kubectl apply -f deploy.yaml

# uninstall space-agon itself
.PHONY: uninstall
uninstall:
	kubectl delete -f deploy.yaml

# TODO: ADD E2e test script
.PHONY: integration-test
integration-test:
	go test -count=1 -v -timeout 60s test/integration_test.go

