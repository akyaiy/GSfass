# gsfaas/Makefile

NAME = gsfaas
ENTRY = ./cmd/$(NAME)/main.go
BIN_DIR = bin
BINARY = ../${BIN_DIR}/$(NAME)
CHECK_LINTER = command -v golangci-lint >/dev/null 2>&1
CHECK_IMPORTS = command -v goimports >/dev/null 2>&1
PATH := $(PATH):$(HOME)/go/bin

lint-tools:
	@if ! $(CHECK_LINTER); then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
			GOLANGCI_LINT_VERSION=v1.62.2 sh -s -- -b $$HOME/go/bin; \
	fi

imports-tools:
	@if ! $(CHECK_IMPORTS); then \
		go install golang.org/x/tools/cmd/goimports@latest; \
	fi

.PHONY: build test lint fmt imports

build:
	@echo "-- building $(NAME)"
	@go build -o $(BINARY) ./...

test:
	@echo "-- testing $(NAME)"
	@go test ./... -count=1 -v

lint: lint-tools
	@golangci-lint run ./...

fmt:
	@go fmt ./...

imports: imports-tools
	@goimports -w .
