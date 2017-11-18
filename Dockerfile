FROM golang:1.8.5-jessie

RUN go get -u "github.com/gin-gonic/gin"

WORKDIR /go/
ADD . /go

EXPOSE 8080

CMD ["go", "run", "main.go"]