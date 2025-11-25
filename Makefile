.PHONY: help generate bundle ogen build up down restart logs test clean

# Переменные
OPENAPI_SOURCE = shared/api/reviewservice/reviewservice.openapi.yaml
OPENAPI_BUNDLE = shared/api/bundled/openapi-bundled.yaml
OGEN_TARGET = shared/pkg/openapi/reviewerservice/v1
OGEN_PACKAGE = reviewerservice


generate: bundle ogen 

bundle: 
	@echo "Сборка OpenAPI спецификации..."
	npx redocly bundle $(OPENAPI_SOURCE) -o $(OPENAPI_BUNDLE)
	@echo "✓ Спецификация собрана: $(OPENAPI_BUNDLE)"

ogen: 
	@echo "Генерация Go кода..."
	.\bin\ogen.exe --target ./$(OGEN_TARGET) --package $(OGEN_PACKAGE) --clean ./$(OPENAPI_BUNDLE)
	@echo "✓ Код сгенерирован в $(OGEN_TARGET)"

build: 
	@echo "Сборка Docker образов..."
	docker-compose build
	@echo "✓ Образы собраны"

up: 
	@echo "Запуск приложения..."
	docker-compose up -d
	@echo "✓ Приложение запущено на http://localhost:8080"

down: 
	@echo "Остановка приложения..."
	docker-compose down
	@echo "✓ Приложение остановлено"

restart: down up ## Перезапустить приложение

