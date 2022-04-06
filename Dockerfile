FROM golang:1.17

WORKDIR /go/src

RUN \
    apt-get update && apt-get upgrade -y &&\
    apt-get -y install postgresql-client

COPY . .
RUN chmod a+x waitPostgres.sh

RUN go build -o ./bin/URLshortener ./cmd/main.go

EXPOSE 8080