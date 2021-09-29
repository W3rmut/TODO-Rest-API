package database

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"restapi/pkg/data"
)

//This file use for BD connections and other actions with database
const salt = "asasd12312as12w"

func PasswordHash(password string) string {
	passwordHash := sha1.New()
	passwordHash.Write([]byte(password))
	return fmt.Sprintf("%x", passwordHash.Sum([]byte(salt)))
}

func CheckUsername(username string) bool {
	creds := bson.D{{"name", username}}

	var resultUser data.UserWithId
	collection := clientGlobal.Database("test").Collection("users")
	collection.FindOne(context.TODO(), creds).Decode(&resultUser)
	if resultUser.Name == "" {
		return true
	} else {
		return false
	}
}

func CreateUser(user data.UserWithoutId) (interface{}, error) {

	user.PasswordHash = PasswordHash(user.PasswordHash)
	resultUser := data.UserWithoutId{Name: user.Name, Email: user.Email, PasswordHash: user.PasswordHash}
	collection := clientGlobal.Database("test").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), resultUser)
	if err != nil {
		return 0, errors.New("Error creating user: ")
	} else {
		return insertResult.InsertedID, nil
	}

}

//Выполняет авторизацию. Возвращает токен авторизации или ошибку
func AuthUser(user data.UserWithId) (data.UserWithId, error) {
	user.PasswordHash = PasswordHash(user.PasswordHash)
	fmt.Println(user.PasswordHash)
	filter := bson.D{{"name", user.Name}, {"passwordhash", user.PasswordHash}}
	var resultUser data.UserWithId
	collection := clientGlobal.Database("test").Collection("users")
	collection.FindOne(context.TODO(), filter).Decode(&resultUser)
	if resultUser.Name == "" {
		return resultUser, errors.New("User not found")
	} else {
		return resultUser, nil
	}

}
