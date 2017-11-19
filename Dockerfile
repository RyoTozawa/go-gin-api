FROM golang:1.8.5-jessie

ENV GOPATH $GOPATH:/go/src

RUN go get -u "github.com/gin-gonic/gin"

RUN mkdir /go/src/api

EXPOSE 8080