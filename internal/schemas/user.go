package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollection = "users"

type User struct {
	Id       primitive.ObjectID   `bson:"_id" json:"_id"`
	Name     string               `bson:"name" json:"name"`
	HouseIds []primitive.ObjectID `bson:"houses" json:"houses"`
}
