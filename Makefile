GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./scripts/build-go.sh
LATEST_BUILD=$(LOCAL_BUILD) latest
LATEST_DARWIN_BUILD=$(LATEST_BUILD) darwin
LATEST_LINUX_BUILD=$(LATEST_BUILD) linux

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: build-latest
build-latest: build-latest-darwin build-latest-linux

.PHONY: build-latest-darwin
build-latest-darwin: build-latest-darwin-386 build-latest-darwin-amd64

.PHONY: build-latest-darwin-386
build-latest-darwin-386:
	$(LATEST_DARWIN_BUILD) 386

.PHONY: build-latest-darwin-amd64
build-latest-darwin-amd64:
	$(LATEST_DARWIN_BUILD) amd64

.PHONY: build-latest-darwin-arm64
build-latest-darwin-arm64:
	$(LATEST_DARWIN_BUILD) arm64

.PHONY: build-latest-linux
build-latest-linux: build-latest-linux-386 build-latest-linux-amd64 build-latest-linux-arm build-latest-linux-arm64

.PHONY: build-latest-linux-386
build-latest-linux-386:
	$(LATEST_LINUX_BUILD) 386

.PHONY: build-latest-linux-amd64
build-latest-linux-amd64:
	$(LATEST_LINUX_BUILD) amd64

.PHONY: build-latest-linux-arm
build-latest-linux-arm:
	$(LATEST_LINUX_BUILD) arm

.PHONY: build-latest-linux-arm64
build-latest-linux-arm64:
	$(LATEST_LINUX_BUILD) arm64

.PHONY: clean-build
clean-build:
	rm -rf .build/release-log*

.PHONY: generate-docs
generate-docs:
	$(GOCMD) run tools/gendocs/main.go

.PHONY: prepare-release
prepare-release: generate-docs
