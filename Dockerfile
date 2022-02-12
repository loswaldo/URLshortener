FROM golang:1.17

WORKDIR /go/src

COPY . .

RUN pwd

RUN go build -o ./bin/URLshortener ./cmd/main.go
CMD ["/go/src/bin/URLshortener"]

#ADD ./pkg ./pkg

#ADD ./servermain ./servermain

#ADD go.mod .
#
#ADD go.sum .
#
#RUN go build -o ./bin/grpc-fibonacci ./servermain
#
#ENTRYPOINT ["/go/src/bin/grpc-fibonacci"]

EXPOSE 8080