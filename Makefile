DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=ozon

include .env

DB_DSN="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

.PHONY: setup
setup: ## Установить все необходимые утилиты
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.1

.PHONY: up
up: ## Поднять контейнеры
	docker compose up -d

.PHONY: down
down: ## Остановить контейнеры
	docker compose down

.PHONY: migrations-up
migrations-up: ## Накатить миграции
	 goose -dir db/migrations postgres $(DB_DSN) up

.PHONY: migrations-down
migrations-down: ## Откатить миграции
	 goose -dir db/migrations postgres $(DB_DSN) down

.PHONY: run-app
run-app: up ## Запустить приложение
	make migrations-up

.PHONY: stop-app
stop-app: down ## Остановить приложение

.PHONY: build
build: ## Сбилдить бинарь приложения
	go build -o ./bin/app ./cmd/main/main.go

.PHONY: lint
lint: ## Проверить код линтерами
	golangci-lint run ./... -c golangci.local.yaml

.PHONY: test
test: ## Запустить тесты
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

.PHONY: ci
ci: lint test ## Запустить линтеры + тесты

.PHONY: clean
clean: ## Удалить временные файлы
	rm -f testresult.txt
	rm -f ./bin/app

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:=help
