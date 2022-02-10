GO=$(shell which go)
BUILDPATH=$(CURDIR)
EXENAME=app

all: makedir deps build test

makedir:
	@echo "Ensure directories exist..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi

deps:
	@echo "Installing dependencies..."
	$(GO) mod download
	@echo "Done."

build: deps makedir
	@echo "Building binary..."
	$(GO) build -o $(BUILDPATH)/bin/$(EXENAME)
	@echo "Done."

test: deps
	@echo "Running tests..."
	$(GO) test ./...
	@echo "Done."

run: build
	@echo "Running app..."
	$(BUILDPATH)/bin/$(EXENAME)


clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILDPATH)/bin
	@echo "Done."
