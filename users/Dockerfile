# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8-alpine3.6
ADD . /go/src/github.com/oshankfriends/mesher-demo/users
WORKDIR /go/src/github.com/oshankfriends/mesher-demo/users
RUN go install github.com/oshankfriends/mesher-demo/users
ENTRYPOINT /go/bin/users
# Service listens on port 5000.
EXPOSE 5000
