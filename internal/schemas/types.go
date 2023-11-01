package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Id is an alias for the MongoDB ObjectID
type Id = primitive.ObjectID

// Timestamp is an alias for the MongoDB DateTime
type Timestamp = primitive.DateTime

// Cost represents costs as float32
type Cost = float32
