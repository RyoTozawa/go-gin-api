FROM golang:1.8.5-jessie

ENV GOPATH $GOPATH:/go/src

RUN go get -u "github.com/gin-gonic/gin" \
&& go get -u "github.com/jinzhu/gorm" \
&& go get -u "github.com/lib/pq"
RUN mkdir /go/src/api

EXPOSE 8080