package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"restapi/pkg/config"
	"restapi/pkg/data"
	"restapi/pkg/logger"
	"strings"
	"time"
)

var Key string

func UpdateKey() {
	Key = config.ResultConfig.JwtConfig.JwtKey
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
		return []byte(Key), nil
	})
	if err != nil {
		logger.Logger.Error("Error parse token", err)
	}

	if err != nil {
		return "", fmt.Errorf("Can't read the token")
	} else {
		return claims.Id, nil
	}

}

func CheckToken(acessToken string, signingKey []byte) error {

	token, err := jwt.ParseWithClaims(acessToken, &jwtStruct{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v poshel nahui otusda mosheinik", token.Header["alg"])
		} else {
			return signingKey, nil
		}
	})

	if err != nil {
		return fmt.Errorf("Invalid token")
	}

	if _, claimsOk := token.Claims.(*jwtStruct); claimsOk && token.Valid {
		return nil
	} else {
		return fmt.Errorf("Invalid struct token")
	}

}

func CreateToken(user data.UserWithId) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtStruct{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
		user.Name})

	return token.SignedString([]byte(Key))
}
