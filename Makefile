MAIN_PACKAGE_PATH := ./penprep.go
BINARY_NAME := penprep

## build: build the application
.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: 
	/tmp/bin/${BINARY_NAME}

## install: install the application
.PHONY: install
install: 
	cp /tmp/bin/${BINARY_NAME} ${HOME}/go/bin/
