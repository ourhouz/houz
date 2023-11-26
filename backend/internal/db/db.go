package db

import (
	"github.com/ourhouz/houz/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB = nil

// Connect initializes the singleton connection instance to the MongoDB database
func Connect() {
	var err error

	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN: config.Env.PostgresDSN,
	}))
	if err != nil {
		panic(err)
	}
}
