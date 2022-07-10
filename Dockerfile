# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /AlphaClientServer

COPY go.mod /AlphaClientServer/
COPY go.sum /AlphaClientServer/

RUN go mod download

COPY /src/* /AlphaClientServer/

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]
