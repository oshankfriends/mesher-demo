# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8-alpine3.6
ADD . /go/src/github.com/oshankfriends/mesher-demo/showtimes
WORKDIR /go/src/github.com/oshankfriends/mesher-demo/showtimes
RUN go install github.com/oshankfriends/mesher-demo/showtimes
ENTRYPOINT /go/bin/showtimes
# Service listens on port 5002.
EXPOSE 5002
