BUILDPATH=$(CURDIR)
EXENAME=app

define CREATE_TABLE
CREATE TABLE IF NOT EXISTS books (
	slug 		TEXT	NOT NULL	UNIQUE,
	title		TEXT	NOT NULL,
	description	TEXT	NOT NULL,
	release_time 	INTEGER	NOT NULL
);
endef

define INSERT_BOOKS
INSERT INTO books
VALUES ("AzureWitch", "Death of the Azure Witch", "This is a real book.", 1646006400);
INSERT INTO books
VALUES ("BloodOath", "Blood Oath", "Sometimes, someone needs to die.", 1644796800);
endef

export CREATE_TABLE
export INSERT_BOOKS

all: makedir deps build test

makedir:
	@echo "Ensure directories exist..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi

deps:
	@echo "Installing dependencies..."
	@go mod download
	@echo "Done."

build: deps makedir
	@echo "Building binary..."
	@go build -o $(BUILDPATH)/bin/$(EXENAME)
	@echo "Done."

test: deps
	@echo "Running tests..."
	@go test ./...
	@echo "Done."

run: build
	@echo "Running app..."
	@$(BUILDPATH)/bin/$(EXENAME)

dev-db:
	@echo "Populating development database..."
	@sqlite3 vorona.db "$$CREATE_TABLE"
	@sqlite3 vorona.db "$$INSERT_BOOKS"
	@echo "Done."

docker:
	@-if (docker ps -a | grep -q vorona); then \
		echo "Cleaning up old Docker container..."; \
		docker rm --force vorona; \
	fi
	@touch vorona.db
	@echo "Building Docker image..."
	@docker build -t vorona .
	@echo "Build complete."
	@echo "Running container..."
	@docker run \
		--name vorona \
		-p 8080:8080 \
		-v $(CURDIR)/vorona.db:/usr/src/app/vorona.db \
		-d \
		vorona

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILDPATH)/bin
	@echo "Done."
