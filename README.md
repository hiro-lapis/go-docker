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

## Set up

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
docker exec -it go-docker-go-1 ash
```

-   database
    about following `MYSQL_USER` `MYSQL_PASSWORD`, refer to docker-compose.yml

```shell
 docker exec -it go-docker_db_1 bash -c 'mysql -u <MYSQL_USER> -p'
```

`Enter password: <MYSQL_PASSWORD>`

## app

### Dir structure

```
.
├── Makefile  // defines set of tasks to be commonly executed 
├── README.md
├── cmd // server entry point of as Graphql server
│   └── server.go
├── database
│   ├── driver
│   │   └── driver.go
│   ├── migrations
│   │   ├── 000001_create_users.down.sql
│   │   └── 000001_create_users.up.sql
│   └── users
│       └── query.go
├── domain
│   └── model
│       └── models_gen.go
├── fizzbuzz
│   ├── main.go
│   └── main_test.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graphql
│   ├── db
│   │   ├── boil_queries.go
│   │   ├── boil_table_names.go
│   │   ├── boil_types.go
│   │   └── boil_view_names.go
│   ├── generated
│   │   └── generated.go
│   ├── graph
│   │   └── schema.graphqls
│   └── resolver
│       ├── resolver.go
│       └── schema.resolvers.go
├── main.go
├── sqlboiler.toml
├── sub
│   └── sub.go
└── tmp
    └── main
```

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

### Graphql

#### setup

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


#### run as Graphql server

```go
go run cmd/server.go

2023/06/03 15:47:46 connect to http://localhost:8080/ for GraphQL playground
```

then you can access PlayGround `http://localhost:8080` on browser.  
then you can play graphql by throwing query.

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

#### how to implement graphql function

#### add query

for add query or mutation, you need to edit server side schema `app/graphql/graph/**.graphqls`.

```
// app/graphql/graph/schema.graphqls


type getTodos {
  todos: [Todo!]!
}
```

then do generate command by using serverside Makefile on app/.

```
make gqlgen_c
```

gqlgen will load under app/graphql/graph, generating automatically.

next step, you also add schema on front end schema `frontend/src/graphql/**.graphql`.

```
# frontend/src/graphql/Todo.graphql

query getTodos {
  getTodos { # <- this point is `query field`, is need correspond for serverside query-type name
    id
    text
  }
}
```

and, you run pnpm to generate.
Note: for thanks of codegen.ts settings, frontend schema will be generated by according to **serverside schema**.  
thus, you need **run app server in advance**.

```
pnpm graphql:generate
```

Gotcha! you can import and using query composable imediately in SFC!


```
// in, sfc

import {
  useGetTodosQuery,
 } from '~/src/generated/graphql'
```

#### implement resolver

so far, you still not implement actual logic.

To implement actual logic, in this case `getTodos` logic, 




## frontend

```
docker exec -it go-docker-node-1 ash
```

- initialize project

```
// create with overwrite frontend directory

cd ../ && pnpm dlx nuxi init --force frontend && pnpm i --frozen-lockfile	 
```

- run frontend dev server

```
pnpm run dev
```
then, access http://localhost:4000

## FAQ

Q. `gqlgen generate` results in `gqlgen: not found`   
A. check PATH.

```
echo $PATH
# /go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin


echo $GOPATH
# /go

which go
# /usr/local/go/bin/go
export PATH="$HOME/go/bin:$PATH"
```



## References

## Dockerfile docker-compose.yml

https://docs.docker.jp/engine/reference/builder.html

## Golang image

https://hub.docker.com/_/golang?tab=description&page=1
