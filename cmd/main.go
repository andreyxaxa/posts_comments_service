package main

import (
	"os"

	"github.com/andreyxaxa/posts_comments_service/configs"
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
}
