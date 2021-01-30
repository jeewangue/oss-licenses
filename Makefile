MODULE_NAME=github.com/jeewangue/oss-licenses
BIN_DIR=bin
BIN_NAME=oss-licenses

# DEST_DIR is where the program should be installed
DEST_DIR = $$HOME/.local/bin

.PHONY: all
all: build

.PHONY: download-deps
download-deps:
	@echo Download go.mod dependencies
	@go mod download

.PHONY: install-tools
install-tools: download-deps
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: test
test:
	go test -cover ./...

.PHONY: fmt
fmt:
	test -z $(shell go fmt ./...)

.PHONY: clean
clean:
	@rm -rf ${BIN_DIR} || true

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build: | fmt lint test
	go build -o ${BIN_DIR}/${BIN_NAME} ./cmd

.PHONY: install
install: build
	mkdir -p ${DEST_DIR}
	install ${BIN_DIR}/${BIN_NAME} ${DEST_DIR}/${BIN_NAME}

