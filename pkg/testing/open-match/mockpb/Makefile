CONTAINER := run-mockgen

all: docker gen

gen:
	$(eval UID := $(shell stat --print "%u" .))
	$(eval GID := $(shell stat --print "%g" .))
	docker run --rm -it -v `pwd`:/working $(CONTAINER) bash -c 'git clone https://github.com/googleforgames/open-match.git /tmp/open-match \
	&& cd /tmp/open-match/pkg/pb \
	&& mockgen -source /tmp/open-match/pkg/pb/frontend.pb.go -package mockpb -destination /working/mock_frontend_pb.go \
	&& mockgen -source /tmp/open-match/pkg/pb/backend.pb.go -package mockpb -destination /working/mock_backend_pb.go \
	&& chown $(UID):$(GID) /working/mock_backend_pb.go /working/mock_frontend_pb.go'

docker:
	docker build -t $(CONTAINER) -f Dockerfile-mockgen .
