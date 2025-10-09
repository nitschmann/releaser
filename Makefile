GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./scripts/build-go.sh
BUILD=$(LOCAL_BUILD) local
BUILD_DARWIN=$(BUILD) darwin
BUILD_LINUX=$(BUILD) linux
LATEST_BUILD=$(LOCAL_BUILD) latest
LATEST_DARWIN_BUILD=$(LATEST_BUILD) darwin
LATEST_LINUX_BUILD=$(LATEST_BUILD) linux
NEW_VERSION_BUILD=$(LOCAL_BUILD) new-version
NEW_VERSION_BUILD_DARWIN=$(NEW_VERSION_BUILD) darwin
NEW_VERSION_BUILD_LINUX=$(NEW_VERSION_BUILD) linux

# local builds

.PHONY: build
build: build-darwin build-linux

.PHONY: build-darwin
build-darwin: build-darwin-amd64 build-darwin-arm64

.PHONY: build-darwin-amd64
build-darwin-amd64:
	$(BUILD_DARWIN) amd64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	$(BUILD_DARWIN) arm64

.PHONY: build-linux
build-linux: build-linux-386 build-linux-amd64 build-linux-arm build-linux-arm64

.PHONY: build-linux-386
build-linux-386:
	$(BUILD_LINUX) 386

.PHONY: build-linux-amd64
build-linux-amd64:
	$(BUILD_LINUX) amd64

.PHONY: build-linux-arm
build-linux-arm:
	$(BUILD_LINUX) arm

.PHONY: build-linux-arm64
build-linux-arm64:
	$(BUILD_LINUX) arm64

# latest builds

.PHONY: build-latest
build-latest: build-latest-darwin build-latest-linux

.PHONY: build-latest-darwin
build-latest-darwin: build-latest-darwin-amd64 build-latest-darwin-arm64

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

# new version builds

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

# Clean ./build folder
.PHONY: clean-build
clean-build:
	rm -rf .build/releaser*

# Run the tests locally on the machine
.PHONY: run-tests
run-tests:
	./scripts/run-tests $(path)

TEST_CMD = go test -parallel=4 -v -cover -race ./...
.PHONY: docs
docs:
	@echo "==> Generating CLI docs..."
	$(GOCMD) run ./tools/gendocs
	@echo "==> Docs generated into docs/cli/"


.PHONY: test
test:
	docker-compose build --pull dev
		docker-compose run --rm dev \
		sh -c '${TEST_CMD}'

.PHONY: lint lint-docker lint-local
lint:
	@echo "==> Running golint (local)..."
	golangci-lint run --timeout 2m

lint-docker:
	@echo "==> Running golint (docker)..."
	docker-compose build --pull dev
		docker-compose run --rm dev \
		sh -c 'go mod download && golangci-lint run --timeout 2m --modules-download-mode=mod'

# Kill and remove the local docker-compose containers
.PHONY: down
down:
	docker-compose down --remove-orphans --volumes
