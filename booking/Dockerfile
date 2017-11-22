# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8-alpine3.6
ADD . /go/src/github.com/oshankfriends/mesher-demo/booking
WORKDIR /go/src/github.com/oshankfriends/mesher-demo/booking
RUN go install github.com/oshankfriends/mesher-demo/booking
ENTRYPOINT /go/bin/booking
# Service listens on port 5003.
EXPOSE 5003
