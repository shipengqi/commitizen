TOOLS ?=$(BUILD_TOOLS)

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.golangci-lint
install.golangci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: install.gsemver
install.gsemver:
	@go install github.com/arnaud-deprez/gsemver@latest

.PHONY: install.releaser
install.releaser:
	@go install github.com/goreleaser/goreleaser@latest

.PHONY: install.ginkgo
install.ginkgo:
	@go install github.com/onsi/ginkgo/v2/ginkgo@latest