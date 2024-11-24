FROM golang

COPY . /app
WORKDIR /app

RUN go mod download && go mod verify
RUN go build -v -o app *.go

CMD ["./app"]
