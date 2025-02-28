package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andreyxaxa/posts_comments_service/configs"
	"github.com/andreyxaxa/posts_comments_service/graph"
	"github.com/andreyxaxa/posts_comments_service/internal/db"
	"github.com/andreyxaxa/posts_comments_service/internal/gateway"
	"github.com/andreyxaxa/posts_comments_service/internal/gateway/postgres"
	resolvers "github.com/andreyxaxa/posts_comments_service/internal/server/graphql"
	"github.com/andreyxaxa/posts_comments_service/internal/service"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
	_ "github.com/lib/pq"
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

	logger.Info.Print("Connecting to Postgres")
	options := db.PostgresOptions{
		Name:     os.Getenv("POSTGRES_DBNAME"),
		Post:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	logger.Info.Print(options)

	postgresDB, err := db.NewPostgresDB(options)
	if err != nil {
		logger.Error.Fatalf(err.Error())
	}

	var gateways *gateway.Gateways

	logger.Info.Print("Creating Gateways")
	logger.Info.Print("USE_IN_MEMORY = ", os.Getenv("USE_IN_MEMORY"))
	if os.Getenv("USE_IN_MEMORY") == "true" {
		// TODO: inmemory strg
	} else {
		posts := postgres.NewPostsPostgres(postgresDB)
		comments := postgres.NewCommentsPostgres(postgresDB)
		gateways = gateway.NewGateways(posts, comments)
	}

	logger.Info.Print("Creating services")
	services := service.NewServices(gateways, logger) // 1

	logger.Info.Print("Creating graphQL server")
	port := os.Getenv("PORT")
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		Posts:    services.Posts,
		Comments: services.Comments,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	logger.Info.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Error.Fatal(http.ListenAndServe(":"+port, nil))
}
