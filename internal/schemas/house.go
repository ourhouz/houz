package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const HouseCollection = "houses"

type House struct {
	Id      primitive.ObjectID   `bson:"_id"`
	Name    string               `bson:"name"`
	UserIds []primitive.ObjectID `bson:"user_ids"`
}
