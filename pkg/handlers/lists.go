package handlers

import (
	"C"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/pkg/data"
	listsdb "restapi/pkg/database"
	"restapi/pkg/jwt"
	"restapi/pkg/logger"
)

//Swagger tags

func GetAllLists(w http.ResponseWriter, r *http.Request) {

	idUser, errName := jwt.GetUserId(r)

	if errName != nil {
		logger.Logger.Errorf("Error parse lists")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error parse data"})
	} else {
		var result []data.ListWithId
		result = listsdb.GetAllListsDB(idUser)
		json.NewEncoder(w).Encode(result)
	}

}

func GetList(w http.ResponseWriter, r *http.Request) {
	idUser, _ := jwt.GetUserId(r)
	result, err := listsdb.GetListDB(mux.Vars(r), idUser)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "List not found"})
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func CreateList(w http.ResponseWriter, r *http.Request) {
	var newList data.ListWithoutId
	idUser, _ := jwt.GetUserId(r)
	err := json.NewDecoder(r.Body).Decode(&newList)
	newList.OwnerId = idUser
	result, err := listsdb.CreateListDB(newList)
	if err != nil {
		logger.Logger.Errorf("Error creating list with id \"%s\", Title: \"%s\", Description: \"%s\"", idUser, newList.Title, newList.Description)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error creating list"})
	} else {
		json.NewEncoder(w).Encode(data.ResponseWithID{Id: result})
	}

}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	var newList data.ListWithoutId
	idUser, _ := jwt.GetUserId(r)
	if err := json.NewDecoder(r.Body).Decode(&newList); err != nil {
	}
	listIDVar := mux.Vars(r)
	listID := listIDVar["id"]

	result, err := listsdb.UpdateListDB(newList, idUser, listID)
	if result != true {
		logger.Logger.Errorf("Error update list with id \"%s\", new Title: \"%s\", new Description: \"%s\", err: \"%e\"", listID, newList.Title, newList.Description, err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error Updating list"})
	} else {
		json.NewEncoder(w).Encode(data.StatusResponse{Result: result})
	}

}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	idUser, _ := jwt.GetUserId(r)
	result, err := listsdb.DeleteListDB(mux.Vars(r), idUser)
	if result != true {
		logger.Logger.Errorf("Error delete List with id %s, error: %e", mux.Vars(r)["id"], err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error Delete list"})

	} else {
		json.NewEncoder(w).Encode(data.StatusResponse{Result: result})
	}

}
