package configs

import "github.com/joho/godotenv"

func InitConfigs(envFile string) error {
	err := godotenv.Load(envFile)

	return err
}
