GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./scripts/build-go.sh
LATEST_BUILD=$(LOCAL_BUILD) latest
LATEST_DARWIN_BUILD=$(LATEST_BUILD) darwin
LATEST_LINUX_BUILD=$(LATEST_BUILD) linux
NEW_VERSION_BUILD=$(LOCAL_BUILD) new-version
NEW_VERSION_BUILD_DARWIN=$(NEW_VERSION_BUILD) darwin
NEW_VERSION_BUILD_LINUX=$(NEW_VERSION_BUILD) linux

.PHONY: install-test-dependencies
install-test-dependencies:
	env GO111MODULE=on go get -u github.com/client9/misspell/cmd/misspell
	env GO111MODULE=on go get -u golang.org/x/lint/golint

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: check-misspell
check-misspell:
	misspell ./**/* -error

.PHONY: build-latest
build-latest: build-latest-darwin build-latest-linux

.PHONY: build-latest-darwin
build-latest-darwin: build-latest-darwin-amd64

.PHONY: build-latest-darwin-amd64
build-latest-darwin-amd64:
	$(LATEST_DARWIN_BUILD) amd64

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

.PHONY: build-new-version
build-new-version: build-new-version-darwin build-new-version-linux

.PHONY: build-new-version-darwin
build-new-version-darwin: build-new-version-darwin-amd64

.PHONY: build-new-version-darwin-amd64
build-new-version-darwin-amd64:
	$(NEW_VERSION_BUILD_DARWIN) amd64

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

.PHONY: generate-docs
generate-docs:
	$(GOCMD) run tools/gendocs/main.go

.PHONY: prepare-release
prepare-release: generate-docs
