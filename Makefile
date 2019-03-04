# Usage:
# make install # prepare all .env files and build docker images
# make serve   # start web api server on http://localhost:8003
# make lint    # run gometalinter v2 for the whole application

.PHONY: install serve lint migrate-status migrate-up migrate-down migrate-create install_confirmation

install: install_confirmation
	cp .docker/golang/api/.env.example .docker/golang/api/.env
	cp .docker/mysql/.env.example .docker/mysql/.env
	docker network create fintech-backend
	docker-compose build


serve:
	DB_USER=root DB_PASS=rootroot DB_HOST=localhost DB_NAME=fintech DB_PORT=3306 WEB_API_ADDRESS=:8003 go run -race cmd/api/main.go

lint:
	docker-compose exec ft-web-api-golang gometalinter.v2 ./...

migrate-status:
	export $(grep -v '^#' .docker/mysql/.env | xargs -0)
	docker-compose exec ft-web-api-golang goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}?parseTime=true&charset=utf8" status

migrate-up:
	export $(grep -v '^#' .docker/mysql/.env | xargs -0)
	docker-compose exec ft-web-api-golang goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}?parseTime=true&charset=utf8" up

migrate-down:
	export $(grep -v '^#' .docker/mysql/.env | xargs -0)
	docker-compose exec ft-web-api-golang goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}?parseTime=true&charset=utf8" down

migrate-create:
	export $(grep -v '^#' .docker/mysql/.env | xargs -0)
	@echo  "Please prove name of your migration(example: create_users_table) or terminate it (ctrl+C)"
	@read input && docker-compose exec ft-web-api-golang goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}?parseTime=true&charset=utf8" create $$input sql

install_confirmation:
	@echo  "Would you like to start installation [y/N]? " && read ans && [ $$ans == y ]
