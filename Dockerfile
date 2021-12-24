FROM golang:1.16-alpine3.15

WORKDIR /go/src

ENTRYPOINT [ "top" ]