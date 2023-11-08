# Directory containing the Makefile.
PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Directories containing independent Go modules.
MODULE_DIRS = . ./src

.PHONY: all
all: lint

.PHONY: lint
lint: golangci-lint tidy-lint

.PHONY: golangci-lint
golangci-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] golangci-lint: $(mod)" && \
		golangci-lint run --path-prefix $(mod)) &&) true

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS), \
		(cd $(dir) && go mod tidy) &&) true

.PHONY: tidy-lint
tidy-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] tidy: $(mod)" && \
		go mod tidy && \
		git diff --exit-code -- go.mod go.sum) &&) true
