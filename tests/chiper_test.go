package main

import (
	"fmt"
	"testing"

	"github.com/Kesha005/go_encryptor"
	"github.com/joho/godotenv"
)

func TestChiper(t *testing.T) {
	godotenv.Load("../.env")

	var word string = "Hello world"
	chiped, err := go_encryptor.Encrypt(word)
	if err != nil {
		fmt.Println(err)
	}
	dechiped, deeer := go_encryptor.Decrypt(chiped)
	if deeer != nil {
		fmt.Println(deeer.Error())
	}
	if dechiped != word {
		t.Error("It is not word which we chiped")
	}

}
