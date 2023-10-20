package schemas

import (
	"time"
)

// represent costs as float32
type cost = float32

const PayItemCollection = "pay_items"
const PayEntryCollection = "pay_entries"
const PayPeriodCollection = "pay_periods"

type PayItem struct {
	Id          docId          `bson:"_id"`
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	CostPerUnit cost           `bson:"cost_per_piece"`
	Quantity    cost           `bson:"quantity"`
	UserIdToDue map[docId]cost `bson:"person_to_dues"`
}

type PayEntry struct {
	Id          docId          `bson:"_id"`
	Time        time.Time      `bson:"time"`
	Location    string         `bson:"location"`
	Description string         `bson:"description"`
	Amount      cost           `bson:"amount"`
	PayerId     docId          `bson:"payer_id"`
	ItemIds     []docId        `bson:"item_ids"`
	UserIdToDue map[docId]cost `bson:"person_to_dues"`
}

type PayPeriod struct {
	Id          docId          `bson:"_id"`
	Start       time.Time      `bson:"start"`
	End         time.Time      `bson:"end,omitempty"`
	Completed   bool           `bson:"completed"`
	EntryIds    []docId        `bson:"entry_ids"`
	UserIdToDue map[docId]cost `bson:"person_to_dues"`
}
