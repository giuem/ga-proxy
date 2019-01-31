.PHONY: build clean

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=ga_proxy
OUTPUT_DIR=build

all: build

build:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf $(OUTPUT_DIR)