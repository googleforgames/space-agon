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

set -e

echo "BUILDING PROTO FILES"

mkdir output

find mount -name '*.proto' -exec echo {} \;

find mount -name '*.proto' \
  -exec protoc {} \
    -Imount/ \
    -I/usr/include/ \
    --go_out=plugins=grpc:output \;

cp -r -f output/github.com/mbychkowski/space-agon/* mount/
