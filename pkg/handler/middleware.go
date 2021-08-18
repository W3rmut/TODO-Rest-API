package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/pkg/service"
	"strings"
)

func CheckTokenHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		splitToken := strings.Split(auth, "Bearer ")
		auth = splitToken[1]
		fmt.Println(auth)
		resultBool, name, err := checkToken(auth, []byte(key))
		fmt.Println(resultBool)
		if resultBool {
			json.NewEncoder(w).Encode("token valid")
			next.ServeHTTP(w, r)
		} else {
			json.NewEncoder(w).Encode("token invalid")
			return
		}
		fmt.Println(resultBool, name, err)

	})
}

func CheckListMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := GetUserId(r)
		listID := mux.Vars(r)["idList"]
		fmt.Println(listID)
		check := service.CheckList(userID, listID)

		if check {
			next.ServeHTTP(w, r)
		} else {
			json.NewEncoder(w).Encode("list not found")
			return
		}

	})
}
