.PHONY: all lint

all:
	docker-compose up
lint:
	golangci-lint run ./...