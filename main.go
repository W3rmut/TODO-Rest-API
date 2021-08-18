package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	handlers "restapi/pkg/handler"
	"restapi/pkg/service"
)

func main() {

	r := mux.NewRouter()

	//Create database connect
	err := service.Connect()
	if err != nil {
		panic("Error DB connection")
	}

	//Users options
	r.HandleFunc("/sign-up", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/sign-in", handlers.AuthUser).Methods("POST")

	//Lists options

	//GetAllLists
	GetAllLists := http.HandlerFunc(handlers.GetAllLists)
	r.Handle("/lists", handlers.CheckTokenHandler(GetAllLists)).Methods("GET")

	//Get one list
	GetOneList := http.HandlerFunc(handlers.GetList)
	r.Handle("/lists/{id}", handlers.CheckTokenHandler(GetOneList)).Methods("GET")

	//Create List
	CreateList := http.HandlerFunc(handlers.CreateList)
	r.Handle("/lists", handlers.CheckTokenHandler(CreateList)).Methods("POST")

	//Update List
	UpdateList := http.HandlerFunc(handlers.UpdateList)
	r.Handle("/lists/{id}", handlers.CheckTokenHandler(UpdateList)).Methods("PUT")

	//DeleteList
	DeleteList := http.HandlerFunc(handlers.DeleteList)
	r.Handle("/lists/{id}", DeleteList).Methods("DELETE")

	//Items options

	//Get all items in list handler
	GetAllItems := http.HandlerFunc(handlers.GetAllItemsInList)
	r.Handle("/lists/{idList}/items", handlers.CheckTokenHandler(handlers.CheckListMiddleware(GetAllItems))).Methods("GET")

	//Add item in list handler
	AddItem := http.HandlerFunc(handlers.AddItemInList)
	r.Handle("/lists/{idList}/items", handlers.CheckTokenHandler(handlers.CheckListMiddleware(AddItem))).Methods("POST")

	//Get one item in list handler
	GetOneItem := http.HandlerFunc(handlers.GetItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CheckTokenHandler(handlers.CheckListMiddleware(GetOneItem))).Methods("GET")

	//Update one item in list handler
	UpdateItem := http.HandlerFunc(handlers.UpdateItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CheckTokenHandler(handlers.CheckListMiddleware(UpdateItem))).Methods("PUT")

	//Delete one item in list handler
	DeleteItem := http.HandlerFunc(handlers.DeleteItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CheckTokenHandler(handlers.CheckListMiddleware(DeleteItem))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
