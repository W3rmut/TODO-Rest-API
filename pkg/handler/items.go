package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	items "restapi/pkg/service"
)

func GetAllItemsInList(w http.ResponseWriter, r *http.Request) {
	var Items []items.ItemOutput
	listID := mux.Vars(r)["idList"]
	Items, err := items.GetAllItemsDB(listID)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(Items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	var Item items.ItemOutput
	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]
	Item, err := items.GetItemBD(listID, ItemID)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(Item)
}

func AddItemInList(w http.ResponseWriter, r *http.Request) {
	var Item items.ItemInput
	listID := mux.Vars(r)["idList"]
	json.NewDecoder(r.Body).Decode(&Item)
	result, err := items.AddItemBD(listID, Item)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var newItem items.ItemInput
	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]
	json.NewDecoder(r.Body).Decode(&newItem)
	result, err := items.UpdateItemDB(listID, ItemID, newItem)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	listID, ItemID := mux.Vars(r)["idList"], mux.Vars(r)["idItem"]

	result, err := items.DeleteItemDB(listID, ItemID)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(result)

}
