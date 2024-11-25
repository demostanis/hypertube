GO = go

run:
	docker-compose up --build

stop:
	docker-compose down

re: stop run

tidy:
	$(GO) mod tidy
