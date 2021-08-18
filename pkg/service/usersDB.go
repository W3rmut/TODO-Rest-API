package service

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

//This file use for BD connections and other actions with database

type UserInput struct {
	Id           string `bson:"_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserOutput struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func CheckUsername(username string) bool {
	creds := bson.D{{"name", username}}

	var resultUser UserInput
	collection := clientGlobal.Database("test").Collection("users")
	collection.FindOne(context.TODO(), creds).Decode(&resultUser)
	if resultUser.Name == "" {
		return true
	} else {
		return false
	}
}

func CreateUser(user UserOutput) (interface{}, error) {
	resultUser := UserOutput{Name: user.Name, Email: user.Email, PasswordHash: user.PasswordHash}
	collection := clientGlobal.Database("test").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), resultUser)
	if err != nil {
		return 0, errors.New("User not found")
	} else {
		return insertResult.InsertedID, nil
	}

}

//Выполняет авторизацию. Возвращает токен авторизации или ошибку
func AuthUser(user UserInput) (UserInput, error) {
	creds := bson.D{{"name", user.Name}, {"password_hash", user.PasswordHash}}
	var resultUser UserInput
	collection := clientGlobal.Database("test").Collection("users")
	collection.FindOne(context.TODO(), creds).Decode(&resultUser)
	fmt.Println(resultUser.Name)
	if resultUser.Name == "" {
		return resultUser, errors.New("User not found")
	} else {
		fmt.Println(resultUser)
		return resultUser, nil
	}

}
