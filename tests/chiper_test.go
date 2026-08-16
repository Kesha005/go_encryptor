package tests

import (
	"fmt"
	"testing"

	"github.com/Kesha005/go_encryptor"
)

func TestHasher(t *testing.T) {
	hash, err := go_encryptor.HashPassword("123456")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestHashChecke(t *testing.T) {
	hash, err := go_encryptor.HashPassword("123456")

	if err != nil {
		t.Error(err)
	}

	ok:= go_encryptor.CheckHash("123456",hash)

	if !ok{
		t.Error("Something went wrong")
	}

}
