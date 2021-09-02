.PHONY: build

build:
	go build ./cmd/api

.PHONY: build-worker

build-worker:
	go build ./cmd/worker

.PHONY: run

run:
	go run ./cmd/api/main.go
.PHONY: run-worker

run-worker:
	go run ./cmd/worker/main.go

.DEFAULT_GOAL := run