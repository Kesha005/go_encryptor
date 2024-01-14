package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type UserToken struct {
	Id int
	Username string //it maybe email
}


type JWT struct{
	token string
	exp time.Time
}


func ReturnSecret() []byte {
	token := []byte(os.Getenv("SECRET_KEY"))
	return token
}

func (user UserToken)GenerateToken() (string, error) {
	var secret = ReturnSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
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

func GetTokenData(tokenString string)(UserToken,error){
	var secret =  ReturnSecret()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return UserToken{}, err
	}
	if !token.Valid {
		return UserToken{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["username"].(string)
		id := claims["id"].(int)
		return UserToken{Id: id, Username: username}, nil
	}
	return UserToken{},errors.New("It is invalid token")
}

