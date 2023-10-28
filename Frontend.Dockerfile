# Copyright 2019 Google LLC
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

FROM golang:1.21.3 as builder
ENV GO111MODULE=on

WORKDIR /go/src/github.com/googleforgames/space-agon

COPY go.sum go.mod ./
RUN go mod download

RUN mkdir /app
COPY frontend ./frontend
COPY game ./game
COPY client ./client
COPY static /app/static
RUN CGO_ENABLED=0 go build -installsuffix cgo -o /app/frontend github.com/googleforgames/space-agon/frontend
RUN GOOS=js GOARCH=wasm go build -o /app/static/client.wasm github.com/googleforgames/space-agon/client
RUN cp /usr/local/go/misc/wasm/wasm_exec.js /app/static/

FROM gcr.io/distroless/static:nonroot
COPY --from=builder --chown=nonroot "/app" "/app"
ENTRYPOINT ["/app/frontend"]
