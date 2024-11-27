GO = go
COMPOSE = docker-compose

-include conf.mk

build:
	$(COMPOSE) up --build --watch

stop:
	$(COMPOSE) down

re: stop build

tidy:
	$(GO) mod tidy
