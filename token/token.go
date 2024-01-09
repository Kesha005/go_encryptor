package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)



var secret = []byte("121212121")

func GenerateToken(id int, username string ) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":id ,
			"username":username,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString ,err:= token.SignedString(secret)
	if err !=nil{
		return "There are some errors", err
	}
	return tokenString, nil

}

func ControlToken(input_token string)(string,error){

}
