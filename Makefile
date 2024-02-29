.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run cmd/main/main.go

.DEFAULT_GOAL:=run
