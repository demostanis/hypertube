GO = go
COMPOSE = docker-compose

-include conf.mk

build: secrets
	$(COMPOSE) up --build -d

watch:
	$(KILL_WATCH) 2>/dev/null || :
	$(COMPOSE) watch --no-up; echo $$! >.compose-watch.pid

logs:
	$(COMPOSE) logs -f $(S)

secrets:
	@./secrets.sh

down:
	$(COMPOSE) down

re: down build watch

tidy:
	$(GO) mod tidy
