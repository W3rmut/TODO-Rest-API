package data

type ItemWithoutId struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	ListID      string `bson:"list_id"`
}

type ItemWithId struct {
	Id          string `bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	ListID      string `bson:"list_id"`
}
