BUILD_DIR = ./bin
BINARY_NAME = gear
PKG_PATH = ./cmd/cli/main.go

## help: get info about the targets within this makefile
.PHONY: help
help:
	@echo "gear usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
	@echo "gear variables:"
	@echo "  BUILD_DIR: ${BUILD_DIR}"
	@echo "  BINARY_NAME: ${BINARY_NAME}"
	@echo "  PKG_PATH: ${PKG_PATH}"

## cli: runs the cli app in the package path
.PHONY: cli
cli:
	@go run ${PKG_PATH}

## build: build the application
.PHONY: build
build:
	@mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_DIR}/${BINARY_NAME}-darwin ${PKG_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME}-linux ${PKG_PATH}
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe ${PKG_PATH}

.PHONY: confirm
confirm:
	@echo -n "are you sure? [y/n] " && read ans && [ $${ans:-n} = y ]

## clean: clean up any build artifacts in the build directory
.PHONY: clean
clean: confirm
	@echo "cleaning up..."
	@rm -rf ${BUILD_DIR}

## tidy: tidies up the module and the test cache
.PHONY: tidy
tidy:
	@go clean -testcache
	@go mod tidy

## install: installs all dependencies for the module
.PHONY: install
install: tidy
	@go install ./...

## update: updates all packages used within the module
.PHONY: update
update: tidy
	@go get -u ./...

## test: runs unit tests and gives a coverage report
.PHONY: test
test:
	@go test -cover ./...

# bench: runs benchmarks for the module
.PHONY: bench
bench:
	@go test -bench=. ./...

## fmt: format the code
.PHONY: fmt
fmt:
	@gofmt -l .

## info: get info about the build environment, go version, packages, etc.
.PHONY: info
info:
	@go version
	@go env
	@go vet ./...
	@go list ./...
