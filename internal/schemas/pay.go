package schemas

// represent costs as float32
type cost = float32

const PayItemCollection = "pay_items"
const PayEntryCollection = "pay_entries"
const PayPeriodCollection = "pay_periods"

type PayItem struct {
	Id          id          `bson:"_id" json:"_id"`
	Name        string      `bson:"name" json:"name"`
	Description string      `bson:"description" json:"description"`
	CostPerUnit cost        `bson:"cost_per_piece" json:"cost_per_piece"`
	Quantity    cost        `bson:"quantity" json:"quantity"`
	UserIdToDue map[id]cost `bson:"person_to_dues" json:"person_to_dues"`
}

type PayEntry struct {
	Id          id          `bson:"_id" json:"_id"`
	Time        time        `bson:"time" json:"time"`
	Location    string      `bson:"location" json:"location"`
	Description string      `bson:"description,omitempty" json:"description,omitempty"`
	Amount      cost        `bson:"amount" json:"amount"`
	PayerId     id          `bson:"payer_id" json:"payer_id"`
	Items       []PayItem   `bson:"item_ids" json:"item_ids"`
	UserIdToDue map[id]cost `bson:"person_to_dues" json:"person_to_dues"`
}

type PayPeriod struct {
	Id          id          `bson:"_id" json:"_id"`
	Start       time        `bson:"start" json:"start"`
	End         time        `bson:"end,omitempty" json:"end,omitempty"`
	Completed   bool        `bson:"completed" json:"completed"`
	Entries     []PayEntry  `bson:"entry_ids" json:"entry_ids"`
	UserIdToDue map[id]cost `bson:"person_to_dues" json:"person_to_dues"`
}
