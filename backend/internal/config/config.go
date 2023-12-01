package config

import (
	"github.com/joho/godotenv"

	"os"
)

type Config struct {
	Port        string
	PostgresDSN string
	JWTKey      []byte
}

var Env Config

func Load() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	Env = Config{
		Port:        os.Getenv("PORT"),
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		JWTKey:      []byte(os.Getenv("JWT_KEY")),
	}
}
