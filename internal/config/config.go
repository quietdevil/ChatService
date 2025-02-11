package config

import "github.com/joho/godotenv"

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
