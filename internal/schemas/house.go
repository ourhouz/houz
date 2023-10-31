package schemas

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const HouseCollection = "houses"

type House struct {
	Id         id          `bson:"_id" json:"_id"`
	Name       string      `bson:"name" json:"name"`
	Owner      id          `bson:"owner_id" json:"owner_id"`
	UserIds    []id        `bson:"user_ids,omitempty" json:"user_ids,omitempty"`
	PayPeriods []PayPeriod `bson:"pay_periods,omitempty" json:"pay_periods,omitempty"`
}

// NewHouse creates a new house with the given name and owner ID
func NewHouse(name, owner string) (House, error) {
	if name == "" {
		return House{}, errors.New("house name is required")
	}
	if owner == "" {
		return House{}, errors.New("owner ID is required")
	}

	ownerId, err := primitive.ObjectIDFromHex(owner)
	if err != nil {
		return House{}, err
	}

	return House{
		Name:  name,
		Owner: ownerId,
		UserIds: []id{
			ownerId,
		},
		PayPeriods: []PayPeriod{},
	}, nil
}
