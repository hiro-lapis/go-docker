FROM golang:1.16.9-alpine3.14
# TODO: adpating version up, after air updates
# FROM golang:1.17.2-alpine

ENV GOPATH /go
ENV GO111MODULE on

# update and git install (go get`s inner process using git...)
RUN apk update && \
    apk --no-cache add git

# setting working dir
RUN mkdir /go/src/app
WORKDIR /go/src/app

# adding host files to containers
ADD . /go/src/app

# import HOT RELOAD libraly
# in v1.17,rewrite go install github.com/cosmtrek/air@1.27.3
# not using go get (its deprecated)
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]

# if you adding libraly, rerun build and up command...
