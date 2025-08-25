BUILD_DIR = ./bin
BINARY_NAME = gear
PKG_PATH = ./cmd/cli/main.go

## help: get info about the targets within this makefile
.phony: help
help:
	@echo "gear usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
	@echo "gear variables:"
	@echo "  BUILD_DIR: ${BUILD_DIR}"
	@echo "  BINARY_NAME: ${BINARY_NAME}"
	@echo "  PKG_PATH: ${PKG_PATH}"

## cli: runs the cli app in the package path
.phony: cli
cli:
	@go run ${PKG_PATH}

## build: build the application
.phony: build
build:
	@mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_DIR}/${BINARY_NAME}-darwin ${PKG_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME}-linux ${PKG_PATH}
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe ${PKG_PATH}

.phony: confirm
confirm:
	@echo -n "are you sure? [y/n] " && read ans && [ $${ans:-n} = y ]

## clean: clean up any build artifacts in the build directory
.phony: clean
clean: confirm
	@echo "cleaning up..."
	@rm -rf ${BUILD_DIR}

## tidy: tidies up the module and the test cache
.phony: tidy
tidy:
	@go clean -testcache
	@go mod tidy

## install: installs all dependencies for the module
.phony: install
install: tidy
	@go install ./...

## update: updates all packages used within the module
.phony: update
update: tidy
	@go get -u ./...

## test: runs unit tests and gives a coverage report
.phony: test
test: tidy
	@go test -cover ./...

## fmt: format the code
.phony: fmt
fmt:
	@gofmt -l .

## info: get info about the build environment, go version, packages, etc.
.phony: info
info:
	@go version
	@go env
	@go vet ./...
	@go list ./...
