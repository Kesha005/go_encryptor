package main

import (
	"fmt"
	"testing"

	"github.com/Kesha005/go_encryptor/token"
	"github.com/joho/godotenv"
)



func TestToken(t *testing.T){
	godotenv.Load("../.env")
	var username = "Kerim"
	var id = 2
	stringToken, err := token.GenerateToken(id ,username)
	if err !=nil{
		t.Error(err)
	}
	data, terr := token.ControlToken(stringToken)
	if terr!=nil{
		t.Error(terr)
	}
	fmt.Println(data)
}

