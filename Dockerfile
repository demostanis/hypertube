FROM golang

WORKDIR /app

RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=bind,src=app,target=. \
	go build -v -o /exe *.go

CMD ["/exe"]
