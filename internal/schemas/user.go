package schemas

const UserCollection = "users"

type User struct {
	Id       docId   `bson:"_id"`
	Name     string  `bson:"name"`
	HouseIds []docId `bson:"houses"`
}
