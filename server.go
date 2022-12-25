package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sugaml/developer-news/graph"
	"github.com/Sugaml/developer-news/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error in load env file %v", err)
	} else {
		fmt.Println("Loaded env files...")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// conf, err := postgres.LoadConfig()
	// if err != nil {
	// 	log.Fatal("Config load error", err)
	// }
	// db := postgres.NewDB(conf)
	// migrations.MigarateUp(db)

	// router := chi.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
