exename := "vorona"
docker-image := "registry.digitalocean.com/vorona/vorona"
dev-db := justfile_directory() + "/vorona.db"

# Replace vorona.db with fresh copy
dev-db:
    go run *.go --sqlite {{dev-db}} dev-db

# Create fresh vorona.db; lint and run tests
test: dev-db
    go fmt ./...
    go vet ./...
    go test ./...

# Build binary in bin/vorona
build:
    @if [ ! -d bin ]; then mkdir bin; fi
    go build -o bin/{{exename}}

# Build and run binary in dev mode
run: build
    ./bin/vorona --sqlite {{dev-db}} start

# Docker Commands
# Build latest vorona image
docker-build:
    docker build -t {{docker-image}} .

# Remove old container and build & run latest
docker-run: docker-build
    #!/usr/bin/env bash
    if (docker ps -a | grep -q vorona); then
        echo "Cleaning up old docker container...";
        docker rm --force vorona;
    fi
    docker run \
        --name vorona \
        -p 8080:8080 \
        -v {{justfile_directory()}}/vorona.db:/opt/vorona/vorona.db \
        -d \
        {{docker-image}}

# Build and Publish Image to DigitalOcean Registry
docker-publish: docker-build
    doctl registry login
    docker push {{docker-image}}
