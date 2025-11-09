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

lint:
	@for m in $(MODULES); do \
		echo "- lint $$m"; \
		$(MAKE) -C $$m lint || exit 1; \
	done

fmt:
	@for m in $(MODULES); do \
		echo "- go fmt $$m"; \
		$(MAKE) -C $$m fmt; \
	done

imports:
	@for m in $(MODULES); do \
		echo "- goimports $$m"; \
		$(MAKE) -C $$m imports; \
	done
