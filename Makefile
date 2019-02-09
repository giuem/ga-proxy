.PHONY: build clean

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=ga-proxy
OUTPUT_DIR=build

all: build

build:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)