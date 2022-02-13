FROM golang:1.17

WORKDIR /go/src

COPY . .
RUN chmod a+x waitPostgres.sh

RUN go build -o ./bin/URLshortener ./cmd/main.go

RUN apt-get update && apt-get upgrade -y
RUN apt-get -y install postgresql-client

EXPOSE 8080