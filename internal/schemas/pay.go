package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PayItemCollection = "pay_items"
const PayEntryCollection = "pay_entries"
const PayPeriodCollection = "pay_periods"

type PayItem struct {
	Id          primitive.ObjectID             `bson:"_id"`
	Name        string                         `bson:"name"`
	Description string                         `bson:"description"`
	CostPerUnit float32                        `bson:"cost_per_piece"`
	Quantity    float32                        `bson:"quantity"`
	UserIdToDue map[primitive.ObjectID]float32 `bson:"person_to_dues"`
}

type PayEntry struct {
	Id          primitive.ObjectID             `bson:"_id"`
	Time        time.Time                      `bson:"time"`
	Location    string                         `bson:"location"`
	Description string                         `bson:"description"`
	Amount      float32                        `bson:"amount"`
	PayerId     primitive.ObjectID             `bson:"payer_id"`
	ItemIds     []primitive.ObjectID           `bson:"item_ids"`
	UserIdToDue map[primitive.ObjectID]float32 `bson:"person_to_dues"`
}

type PayPeriod struct {
	Id          primitive.ObjectID             `bson:"_id"`
	Start       time.Time                      `bson:"start"`
	End         time.Time                      `bson:"end,omitempty"`
	Completed   bool                           `bson:"completed"`
	EntryIds    []primitive.ObjectID           `bson:"entry_ids"`
	UserIdToDue map[primitive.ObjectID]float32 `bson:"person_to_dues"`
}
