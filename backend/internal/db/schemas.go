package db

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Base struct {
	Timestamps
	ID uint `gorm:"primaryKey" json:"id"`
}

type House struct {
	Base
	Name       string      `gorm:"size:100" json:"name"`
	Owner      User        `json:"owner"`
	Housemates []User      `json:"housemates"`
	PayPeriods []PayPeriod `json:"payPeriods"`
}

func CreateHouse(name string) (h House, err error) {
	if len(name) == 0 {
		err = errors.New("house name cannot be empty")
		return
	}
	if len(name) > 100 {
		err = errors.New("house name cannot be longer than 100 characters")
		return
	}
	return House{
		Name: name,
	}, nil
}

type User struct {
	Timestamps
	HouseID      uint   `gorm:"primaryKey" json:"houseID"`
	Name         string `gorm:"primaryKey;size:100" json:"name"`
	PasswordHash []byte `json:"-"`
}

func CreateUser(houseID uint, name, password string) (u User, err error) {
	if houseID == 0 {
		err = errors.New("houseID cannot be 0")
		return
	}
	if len(name) == 0 {
		err = errors.New("name cannot be empty")
		return
	}
	if len(name) > 100 {
		err = errors.New("name cannot be longer than 100 characters")
		return
	}

	if len(password) == 0 {
		err = errors.New("password cannot be empty")
		return
	}
	if len(password) > 72 {
		// bcrypt limit
		err = errors.New("password cannot be longer than 72 characters")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return User{
		HouseID:      houseID,
		Name:         name,
		PasswordHash: hash,
	}, nil
}

type PayPeriod struct {
	Base
	HouseID    uint       `json:"houseID"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    time.Time  `json:"endTime"`
	Completed  bool       `json:"completed"`
	PayEntries []PayEntry `json:"payEntries"`
}

func CreatePayPeriod(houseID uint, startTime time.Time) (pp PayPeriod, err error) {
	if houseID == 0 {
		err = errors.New("houseID cannot be 0")
		return
	}
	if startTime.IsZero() {
		err = errors.New("start time cannot be empty")
		return
	}
	return PayPeriod{
		HouseID:   houseID,
		StartTime: startTime,
	}, nil
}

type PayEntry struct {
	Base
	PayPeriodID uint      `json:"payPeriodID"`
	Location    string    `gorm:"size:100" json:"location"`
	Description string    `gorm:"size:500" json:"description"`
	TotalCost   float32   `json:"totalCost"`
	PayItems    []PayItem `json:"payItems"`
}

type PayItem struct {
	Base
	PayEntryID  uint         `json:"payEntryID"`
	Name        string       `gorm:"size:100" json:"name"`
	Description string       `gorm:"size:500" json:"description"`
	CostPerUnit float32      `json:"costPerUnit"`
	Quantity    float32      `json:"quantity"`
	PayItemDues []PayItemDue `json:"payItemDues"`
}

type PayItemDue struct {
	Base
	PayItemID uint    `json:"payItemID"`
	AmountDue float32 `json:"amountDue"`
	UserId    uint    `json:"userID"` // many-to-one?
}
