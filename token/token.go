package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type User struct {
	id int
	username string
}

func ReturnSecret() []byte {
	token := []byte(os.Getenv("SECRET_KEY"))
	return token
}

func GenerateToken(id int, username string) (string, error) {
	var secret = ReturnSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "There are some errors", err
	}
	return tokenString, nil

}

func ControlToken(input_token string) (string, error) {
	var secret =  ReturnSecret()
	token, err := jwt.Parse(input_token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return "There are error in token", err
	}
	if !token.Valid {
		return "Invalid token", err
	}

	return "It is ok ", nil
}

func GetTokenData(tokenString string)(User,error){
	var secret =  ReturnSecret()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return User{}, err
	}
	if !token.Valid {
		return User{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["username"].(string)
		id := claims["id"].(int)
		return User{id: id, username: username}, nil
	}
	return User{},errors.New("It is invalid token")
}

