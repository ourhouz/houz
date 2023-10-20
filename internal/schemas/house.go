package schemas

const HouseCollection = "houses"

type House struct {
	Id      docId   `bson:"_id"`
	Name    string  `bson:"name"`
	UserIds []docId `bson:"user_ids"`
}
