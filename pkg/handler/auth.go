package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	dbusers "restapi/pkg/service"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user dbusers.UserOutput
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	usernameCheck := dbusers.CheckUsername(user.Name)
	if usernameCheck {
		result, err := dbusers.CreateUser(user)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
		fmt.Println(user, err)
	} else {
		json.NewEncoder(w).Encode("username used")
	}

}

func AuthUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user dbusers.UserInput
	fmt.Println(user)
	err := json.NewDecoder(r.Body).Decode(&user)
	user, errAuth := dbusers.AuthUser(user)
	if errAuth != nil {
		json.NewEncoder(w).Encode("user not found")
	} else {
		fmt.Println(user.Name)
		token, errToken := createToken(user)
		if err != nil {
			panic(errToken)
		}
		json.NewEncoder(w).Encode(token)
		fmt.Println(user, err)
	}

}
