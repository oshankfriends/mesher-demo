# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8-alpine3.6
ADD . /go/src/github.com/oshankfriends/mesher-demo/movies
WORKDIR /go/src/github.com/oshankfriends/mesher-demo/movies
RUN go install github.com/oshankfriends/mesher-demo/movies
ENTRYPOINT /go/bin/movies
# Service listens on port 5001.
EXPOSE 5001
