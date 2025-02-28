package configs

import "github.com/joho/godotenv"

func InitConfigs() error {
	err := godotenv.Load()

	return err
}
