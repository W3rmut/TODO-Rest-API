package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/pkg/database"
	"restapi/pkg/jwt"
	"strings"
)

type ErrorResponse struct {
	Error interface{} `json:"error"`
}

func CheckTokenHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		splitToken := strings.Split(auth, "Bearer ")
		if len(splitToken) < 2 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid token"})
			return
		}
		auth = splitToken[1]

		errToken := jwt.CheckToken(auth, []byte(jwt.Key))
		if errToken != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid token"})
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func CheckListMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := jwt.GetUserId(r)
		listID := mux.Vars(r)["idList"]
		check := database.CheckList(userID, listID)

		if check {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "List not found"})
			return
		}

	})
}

func CorsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
