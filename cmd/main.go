package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andreyxaxa/posts_comments_service/configs"
	"github.com/andreyxaxa/posts_comments_service/graph"
	resolvers "github.com/andreyxaxa/posts_comments_service/internal/server/graphql"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	logger.Info.Print("exec NewLogger")

	envFile := ".env"
	if len(os.Args) >= 2 {
		envFile = os.Args[1]
	}

	logger.Info.Print("exec InitConfig")

	logger.Info.Printf("reading %s\n", envFile)
	if err := configs.InitConfigs(envFile); err != nil {
		logger.Error.Fatal(err.Error())
	}

	port := os.Getenv("PORT")

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	logger.Info.Print("Connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Error.Fatal(http.ListenAndServe(":"+port, nil))
}
