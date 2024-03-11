#!/bin/bash

# Copyright (c) 2022 PengQi Shi
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

set -o errexit
set -o nounset

if [[ -z "${GO_LDFLAGS}" ]]; then
    echo "GO_LDFLAGS must be set"
    exit 1
fi

# set .Env.GO_LDFLAGS for goreleaser.yaml
export GO_LDFLAGS=${GO_LDFLAGS}

RELEASER=$(go env GOPATH)/bin/goreleaser

# $PUBLISH must explicitly be set to '1' for goreleaser
# to publish the release to GitHub.
if [[ "${PUBLISH:-}" != "1" ]]; then
    echo "Not set to publish"
    ${RELEASER} release \
        --clean \
        --snapshot \
        --skip-publish
else
    if [[ -z "${GITHUB_TOKEN}" ]]; then
        echo "GITHUB_TOKEN must be set"
        exit 1
    fi

    echo "Getting ready to publish"
    ${RELEASER} release \
    --clean
fi
