package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollection = "users"

type User struct {
	Id       primitive.ObjectID   `bson:"_id"`
	Name     string               `bson:"name"`
	HouseIds []primitive.ObjectID `bson:"houses"`
}
