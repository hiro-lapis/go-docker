## Intro

This is an Golang(v1.19) and mysql(v8.0) dev-environment with hot reload(air).
If you run `docker compose up`, go container has up with hot reload, instantly!

## Install Guide

### up container with hot Reload

```shell
docker compose up
```

### interactive tty container

-   application

```shell
docker exec -it go-docker_go_1 ash
```

-   database
    about following `MYSQL_USER` `MYSQL_PASSWORD`, refer to docker-compose.yml

```shell
 docker exec -it go-docker_db_1 bash -c 'mysql -u <MYSQL_USER> -p'
```

`Enter password: <MYSQL_PASSWORD>`

### migration

```
// create
migrate create -ext sql -dir database/migrations -seq create_users

// up
migrate -path database/migrations -database $MYSQL_URL up 1
// back
migrate -path database/migrations -database $MYSQL_URL down 1
```

## References

## Dockerfile docker-compose.yml

https://docs.docker.jp/engine/reference/builder.html

## Golang image

https://hub.docker.com/_/golang?tab=description&page=1
