# Root Makefile

MODULES = core gsfaas gsfcmpl
export BIN_DIR = bin

.PHONY: all build test cor lint fmt imports $(MODULES)

all: build

build: $(MODULES)

$(MODULES):
	@echo "- building module $@"
	$(MAKE) -C $@ build

test:
	@for m in $(MODULES); do \
		echo "- testing $$m"; \
		$(MAKE) -C $$m test || exit 1; \
	done

cor: lint fmt imports

CHECK_LINTER = command -v golangci-lint >/dev/null 2>&1
CHECK_IMPORTS = command -v goimports >/dev/null 2>&1
export PATH := $(PATH):$(HOME)/go/bin

lint-tools:
	@if ! $(CHECK_LINTER); then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
			GOLANGCI_LINT_VERSION=v1.62.2 sh -s -- -b $$HOME/go/bin; \
	fi

imports-tools:
	@if ! $(CHECK_IMPORTS); then \
		go install golang.org/x/tools/cmd/goimports@latest; \
	fi


lint: lint-tools
	@for m in $(MODULES); do \
		echo "- lint $$m"; \
		$(MAKE) -C $$m lint || exit 1; \
	done

fmt:
	@for m in $(MODULES); do \
		echo "- go fmt $$m"; \
		$(MAKE) -C $$m fmt; \
	done

imports: imports-tools
	@for m in $(MODULES); do \
		echo "- goimports $$m"; \
		$(MAKE) -C $$m imports; \
	done
