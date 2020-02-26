FROM golang:1.12 as builder

MAINTAINER LHM <cbg.hongminh@gmail.com>

WORKDIR /app

COPY . /app

RUN go mod download

RUN GOOS=linux

RUN go build -o user ./cmd/user.go

FROM ubuntu:16.04

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/user .

EXPOSE 8080

CMD ["/app/user"]
