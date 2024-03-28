.DEFAULT_GOAL:=up

.PHONY: lint up down build test clean

lint:
	golangci-lint run ./... -c golangci.yaml

up:
	docker compose up -d

down:
	docker compose down

build:
	go build -o ./bin/app ./cmd/main/main.go

test:
	@go test ./... -cover > testresult.txt
	@sed -i '/dto/d' testresult.txt
	@sed -i '/server/d' testresult.txt
	@sed -i '/app/d' testresult.txt
	@sed -i '/config/d' testresult.txt
	@sed -i '/mock/d' testresult.txt
	@sed -i '/main/d' testresult.txt
	@sed -i '/pkg/d' testresult.txt
	@sed -i '/models/d' testresult.txt
	@cat testresult.txt
	@rm testresult.txt

clean:
	rm testresult.txt
