# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

MOCKGEN_CONTAINER := run-mockgen
MOCKGEN_BASEDIR := pkg/testing/open-match/mockpb

.PHONY: open-match-mockgen-all
open-match-mockgen-all: docker-for-mockgen open-match-mockgen

.PHONY: open-match-mockgen
open-match-mockgen:
	$(eval UID := $(shell stat --print "%u" .))
	$(eval GID := $(shell stat --print "%g" .))
	docker run --rm -it -v $(CURDIR):/working $(MOCKGEN_CONTAINER) sh -c 'git clone https://github.com/googleforgames/open-match.git /open-match \
	&& cd /working \
	&& mockgen -source /open-match/pkg/pb/frontend.pb.go -package mockpb -destination /working/$(MOCKGEN_BASEDIR)/mock_frontend_pb.go -copyright_file /LICENSE.txt \
	&& mockgen -source /open-match/pkg/pb/backend.pb.go -package mockpb -destination /working/$(MOCKGEN_BASEDIR)/mock_backend_pb.go -copyright_file /LICENSE.txt \
	&& chown $(UID):$(GID) /working/$(MOCKGEN_BASEDIR)/mock_backend_pb.go /working/$(MOCKGEN_BASEDIR)/mock_frontend_pb.go'

.PHONY: docker-for-mockgen
docker-for-mockgen:
	docker build -t $(MOCKGEN_CONTAINER) -f $(CURDIR)/$(MOCKGEN_BASEDIR)/Dockerfile-mockgen $(CURDIR)/$(MOCKGEN_BASEDIR)

