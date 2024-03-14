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

GINKGO := $(shell go env GOPATH)/bin/ginkgo
CLI ?= $(OUTPUT_DIR)/commitizen

.PHONY: test.cover
test.cover:
	@echo "===========> Run unit test"
	@go test -race -cover -coverprofile=$(REPO_ROOT)/coverage.out \
			-timeout=10m -short -v ./...

.PHONY: test.e2e
test.e2e: tools.verify.ginkgo
	@echo "===========> Run e2e test, CLI: $(CLI)"
	@$(GINKGO) -v $(REPO_ROOT)/test/e2e -- -cli=$(CLI) -no-tty=$(NO_TTY)
