# Usage:
# make install # prepare all .env files and build docker images
# make serve   # start web api server on http://localhost:8003
# make lint    # run gometalinter v2 for the whole application

.PHONY: install serve lint migrations


install: install_confirmation
	cp .docker/golang/api/.env.example .docker/golang/api/.env
	cp .docker/mysql/.env.example .docker/mysql/.env
	docker network create fintech-backend
	docker-compose build


serve:
	WEB_API_ADDRESS=:8003 go run -race cmd/api/main.go

lint:
	docker-compose exec ft-web-api-golang gometalinter.v2 ./...

migrations:
	echo 12

install_confirmation:
	@echo  "Would you like to start installation [y/N]? " && read ans && [ $$ans == y ]
