# gsfaas/Makefile

NAME = gsfaas
ENTRY = ./cmd/$(NAME)/main.go
BINARY = ../${BIN_DIR}/$(NAME)
PATH := $(PATH):$(HOME)/go/bin

.PHONY: build test lint fmt imports

build:
	@echo "-- building $(NAME)"
	go build -o $(BINARY) ./...

test:
	@echo "-- testing $(NAME)"
	go test ./... -count=1 -v

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

imports:
	goimports -w .
