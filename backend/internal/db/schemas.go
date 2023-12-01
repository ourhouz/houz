package db

import (
	"time"
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

type User struct {
	Timestamps
	HouseID      uint   `gorm:"primaryKey" json:"houseID"`
	Name         string `gorm:"primaryKey;size:100" json:"name"`
	PasswordHash []byte `json:"-"`
}

type PayPeriod struct {
	Base
	HouseID    uint       `json:"houseID"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    time.Time  `json:"endTime"`
	Completed  bool       `json:"completed"`
	PayEntries []PayEntry `json:"payEntries"`
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
	PayEntryID  uint    `json:"payEntryID"`
	Name        string  `gorm:"size:100" json:"name"`
	Description string  `gorm:"size:500" json:"description"`
	CostPerUnit float32 `json:"costPerUnit"`
	Quantity    float32 `json:"quantity"`
}

type PayItemDue struct {
	Base
	PayItemID uint    `json:"payItemID"`
	UserId    uint    `json:"userID"`
	AmountDue float32 `json:"amountDue"`
}
