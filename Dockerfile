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

ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie

RUN chmod 755 /usr/bin/aws-lambda-rie

COPY /SH/entry.sh /

RUN chmod 755 /entry.sh

ENTRYPOINT [ "/entry.sh" ]        