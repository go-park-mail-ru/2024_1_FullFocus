.DEFAULT_GOAL:=run

.PHONY: run lint test

lint:
	golangci-lint run

run:
	go run cmd/main/main.go

test:
	@go test ./... -cover > testresult.txt
	@sed -i '/mock/d' testresult.txt
	@sed -i '/main/d' testresult.txt
	@sed -i '/pkg/d' testresult.txt
	@sed -i '/models/d' testresult.txt
	@cat testresult.txt
	@rm testresult.txt
