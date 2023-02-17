These are some additional commands which are useful for development.

Generate components file:
```
go generate github.com/mbychkowski/space-agon/game/generation
```

Rebuild proto generates files:
```
docker build -f=Dockerfile.build-protos -t build-protos . && docker run --rm --mount type=bind,source="$(pwd)",target=/workdir/mount build-protos
```

Mother of all commands:
```

# Once
docker build -f=Dockerfile.build-protos -t build-protos .

# Every update

go generate github.com/mbychkowski/space-agon/game/generation && \
docker run --rm --mount type=bind,source="$(pwd)",target=/workdir/mount build-protos && \
TAG=$(date +INDEV-%Y%m%d-%H%M%S) && \
REGISTRY=gcr.io/$(gcloud config list --format 'value(core.project)') && \
GOOS=js GOARCH=wasm go test github.com/mbychkowski/space-agon/client/... && \
go test github.com/mbychkowski/space-agon/...





 && \
docker build . -f Frontend.Dockerfile -t $REGISTRY/space-agon-frontend:$TAG && \
docker build . -f Dedicated.Dockerfile -t $REGISTRY/space-agon-dedicated:$TAG && \
docker build . -f Director.Dockerfile -t $REGISTRY/space-agon-director:$TAG && \
docker build . -f Mmf.Dockerfile -t $REGISTRY/space-agon-mmf:$TAG && \
docker push $REGISTRY/space-agon-frontend:$TAG && \
docker push $REGISTRY/space-agon-dedicated:$TAG && \
docker push $REGISTRY/space-agon-director:$TAG && \
docker push $REGISTRY/space-agon-mmf:$TAG && \
ESC_REGISTRY=$(echo $REGISTRY | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
ESC_TAG=$(echo $TAG | sed -e 's/\\/\\\\/g; s/\//\\\//g; s/&/\\\&/g') && \
sed -E 's/image: (.*)\/([^\/]*):(.*)/image: '$ESC_REGISTRY'\/\2:'$ESC_TAG'/' deploy_template.yaml > deploy.yaml && \
kubectl apply -f deploy.yaml
```

Run locally
```
# Gameserver - RUN FIRST

go generate github.com/mbychkowski/space-agon/game/generation && \
docker run --rm --mount type=bind,source="$(pwd)",target=/workdir/mount build-protos && \
go test github.com/mbychkowski/space-agon/... && \
docker build . -f Dedicated.Dockerfile -t space-agon-dedicated && \
docker run -p 2156:2156/tcp -e DISABLE_AGONES=true space-agon-dedicated

# Frontend - RUN SECOND

GOOS=js GOARCH=wasm go test github.com/mbychkowski/space-agon/client/... && \
go test github.com/mbychkowski/space-agon/... && \
docker build . -f Frontend.Dockerfile -t space-agon-frontend && \
docker run -p 2157:8080/tcp space-agon-frontend

```


This is not an officially supported Google product.
