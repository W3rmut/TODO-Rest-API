package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restapi/pkg/config"
	"restapi/pkg/database"
	handlers "restapi/pkg/handlers"
	"restapi/pkg/jwt"
	"restapi/pkg/logger"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {
	r := mux.NewRouter()

	//Set Configs
	config.ParseConfig("configs/config.toml")

	//set jwt key from struct
	jwt.UpdateKey()

	//Create database connect
	err := database.Connect()
	if err != nil {
		panic("Error DB connection")
	}

	logger.InitLogger()
	logger.Logger.Info("Server successful started")

	//Users options
	r.HandleFunc("/sign-up", handlers.CreateUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/sign-in", handlers.AuthUser).Methods("POST", "OPTIONS")

	//Lists options

	//GetAllLists
	GetAllLists := http.HandlerFunc(handlers.GetAllLists)
	r.Handle("/lists", handlers.CorsMiddleware(handlers.CheckTokenHandler(GetAllLists))).Methods("GET", "OPTIONS")

	//Get one list
	GetOneList := http.HandlerFunc(handlers.GetList)
	r.Handle("/lists/{id}", handlers.CorsMiddleware(handlers.CheckTokenHandler(GetOneList))).Methods("GET", "OPTIONS")

	//Create List
	CreateList := http.HandlerFunc(handlers.CreateList)
	r.Handle("/lists", handlers.CorsMiddleware(handlers.CheckTokenHandler(CreateList))).Methods("POST", "OPTIONS")

	//Update List
	UpdateList := http.HandlerFunc(handlers.UpdateList)
	r.Handle("/lists/{id}", handlers.CorsMiddleware(handlers.CheckTokenHandler(UpdateList))).Methods("PUT", "OPTIONS")

	//DeleteList
	DeleteList := http.HandlerFunc(handlers.DeleteList)
	r.Handle("/lists/{id}", handlers.CorsMiddleware(handlers.CheckTokenHandler(DeleteList))).Methods("DELETE", "OPTIONS")

	//Items options

	//Get all items in list handlers
	GetAllItems := http.HandlerFunc(handlers.GetAllItemsInList)
	r.Handle("/lists/{idList}/items", handlers.CorsMiddleware(handlers.CheckTokenHandler(handlers.CheckListMiddleware(GetAllItems)))).Methods("GET", "OPTIONS")

	//Add item in list handlers
	AddItem := http.HandlerFunc(handlers.AddItemInList)
	r.Handle("/lists/{idList}/items", handlers.CorsMiddleware(handlers.CheckTokenHandler(handlers.CheckListMiddleware(AddItem)))).Methods("POST", "OPTIONS")

	//Get one item in list handlers
	GetOneItem := http.HandlerFunc(handlers.GetItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CorsMiddleware(handlers.CheckTokenHandler(handlers.CheckListMiddleware(GetOneItem)))).Methods("GET", "OPTIONS")

	//Update one item in list handlers
	UpdateItem := http.HandlerFunc(handlers.UpdateItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CorsMiddleware(handlers.CheckTokenHandler(handlers.CheckListMiddleware(UpdateItem)))).Methods("PUT", "OPTIONS")

	//Delete one item in list handlers
	DeleteItem := http.HandlerFunc(handlers.DeleteItem)
	r.Handle("/lists/{idList}/items/{idItem}", handlers.CorsMiddleware(handlers.CheckTokenHandler(handlers.CheckListMiddleware(DeleteItem)))).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(config.ResultConfig.ServerConfig.Port, r))

	return nil
}
