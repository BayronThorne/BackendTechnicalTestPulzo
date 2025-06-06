.PHONY: build up down test

# Runs tests
test:
	go test ./auth ./datafetcher

# Builds the Docker images
build:
	docker compose build

# Starts the services
up:
	docker compose up -d

# Stops the services
down:
	docker compose down
