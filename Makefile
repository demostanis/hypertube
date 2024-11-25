GO = go
COMPOSE = docker-compose

-include conf.mk

# build should probably be the default rule in the future,
# but for now i don't want the hypertube container to rebuild everytime.
up:
	$(COMPOSE) up

build:
	$(COMPOSE) up --build

stop:
	$(COMPOSE) down

re: stop build

tidy:
	$(GO) mod tidy
