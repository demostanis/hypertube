GO = go
COMPOSE = docker-compose

-include conf.mk

build: secrets
	$(COMPOSE) up --build --watch

secrets:
	@./secrets.sh

stop:
	$(COMPOSE) down

re: stop build

tidy:
	$(GO) mod tidy
