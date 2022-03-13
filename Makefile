GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./scripts/build-go.sh
NEW_VERSION_BUILD=$(LOCAL_BUILD) new-version
NEW_VERSION_BUILD_DARWIN=$(NEW_VERSION_BUILD) darwin
NEW_VERSION_BUILD_LINUX=$(NEW_VERSION_BUILD) linux

run-tests:
	./scripts/run-tests $(path)

.PHONY: build-new-version
build-new-version: build-new-version-darwin build-new-version-linux

.PHONY: build-new-version-darwin
build-new-version-darwin: build-new-version-darwin-amd64 build-new-version-darwin-arm64

.PHONY: build-new-version-darwin-amd64
build-new-version-darwin-amd64:
	$(NEW_VERSION_BUILD_DARWIN) amd64

.PHONY: build-new-version-darwin-arm64
build-new-version-darwin-arm64:
	$(NEW_VERSION_BUILD_DARWIN) arm64

.PHONY: build-new-version-linux
build-new-version-linux: build-new-version-linux-386 build-new-version-linux-amd64 build-new-version-linux-arm build-new-version-linux-arm64

.PHONY: build-new-version-linux-386
build-new-version-linux-386:
	$(NEW_VERSION_BUILD_LINUX) 386

.PHONY: build-new-version-linux-amd64
build-new-version-linux-amd64:
	$(NEW_VERSION_BUILD_LINUX) amd64

.PHONY: build-new-version-linux-arm
build-new-version-linux-arm:
	$(NEW_VERSION_BUILD_LINUX) arm

.PHONY: build-new-version-linux-arm64
build-new-version-linux-arm64:
	$(NEW_VERSION_BUILD_LINUX) arm64

.PHONY: clean-build
clean-build:
	rm -rf .build/releaser*
