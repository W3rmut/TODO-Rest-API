package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/pkg/data"
	dbusers "restapi/pkg/database"
	"restapi/pkg/jwt"
	"restapi/pkg/logger"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user data.UserWithoutId
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data.ErrorResponse{Error: "error creating user"})
		return
	}

	usernameCheck := dbusers.CheckUsername(user.Name)
	if usernameCheck {
		_, err := dbusers.CreateUser(user)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Error creating user"})
			return
		} else {
			json.NewEncoder(w).Encode(data.StatusResponse{Result: true})
			return
		}

	} else {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(data.ErrorResponse{Error: "username used"})
		return
	}

}

func AuthUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user data.UserWithId
	err := json.NewDecoder(r.Body).Decode(&user)
	user, errAuth := dbusers.AuthUser(user)
	if errAuth != nil {
		var errorResponse = ErrorResponse{Error: "user not found"}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		token, errToken := jwt.CreateToken(user)
		if err != nil {
			logger.Logger.Errorf("Error auth user %s , error: %e", user.Name, errToken)
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Error auth"})
		} else {
			json.NewEncoder(w).Encode(data.AuthorizationResponse{Token: token})
		}

	}

}
