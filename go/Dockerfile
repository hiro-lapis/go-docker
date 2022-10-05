FROM golang:1.19.2-alpine3.16

ENV GOPATH /go
ENV GO111MODULE on
ENV ROOT=/go/src/app

# update and git install (go get`s inner process using git...)
RUN apk update && \
    apk --no-cache add git

# setting working dir
RUN mkdir ${ROOT}
WORKDIR ${ROOT}
# adding host files to containers
ADD . ${ROOT}

# import HOT RELOAD libraly
# not using go get (its deprecated)
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.40.4

CMD ["air", "-c", ".air.toml"]

# if you adding libraly, rerun build and up command...
