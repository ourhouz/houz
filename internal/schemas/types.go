package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// id is an alias for the MongoDB ObjectID
type id = primitive.ObjectID

// time is an alias for the MongoDB DateTime
type time = primitive.DateTime
