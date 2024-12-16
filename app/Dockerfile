FROM golang:1.23.4

COPY app/static /app/static

WORKDIR /app

RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=bind,src=app,target=.,rw \
	--mount=type=bind,src=db/models,target=./models \
	go build -v -o /exe *.go

CMD ["/exe"]
