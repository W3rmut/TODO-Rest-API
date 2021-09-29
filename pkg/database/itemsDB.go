package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"restapi/pkg/data"
)

func GetAllItemsDB(listID string) ([]data.ItemWithId, error) {

	var result []data.ItemWithId
	collection := clientGlobal.Database("test").Collection("items")
	filter := bson.D{{"list_id", listID}}
	cur, err := collection.Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		var elem data.ItemWithId
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, elem)
	}

	if err != nil {
		return result, fmt.Errorf("No have list for this user")
	} else {
		return result, nil
	}

}

func GetItemBD(listID string, itemID string) (data.ItemWithId, error) {

	var result data.ItemWithId
	collection := clientGlobal.Database("test").Collection("items")
	itemIdBS, _ := primitive.ObjectIDFromHex(itemID)

	filter := bson.D{{"list_id", listID}, {"_id", itemIdBS}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, fmt.Errorf("Element not found")
	}

	return result, nil

}

func AddItemBD(listID string, input data.ItemWithoutId) (interface{}, error) {
	input.ListID = listID
	collection := clientGlobal.Database("test").Collection("items")
	resultID, err := collection.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("Error create item")
	}
	return resultID.InsertedID, nil
}

func UpdateItemDB(listID string, itemID string, newItem data.ItemWithoutId) (bool, error) {

	var objectID, _ = primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"list_id", listID}, {"_id", objectID}}

	newItemBson := bson.D{
		{"$set", bson.D{
			{"title", newItem.Title},
			{"description", newItem.Description},
			{"done", newItem.Done},
			{"list_id", listID},
		}},
	}
	collection := clientGlobal.Database("test").Collection("items")
	_, err := collection.UpdateOne(context.TODO(), filter, newItemBson)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteItemDB(listID, itemID string) (bool, error) {

	collection := clientGlobal.Database("test").Collection("items")
	var objectID, _ = primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"list_id", listID}, {"_id", objectID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}
