# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

EXPOSE 4200

COPY ./ ./

RUN go build -o gostagram server.go

CMD ["./gostagram"]
