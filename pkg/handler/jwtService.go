package handler

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"restapi/pkg/config"
	"restapi/pkg/service"
	"strings"
	"time"
)

var key string

func UpdateKey() {
	key = config.ResultConfig.JwtConfig.JwtKey
}

type jwtStruct struct {
	jwt.StandardClaims
	Id   string
	Name string
}

func GetUserId(request *http.Request) (string, error) {

	tokenString := request.Header.Get("Authorization")
	splitToken := strings.Split(tokenString, "Bearer ")
	tokenString = splitToken[1]
	claims := &jwtStruct{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	fmt.Println("parseerror: ", err)
	if err != nil {
		return "", fmt.Errorf("Can't read the token")
	} else {
		return claims.Id, nil
	}

}

func checkToken(acessToken string, signingKey []byte) (bool, string, error) {

	token, err := jwt.ParseWithClaims(acessToken, &jwtStruct{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return signingKey, nil
		}
	})

	if err != nil {
		return false, "", fmt.Errorf("Invalid token")
	}

	if claims, claimsOk := token.Claims.(*jwtStruct); claimsOk && token.Valid {
		fmt.Println()
		return true, claims.Name, nil
	} else {
		return false, "", errors.New("Invalid struct token")
	}

}

func createToken(user service.UserInput) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtStruct{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
		user.Name})

	return token.SignedString([]byte(key))
}
