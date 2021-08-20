package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ItemInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	ListID      string `bson:"list_id"`
}

type ItemOutput struct {
	Id          string `bson:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	ListID      string `bson:"list_id"`
}

func GetAllItemsDB(listID string) ([]ItemOutput, error) {

	var result []ItemOutput
	collection := clientGlobal.Database("test").Collection("items")
	//listIdBS,_ := primitive.ObjectIDFromHex(listID)
	filter := bson.D{{"list_id", listID}}
	cur, err := collection.Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		var elem ItemOutput
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

func GetItemBD(listID string, itemID string) (ItemOutput, error) {
	fmt.Println("//DB start:\n\n")
	var result ItemOutput
	collection := clientGlobal.Database("test").Collection("items")
	itemIdBS, _ := primitive.ObjectIDFromHex(itemID)
	fmt.Println(listID, itemID)
	filter := bson.D{{"list_id", listID}, {"_id", itemIdBS}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(result)
	if err != nil {
		return result, fmt.Errorf("Element not found")
	}

	return result, nil

}

func AddItemBD(listID string, input ItemInput) (interface{}, error) {
	input.ListID = listID
	collection := clientGlobal.Database("test").Collection("items")
	resultID, err := collection.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("Error create item")
	}
	return resultID, nil
}

func UpdateItemDB(listID string, itemID string, newItem ItemInput) (interface{}, error) {

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
	result, err := collection.UpdateOne(context.TODO(), filter, newItemBson)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteItemDB(listID, itemID string) (interface{}, error) {

	collection := clientGlobal.Database("test").Collection("items")
	var objectID, _ = primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"list_id", listID}, {"_id", objectID}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
