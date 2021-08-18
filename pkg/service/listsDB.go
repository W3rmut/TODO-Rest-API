package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//Структура для работы со списками из БД

type ListInput struct {
	Id          string `bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `bson:"owner_id"`
}

type ListOutput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `bson:"owner_id"`
}

func CheckList(userId string, listID string) bool {
	var result ListInput
	collection := clientGlobal.Database("test").Collection("lists")
	listIdBS, _ := primitive.ObjectIDFromHex(listID)
	filter := bson.D{{"owner_id", userId}, {"_id", listIdBS}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}

func GetAllListsDB(userId string) []ListInput {
	var ListsResult []ListInput
	collection := clientGlobal.Database("test").Collection("lists")

	filter := bson.D{{"owner_id", userId}}

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var elem ListInput
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		ListsResult = append(ListsResult, elem)
	}
	return ListsResult
}

func GetListDB(params map[string]string, userId string) (ListInput, error) {

	var result ListInput

	collection := clientGlobal.Database("test").Collection("lists")
	var objectID, _ = primitive.ObjectIDFromHex(params["id"])
	filter := bson.D{{"owner_id", userId}, {"_id", objectID}}
	collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, nil
}

func CreateListDB(list ListOutput) int {
	collection := clientGlobal.Database("test").Collection("lists")
	result, err := collection.InsertOne(context.TODO(), list)
	fmt.Println(result, err)
	return 1
}

func UpdateListDB(list ListOutput, userID string, listID string) int {
	var objectID, _ = primitive.ObjectIDFromHex(listID)
	fmt.Println(objectID)
	filter := bson.D{{"owner_id", userID}, {"_id", objectID}}
	newListBson := bson.D{
		{"$set", bson.D{
			{"title", list.Title},
			{"description", list.Description},
		}},
	}

	fmt.Println(newListBson)
	collection := clientGlobal.Database("test").Collection("lists")
	result, err := collection.UpdateOne(context.TODO(), filter, newListBson)
	fmt.Println(result, err)
	return 1
}

func DeleteListDB(params map[string]string, userId string) error {

	collection := clientGlobal.Database("test").Collection("lists")
	var objectID, _ = primitive.ObjectIDFromHex(params["id"])
	filter := bson.D{{"owner_id", userId}, {"_id", objectID}}
	collection.DeleteOne(context.TODO(), filter)

	return nil
}
