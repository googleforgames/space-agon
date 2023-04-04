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

AGONES_NS:=agones-system
OM_NS:=open-match
AGONES_VER:=1.29.0
OM_VER:=1.7.0

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
	@echo "Build Docker images in local environment"
	@echo "    make build-local"
	@echo ""
	@echo "Build Docker images"
	@echo "    make build"
	@echo ""
	@echo "Create GKE Cluster"
	@echo "    make gcloud-test-cluster"
	@echo ""
	@echo "Add Helm Repositories"
	@echo "    make helm-repo-add"
	@echo ""
	@echo "Remove Helm Repositories"
	@echo "    make helm-repo-remove"
	@echo ""
	@echo "Install Agones in local-cluster"
	@echo "    make agones-install-local"
	@echo ""
	@echo "Install Agones"
	@echo "    make agones-install"
	@echo ""
	@echo "Install Open Match in local-cluster"
	@echo "    make openmatch-install-local"
	@echo ""
	@echo "Install Open Match"
	@echo "    make openmatch-install"
	@echo ""
	@echo "Install Space Agon"
	@echo "    make install"
	@echo ""
	@echo "Uninstall Agones in local-cluster"
	@echo "    make agones-uninstall-local"
	@echo ""
	@echo "Uninstall Agones"
	@echo "    make agones-uninstall"
	@echo ""
	@echo "Uninstall Open Match in local-cluster"
	@echo "    make openmatch-uninstall-local"
	@echo ""
	@echo "Uninstall Open Match"
	@echo "    make openmatch-uninstall"
	@echo ""
	@echo "Upgrade Space Agon parameters"
	@echo "    make upgrade"
	@echo ""
	@echo "Uninstall Space Agon"
	@echo "    make uninstall"
	@echo ""
	@echo "Setup Cloud Build for building your image remotely"
	@echo "    make cloudbuild-setup"
	@echo ""
	@echo "Run integration test"
	@echo "    make integration-test"
	@echo ""

# build space-agon docker images in local
.PHONY: build-local
build-local:
	./scripts/build.sh test

# build space-agon docker images
.PHONY: build
build:
	./scripts/build.sh ${REGISTRY} ${PROJECT} ${LOCATION}

# create gke cluster
.PHONY: gcloud-test-cluster
gcloud-test-cluster: GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT ?= 4
gcloud-test-cluster: GCP_CLUSTER_NODEPOOL_MACHINETYPE ?= e2-standard-4
gcloud-test-cluster: NETWORK ?= default
gcloud-test-cluster:
	./scripts/create-cluster.sh ${GCP_CLUSTER_NODEPOOL_INITIALNODECOUNT} ${GCP_CLUSTER_NODEPOOL_MACHINETYPE} ${LOCATION} ${NETWORK}

.PHONY: helm-repo-add
helm-repo-add: 
	helm repo add $(AGONES_NS) https://agones.dev/chart/stable
	helm repo add $(OM_NS) https://open-match.dev/chart/stable
	helm repo update

.PHONY: helm-repo-remove
helm-repo-remove: 
	helm repo remove $(AGONES_NS)
	helm repo remove $(OM_NS) 

# install agones in local-cluster
.PHONY: agones-install-local
agones-install-local:
	helm install $(AGONES_NS) --namespace $(AGONES_NS) \
		--create-namespace $(AGONES_NS)/agones \
		--version $(AGONES_VER) \
		--set agones.ping.install=false \
		--set agones.allocator.replicas="1"

# install agones
.PHONY: agones-install
agones-install:
	helm install ${AGONES_NS} --namespace ${AGONES_NS} \
		--create-namespace $(AGONES_NS)/agones \
		--version ${AGONES_VER}

# uninstall agones and agones resources in local-cluster
.PHONY: agones-uninstall-local
agones-uninstall-local:
	helm uninstall $(AGONES_NS) --namespace $(AGONES_NS)
	kubectl delete namespace $(AGONES_NS)

# uninstall agones and agones resources
.PHONY: agones-uninstall
agones-uninstall:
	helm uninstall $(AGONES_NS) --namespace $(AGONES_NS)
	kubectl delete namespace $(AGONES_NS)

# install open-match in local-cluster
.PHONY: openmatch-install-local
openmatch-install-local:
	helm install $(OM_NS) \
	--create-namespace --namespace $(OM_NS) $(OM_NS)/open-match \
	--version $(OM_VER) \
	--set open-match-customize.enabled=true \
	--set open-match-customize.evaluator.enabled=true \
	--set open-match-customize.evaluator.replicas=1 \
	--set open-match-override.enabled=true \
	--set open-match-core.swaggerui.enabled=false \
	--set global.kubernetes.horizontalPodAutoScaler.frontend.maxReplicas=1 \
	--set global.kubernetes.horizontalPodAutoScaler.backend.maxReplicas=1 \
	--set global.kubernetes.horizontalPodAutoScaler.query.minReplicas=1 \
	--set global.kubernetes.horizontalPodAutoScaler.query.maxReplicas=1 \
	--set global.kubernetes.horizontalPodAutoScaler.evaluator.maxReplicas=1 \
	--set query.replicas=1 \
	--set frontend.replicas=1 \
	--set backend.replicas=1 \
	--set redis.sentinel.enabled=false \
	--set redis.master.resources.requests.cpu=0.1 \
	--set redis.master.persistence.enabled=false \
	--set redis.replica.replicaCount=0 \
	--set redis.metrics.enabled=false

# install open-match
.PHONY: openmatch-install
openmatch-install:
	helm install ${OM_NS} --create-namespace --namespace ${OM_NS} $(OM_NS)/open-match --version ${OM_VER} \
	--set open-match-customize.enabled=true \
	--set open-match-customize.evaluator.enabled=true \
	--set open-match-customize.evaluator.replicas=1 \
	--set open-match-override.enabled=true \
	--set open-match-core.swaggerui.enabled=false \
	--set redis.sentinel.enabled=false \
	--set redis.master.resources.requests.cpu=0.1 \
	--set redis.master.persistence.enabled=false \
	--set redis.replica.replicaCount=0 \
	--set redis.metrics.enabled=false

# uninstall open-match in local-cluster
.PHONY: openmatch-uninstall-local
openmatch-uninstall-local:
	helm uninstall -n $(OM_NS) $(OM_NS)
	kubectl delete namespace $(OM_NS)

# uninstall open-match
.PHONY: openmatch-uninstall
openmatch-uninstall:
	helm uninstall -n ${OM_NS} ${OM_NS}
	kubectl delete namespace ${OM_NS}

.PHONY: cloudbuild-setup
cloudbuild-setup:
	./scripts/setup-cloudbuild.sh ${PROJECT} ${REGISTRY}

# install space-agon itself
.PHONY: install
install:
	helm install space-agon -f install/helm/space-agon/values.yaml ./install/helm/space-agon

# uninstall space-agon itself
.PHONY: uninstall
uninstall:
	helm uninstall space-agon 

# upgrade space-agon after changing parameters
.PHONY: upgrade
upgrade:
	helm upgrade space-agon -f install/helm/space-agon/values.yaml ./install/helm/space-agon

# integration test
.PHONY: integration-test
integration-test:
	go test -count=1 -v -timeout 60s test/integration_test.go
