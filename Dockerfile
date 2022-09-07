FROM golang:1.18-alpine

RUN apk add git
RUN go install github.com/YaromyrO/watch-out@v0.1.0

CMD /go/bin/watch-out