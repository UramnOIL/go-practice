package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go-practice/graph"
	"go-practice/graph/generated"
	"log"
	"net/http"
)

func main() {
	println("Hello, world!")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("Playground - GraphQL", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
