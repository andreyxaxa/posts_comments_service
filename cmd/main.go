package main

import (
	"github.com/andreyxaxa/posts_comments_service/configs"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	logger.Info.Print("exec NewLogger")

	logger.Info.Print("exec InitConfig")
	if err := configs.InitConfigs(); err != nil {
		logger.Error.Fatalf(err.Error())
	}
}
