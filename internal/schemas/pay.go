package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PayItemCollection = "pay_items"

type PayItem struct {
	Id          Id          `json:"_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CostPerUnit Cost        `json:"cost_per_piece"`
	Quantity    Cost        `json:"quantity"`
	UserIdToDue map[Id]Cost `json:"person_to_dues"`
}

func NewPayItem() PayItem {
	return PayItem{
		Id: primitive.NewObjectID(),
	}
}

type PayEntry struct {
	Id          Id          `json:"_id"`
	Time        Timestamp   `json:"Timestamp"`
	Location    string      `json:"location"`
	Description string      `json:"description,omitempty"`
	Amount      Cost        `json:"amount"`
	PayerId     Id          `json:"payer_id"`
	Items       []Id        `json:"item_ids"`
	UserIdToDue map[Id]Cost `json:"person_to_dues"`
}

func NewPayEntry() PayEntry {
	return PayEntry{
		Id: primitive.NewObjectID(),
	}
}

type PayPeriod struct {
	Id          Id          `json:"_id"`
	Start       Timestamp   `json:"start"`
	End         Timestamp   `json:"end,omitempty"`
	Completed   bool        `json:"completed"`
	EntryIds    []Id        `json:"entry_ids"`
	UserIdToDue map[Id]Cost `json:"person_to_dues"`
}

func NewPayPeriod(start int64) PayPeriod {
	return PayPeriod{
		Id:          primitive.NewObjectID(),
		Start:       Timestamp(start),
		Completed:   false,
		EntryIds:    make([]Id, 0),
		UserIdToDue: make(map[Id]Cost),
	}
}
