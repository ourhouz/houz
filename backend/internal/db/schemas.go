package db

import (
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	Name       string `gorm:"size:100" json:"name"`
	Owner      User   `gorm:"foreignKey:HouseId" json:"owner"`
	Housemates []User `gorm:"foreignKey:HouseId" json:"housemates"`
}

type User struct {
	gorm.Model
	HouseId      uint   `gorm:"primaryKey"`
	Name         string `gorm:"primaryKey;size:100"`
	PasswordHash []byte `json:"-"`
}
