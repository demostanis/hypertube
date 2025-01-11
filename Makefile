GO = go
COMPOSE = docker-compose

-include conf.mk

all: up watch

up: secrets
	$(COMPOSE) up -d

build:
	$(COMPOSE) build $(S)

watch:
	$(KILL_WATCH) 2>/dev/null || :
	$(COMPOSE) watch --no-up; echo $$! >.compose-watch.pid

logs:
	$(COMPOSE) logs -f $(S)

exec:
	$(COMPOSE) exec $(S) bash || $(COMPOSE) exec $(S) sh

secrets:
	@./secrets.sh

down:
	$(COMPOSE) down

re: down build up watch

tidy:
	$(GO) mod tidy
