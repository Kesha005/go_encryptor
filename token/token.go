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
	Token string
}


func returnSecret() []byte {
	token := []byte(os.Getenv("SECRET_KEY"))
	return token
}

func (user UserToken)GenerateToken() (string, error) {
	var secret = returnSecret()
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
	var secret =  returnSecret()
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

func (tokenString JWT)GetTokenData()(UserToken,error){
	var secret =  returnSecret()
	token, err := jwt.Parse(tokenString.Token, func(token *jwt.Token) (interface{}, error) {
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
		id := claims["id"].(float64)
		return UserToken{Id: int(id), Username: username}, nil
	}
	return UserToken{},errors.New("It is invalid token")
}

func (tokenString JWT)RefreshToken()(string, error){
	var secret = returnSecret()
	token, err := jwt.Parse(tokenString.Token, func(token *jwt.Token)(interface{},error){
		return secret,nil
	})
	if err!=nil{
		return "",err
	}
	if !token.Valid{
		data,err := tokenString.GetTokenData()
		if err!=nil{
			panic(err)
		}
		newtoken_data := UserToken{Id: data.Id, Username: data.Username}
		refreshed,token_err := newtoken_data.GenerateToken()
		if token_err!=nil{
			panic(token_err)
		}
		return refreshed, nil
	}
	return tokenString.Token, nil
}


