package config

import (
	"github.com/joho/godotenv"

	"os"
)

type Config struct {
	Port         string
	MongoURI     string
	MongoCluster string
}

var Env Config

func Load() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	Env = Config{
		Port:         os.Getenv("PORT"),
		MongoURI:     os.Getenv("MONGO_URI"),
		MongoCluster: os.Getenv("MONGO_CLUSTER"),
	}
}
