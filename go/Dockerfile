FROM golang:1.19.2-alpine3.16

ENV GOPATH /go
ENV GO111MODULE on
ENV ROOT=/go/src/app
ENV MYSQL_URL mysql://user:pw@tcp(go-docker_db_1:3306)/db
ENV DSN user:pw@tcp(go-docker_db_1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local

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
    go install github.com/cosmtrek/air@v1.40.4 && \
    go get github.com/go-sql-driver/mysql && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    # go get gorm.io/gorm && \
    go get gorm.io/driver/mysql
CMD ["air", "-c", ".air.toml"]

# if you adding libraly, rerun build and up command...
