GO = go

run:
	docker-compose up --build

tidy:
	$(GO) mod tidy
