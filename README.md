## Intro

- Golang(v.1.19)
- mysql(v.8.0)
- nuxt(v3.2) with node(v18.14)

This is an Golang and mysql dev-environment with hot reload(air).
If you run `docker compose up`, go container has up with hot reload, instantly!
frontend is also ignitionable.


| container | HOST PORT  | CONTAINER PORT |
|-----------| ---- |----------------|
| golang    |  8080  | 8080           |
| mysql     |  3306  | 3306           |
| node      |  4000  | 3000           |

## Install Guide

### up all container

```shell
docker compose up -d --build
```

### up back container with hot Reload

```shell
docker compose up go
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

### run as web server

comment in below at main.go

```.go
	/*
	 * run as web server
	 * listen request
	 * access to http://localhost:8080
	 */
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
```

then, `go run main.go`

## Graphql setup

```
// in go container, at ./app
go install github.com/99designs/gqlgen@latest
gqlgen version

v0.17.22
```

then, generate sample code (on this project, already generated)

```
go get -u github.com/99designs/gqlgen

gqlgen init
```

then you can run sample Graphql server
Avoiding duplicates listen port, make sure modifing main.go's code, http port except '8080'


```go
go run graphql/server.go

2023/06/03 15:47:46 connect to http://localhost:8080/ for GraphQL playground
```

then you can access PlayGround `http://localhost:8080` on browser.  
then you can throw query.

```
query {
  todos {
    id
    text
    done
    user {
      name
    }
  }
}
```

## frontend

```
docker exec -it go-docker_node_1 ash
```

- initialize project

```
// create with overwrite frontend directory

cd ../ && pnpm dlx nuxi init --force frontend && pnpm install
```

- run frontend dev server

```
pnpm run dev
```
then, access http://localhost:4000


## References

## Dockerfile docker-compose.yml

https://docs.docker.jp/engine/reference/builder.html

## Golang image

https://hub.docker.com/_/golang?tab=description&page=1
