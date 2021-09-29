package data

type UserWithId struct {
	Id           string `bson:"_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserWithoutId struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type AuthorizationResponse struct {
	Token string
}
