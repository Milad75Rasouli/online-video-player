package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Password string
}

func (c *Config) Read() error {
	var (
		err error
	)
	err = godotenv.Load()
	if err != nil {
		return err
	}
	c.Password = os.Getenv("USER_PASSWORD")
	return nil
}