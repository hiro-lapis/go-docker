# environment variable
DB_HOST ?= db
DB_PORT ?= 3306
DB_USER ?= user
DB_PASSWORD ?= password

# up api server
up_server:
	go run ./cmd/server/main.go
# initialize graphql schema
gqlgen_c:
	gqlgen generate

# update graphql schema
gqlgen_u:
	rm -vf ./graphql/generated/* && gqlgen generate
