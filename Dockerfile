# syntax=docker/dockerfile:1

FROM golang:alpine

RUN apk add git

WORKDIR /app

COPY go.mod /app

RUN go mod download

COPY * /app

RUN go build -o /app/http-test

EXPOSE 8080

CMD [ "/app/http-test" ]