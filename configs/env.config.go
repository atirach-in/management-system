package configs

import "github.com/joho/godotenv"

func InitDotEnv() *error {
	err := godotenv.Load(".env")
	if err != nil {
		return &err
	}
	return nil
}
