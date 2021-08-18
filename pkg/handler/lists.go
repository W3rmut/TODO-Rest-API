package handler

import (
	"C"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	listsdb "restapi/pkg/service"
)

var UserID = "6112c21bb83141242317ecc7"

func GetAllLists(w http.ResponseWriter, r *http.Request) {

	idUser, errName := GetUserId(r)
	fmt.Println(idUser)
	if errName != nil {
		json.NewEncoder(w).Encode(errName)
	} else {
		var result []listsdb.ListInput
		result = listsdb.GetAllListsDB(idUser)
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
	}

}

func GetList(w http.ResponseWriter, r *http.Request) {
	idUser, _ := GetUserId(r)
	if result, err := listsdb.GetListDB(mux.Vars(r), idUser); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func CreateList(w http.ResponseWriter, r *http.Request) {
	var newList listsdb.ListOutput
	idUser, _ := GetUserId(r)
	if err := json.NewDecoder(r.Body).Decode(&newList); err != nil {
	}
	newList.OwnerId = idUser
	result := listsdb.CreateListDB(newList)
	json.NewEncoder(w).Encode(result)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	var newList listsdb.ListOutput
	idUser, _ := GetUserId(r)
	if err := json.NewDecoder(r.Body).Decode(&newList); err != nil {
	}
	listIDVar := mux.Vars(r)
	listID := listIDVar["id"]
	fmt.Println(listID, idUser)

	result := listsdb.UpdateListDB(newList, UserID, listID)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	idUser, _ := GetUserId(r)
	result := listsdb.DeleteListDB(mux.Vars(r), idUser)
	json.NewEncoder(w).Encode(result)
}
