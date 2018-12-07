# Usage:
# make install # prepare all .env files and build docker images
# make serve   # start web api server on http://localhost:8003
# make lint    # run gometalinter v2 for the whole application

.PHONY: install serve lint migrations

include .docker/mysql/.env

install: install_confirmation
	cp .docker/golang/api/.env.example .docker/golang/api/.env
	cp .docker/mysql/.env.example .docker/mysql/.env
	docker network create fintech-backend
	docker-compose build


serve:
	WEB_API_ADDRESS=:8003 go run -race cmd/api/main.go

lint:
	docker-compose exec ft-web-api-golang gometalinter.v2 ./...

migrate-status:
	goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}" status

migrate-up:
	goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}" up

migrate-down:
	goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}" down

migrate-create:
	@echo  "Please prove name of your migration(example: create_users_table) or terminate it (ctrl+C)"
	@read input && goose -dir migrations mysql "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(ft-mysql)/${MYSQL_DATABASE}" create $$input sql


install_confirmation:
	@echo  "Would you like to start installation [y/N]? " && read ans && [ $$ans == y ]
