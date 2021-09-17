BUILD_IMAGE ?= golang:1.16
GOPATH ?= $(shell go env GOPATH)
GOPATH_SRC := $(GOPATH)/src/
CURRENT_WORK_DIR := $(shell pwd)
ALL_FILES=$(shell find . -path ./vendor -prune -type f -o -name *.proto)
PKG=$(shell git config -l | grep "remote.origin.url" | sed "s/remote.origin.url=git@//" | sed "s/:/\//" | sed "s/\.git//")

GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION ?= $(GIT_COMMIT)

BIN ?= ""
TAG ?= $(VERSION)
IMAGE ?= ""
YES ?= ""

all: build build-linux

clean: clean-dirs

build: clean
	@$(MAKE) $(subst cmd, dist/linux/arm, $(wildcard cmd/*))
	@$(MAKE) $(subst cmd, dist/linux/arm64, $(wildcard cmd/*))
	@$(MAKE) $(subst cmd, dist/linux/amd64, $(wildcard cmd/*))
	@$(MAKE) $(subst cmd, dist/darwin/arm64, $(wildcard cmd/*))
	@$(MAKE) $(subst cmd, dist/darwin/amd64, $(wildcard cmd/*))

dist/%: build-dirs
	$(eval OS := $(shell echo "$*" | cut -d'/' -f1))
	$(eval ARCH := $(shell echo "$*" | cut -d'/' -f2))
	$(eval BINARY_NAME := $(shell echo "$*" | cut -d'/' -f3))
	$(info building binary $(notdir $@) for OS $(OS) and ARCH $(ARCH))

	@mkdir -p dist/$(OS)/$(ARCH)
	@mkdir -p .go/std/$(ARCH)

	@docker run \
		--rm \
		-u $$(id -u):$$(id -g) \
		-v "$$(pwd):/src" \
		-v "$$(pwd)/dist/$(OS)/$(ARCH):/go/bin" \
		-v "$$(pwd)/.gocache/:/go/cache" \
		-w /src \
		$(BUILD_IMAGE) \
		/bin/sh -c " \
			ARCH=$(ARCH) \
			OS=$(OS) \
			VERSION=$(VERSION) \
			COMMIT=$(GIT_COMMIT) \
			PKG=$(PKG) \
			BIN=$(notdir $@) \
			GO111MODULE=on \
			./hack/build.sh \
		"

test: build-dirs
	$(info run test)
	@docker run \
		--rm \
		-u $$(id -u):$$(id -g) \
		-v "$$(pwd):/src" \
		-v "$$(pwd)/dist/linux/amd64:/go/bin" \
		-v "$$(pwd)/.gocache/:/go/cache" \
		-w /src \
		$(BUILD_IMAGE) \
		/bin/sh -c "CGO_ENABLED=1 GO111MODULE=on GOCACHE=/go/cache go test -race -mod=vendor ./..."

build-dirs:
	@echo "build-dirs"
	@mkdir -p dist
	@mkdir -p .go/src/$(PKG) .go/pkg .go/bin

clean-dirs:
	$(info clean up cache and dist folders)
	@rm -rf ./dist
	@rm -rf ./.go