package data

type ListWithId struct {
	Id          string `bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `bson:"owner_id"`
}

type ListWithoutId struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `bson:"owner_id"`
}
