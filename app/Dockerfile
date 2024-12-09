FROM golang:1.23.4

COPY static /app/static

WORKDIR /app

RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=bind,src=.,target=. \
	go build -v -o /exe *.go

CMD ["/exe"]
