package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"go_docker/graphql/generated"
	"go_docker/graphql/resolver"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	// frontend connection, setting CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", corsMiddleware.Handler(srv))
	http.Handle("/query", corsMiddleware.Handler(loggingMiddleware(srv)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received at:", time.Now())
		next.ServeHTTP(w, r)
	})
}
