# core/Makefile

NAME = core
ENTRY = ./cmd/$(NAME)/main.go

.PHONY: build test lint fmt imports

build:
	@echo "-- ignoring $(NAME)"

test:
	@echo "-- testing $(NAME)"
	go test ./... -count=1 -v

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

imports:
	goimports -w .
