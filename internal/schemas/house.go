package schemas

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const HouseCollection = "houses"

type House struct {
	Id         Id          `json:"_id"`
	Name       string      `json:"name"`
	Owner      Id          `json:"owner_id"`
	UserIds    []Id        `json:"user_ids"`
	PayPeriods []PayPeriod `json:"pay_periods"`
}

// NewHouse creates a new house with the given name and owner ID
func NewHouse(name string, owner Id) (House, error) {
	if name == "" {
		return House{}, errors.New("house name is required")
	}

	return House{
		Id:    primitive.NewObjectID(),
		Name:  name,
		Owner: owner,
		UserIds: []Id{
			owner,
		},
		PayPeriods: []PayPeriod{},
	}, nil
}
