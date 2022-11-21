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


CONTAINER := run-mockgen

all: docker gen

gen:
	$(eval UID := $(shell stat --print "%u" .))
	$(eval GID := $(shell stat --print "%g" .))
	docker run --rm -it -v `pwd`:/working $(CONTAINER) sh -c 'git clone https://github.com/googleforgames/open-match.git /tmp/open-match \
	&& cd /tmp/open-match/pkg/pb \
	&& mockgen -source /tmp/open-match/pkg/pb/frontend.pb.go -package mockpb -destination /working/mock_frontend_pb.go \
	&& mockgen -source /tmp/open-match/pkg/pb/backend.pb.go -package mockpb -destination /working/mock_backend_pb.go \
	&& chown $(UID):$(GID) /working/mock_backend_pb.go /working/mock_frontend_pb.go'

docker:
	docker build -t $(CONTAINER) -f Dockerfile-mockgen .
