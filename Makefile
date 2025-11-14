.PHONY: up down create-service update-dependencies build

ifneq ("$(wildcard .env)","")
    include .env
endif

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

reup:
	make down
	make build
	make up

create-service:
	@mkdir -p app/back/services/$(SERVICE_NAME); \
	cd app/back/services/$(SERVICE_NAME) && go mod init $(SERVICE_NAME) && \
	go get -u gorm.io/gorm && \
	go get -u gorm.io/driver/postgres && \
	go get -u github.com/gin-gonic/gin && \
	go get -u github.com/joho/godotenv && \
	sed -i 's/go 1.23.1/go 1.23/' go.mod && \
	echo "Сервис ${SERVICE_NAME} успешно создан в директории app/back/services/${SERVICE_NAME}"

update-dependencies:
	@read -p "Введите имя пакета для установки: " package_name; \
	cd app/back/services/${SERVICE_NAME} && go get -u $$package_name && \
	echo "Пакет $$package_name успешно установлен в сервис ${SERVICE_NAME}"