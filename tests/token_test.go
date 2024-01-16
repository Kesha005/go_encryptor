package main

import (
	"fmt"
	"testing"

	"github.com/Kesha005/go_encryptor/token"
	"github.com/joho/godotenv"
)



func TestToken(t *testing.T){
	godotenv.Load("../.env")
	user := token.UserToken{2,"Kerimberdi"}
	

	stringToken, err := user.GenerateToken()
	if err !=nil{
		t.Error(err)
	}
	data, terr := token.ControlToken(stringToken)
	if terr!=nil{
		t.Error(terr)
	}
	fmt.Println(data)

	tokendata, dataerr:= token.GetTokenData(stringToken)
	if dataerr !=nil{
		t.Error("There are some error")
	}
	if tokendata.Username !="Kerimberdi"{
		t.Error("Its error")
	}
	fmt.Println(tokendata.Id)
}

