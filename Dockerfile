# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

WORKDIR /AlphaClientServer

COPY go.mod /AlphaClientServer/
COPY go.sum /AlphaClientServer/

RUN go mod download

COPY /src/* /AlphaClientServer/

RUN go build -o /main


FROM golang:1.16-alpine

COPY --from=build /main /main

EXPOSE 8080

CMD [ "/main" ]     