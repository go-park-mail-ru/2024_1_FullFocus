.DEFAULT_GOAL:=run

.PHONY: lint run build test

lint:
	golangci-lint run ./...

run:
	go run cmd/main/main.go

build:
	go build -o ./bin/app ./cmd/main/main.go

test:
	@go test ./... -cover > testresult.txt
	@sed -i '/mock/d' testresult.txt
	@sed -i '/main/d' testresult.txt
	@sed -i '/pkg/d' testresult.txt
	@sed -i '/models/d' testresult.txt
	@cat testresult.txt
	@rm testresult.txt
