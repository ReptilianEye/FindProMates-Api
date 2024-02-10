package main

import (
	"example/FindProMates-Api/graph"
	"example/FindProMates-Api/internal/app"
	"example/FindProMates-Api/internal/auth"
	"example/FindProMates-Api/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const defaultPort = "8080"

func main() {
	cancel := app.InitApp()
	defer cancel()
	defer database.CloseDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(auth.Middleware)
	router.Use(middleware.Recoverer)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
