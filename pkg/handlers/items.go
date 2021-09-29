package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/pkg/data"
	items "restapi/pkg/database"
	"restapi/pkg/logger"
)

func GetAllItemsInList(w http.ResponseWriter, r *http.Request) {
	var Items []data.ItemWithId
	listID := mux.Vars(r)["idList"]
	Items, err := items.GetAllItemsDB(listID)

	if err != nil {
		logger.Logger.Errorf("Error parse items")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error parse data"})
	} else {
		json.NewEncoder(w).Encode(Items)
	}

}

func GetItem(w http.ResponseWriter, r *http.Request) {
	var Item data.ItemWithId
	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]
	Item, err := items.GetItemBD(listID, ItemID)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Item not found"})
	} else {
		json.NewEncoder(w).Encode(Item)
	}

}

func AddItemInList(w http.ResponseWriter, r *http.Request) {
	var Item data.ItemWithoutId
	listID := mux.Vars(r)["idList"]
	json.NewDecoder(r.Body).Decode(&Item)
	result, err := items.AddItemBD(listID, Item)

	if err != nil {
		logger.Logger.Errorf("Error creating item")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error creating Item"})

	} else {
		json.NewEncoder(w).Encode(data.ResponseWithID{Id: result})
	}

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var newItem data.ItemWithoutId
	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]
	json.NewDecoder(r.Body).Decode(&newItem)
	result, err := items.UpdateItemDB(listID, ItemID, newItem)

	if err != nil {
		logger.Logger.Errorf("Error creating item")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error updating Item"})
	} else {
		json.NewEncoder(w).Encode(data.StatusResponse{Result: result})
	}

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]

	result, err := items.DeleteItemDB(listID, ItemID)
	if err != nil {
		logger.Logger.Errorf("Error creating item")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error updating Item"})
	} else {
		json.NewEncoder(w).Encode(data.StatusResponse{Result: result})
	}

}
