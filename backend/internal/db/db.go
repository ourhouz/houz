package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ourhouz/houz/internal/config"
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

func Init() {
	err := Database.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	err = Database.AutoMigrate(&House{})
	if err != nil {
		panic(err)
	}

}
