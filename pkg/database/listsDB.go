package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"restapi/pkg/data"
	"restapi/pkg/logger"
)

//Структура для работы со списками из БД

func CheckList(userId string, listID string) bool {
	var result data.ListWithId
	collection := clientGlobal.Database(DatabaseName).Collection("lists")
	listIdBS, _ := primitive.ObjectIDFromHex(listID)
	filter := bson.D{{"owner_id", userId}, {"_id", listIdBS}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}

func GetAllListsDB(userId string) []data.ListWithId {
	var ListsResult []data.ListWithId
	collection := clientGlobal.Database(DatabaseName).Collection("lists")

	filter := bson.D{{"owner_id", userId}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		logger.Logger.Error(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.ListWithId
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		ListsResult = append(ListsResult, elem)
	}

	return ListsResult
}

func GetListDB(params map[string]string, userId string) (data.ListWithId, error) {

	var result data.ListWithId

	collection := clientGlobal.Database(DatabaseName).Collection("lists")
	var objectID, _ = primitive.ObjectIDFromHex(params["id"])
	filter := bson.D{{"owner_id", userId}, {"_id", objectID}}
	collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, nil
}

func CreateListDB(list data.ListWithoutId) (interface{}, error) {

	collection := clientGlobal.Database(DatabaseName).Collection("lists")
	result, err := collection.InsertOne(context.TODO(), list)
	if err != nil {
		return nil, err
	} else {
		return result.InsertedID, nil
	}

}

func UpdateListDB(list data.ListWithoutId, userID string, listID string) (bool, error) {
	var objectID, _ = primitive.ObjectIDFromHex(listID)
	filter := bson.D{{"owner_id", userID}, {"_id", objectID}}
	newListBson := bson.D{
		{"$set", bson.D{
			{"title", list.Title},
			{"description", list.Description},
		}},
	}

	collection := clientGlobal.Database(DatabaseName).Collection("lists")
	_, err := collection.UpdateOne(context.TODO(), filter, newListBson)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func DeleteListDB(params map[string]string, userId string) (bool, error) {

	collection := clientGlobal.Database(DatabaseName).Collection("lists")
	var objectID, _ = primitive.ObjectIDFromHex(params["id"])
	filter := bson.D{{"owner_id", userId}, {"_id", objectID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.Logger.Error("Error Delete List ")
		return false, err
	}
	_, err = DeleteItemsFromList(params["id"])
	if err != nil {
		logger.Logger.Error("Error Delete List Items")
		return false, err
	}
	return true, err
}

func DeleteItemsFromList(listID string) (*mongo.DeleteResult, error) {

	collection := clientGlobal.Database(DatabaseName).Collection("items")

	filter := bson.D{{"list_id", listID}}
	result, err := collection.DeleteMany(context.TODO(), filter)

	return result, err
}
