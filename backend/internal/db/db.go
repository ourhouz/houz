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
	models := []interface{}{
		&User{},
		&House{},
		&PayPeriod{},
		&PayEntry{},
		&PayItem{},
		&PayItemDue{},
	}

	for _, model := range models {
		err := Database.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}
}
