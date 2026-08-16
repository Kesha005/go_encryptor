package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kesha005/go_encryptor/token"
)

func TestTokenMake(t *testing.T) {

	tserv := jwt.New(jwt.JWTConfig{
		Secret:               "12222",
		AccessTokenDuration:  10,
		RefreshTokenDuration: 150,
		Issuer:               "test",
		Audience:             "test-app",
	})

	tokenstring, err := tserv.CreateAccessToken(1, "phone", "admin")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(tokenstring)
}

func TestTokenValidate(t *testing.T) {
	tserv := jwt.New(jwt.JWTConfig{
		Secret:               "12222",
		AccessTokenDuration:  10,
		RefreshTokenDuration: 150,
		Issuer:               "test",
		Audience:             "test-app",
	})

	tokenstring, err := tserv.CreateAccessToken(1, "phone", "admin")

	if err != nil {
		t.Error(err)
	}

	ok, err := tserv.IsAccessToken(tokenstring)

	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("Token verify failed")
	}

	payload, err := tserv.Verify(tokenstring)

	if err != nil {
		t.Error(err)
	}

	if payload.UserID != 1 || payload.Role != "admin" {
		t.Error("Failed to get correct token payload data")
	}
}

func TestTokenExpireTime(t *testing.T) {
	tserv := jwt.New(jwt.JWTConfig{
		Secret:               "12222",
		AccessTokenDuration:  3,
		RefreshTokenDuration: 150,
		Issuer:               "test",
		Audience:             "test-app",
	})

	tokenstring, err := tserv.CreateAccessToken(1, "phone", "admin")

	if err != nil {
		t.Error(err)
	}

	ok, err := tserv.IsAccessToken(tokenstring)

	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("Token verify failed")
	}

	time.Sleep(time.Second * 4)

	_,eerr:= tserv.Verify(tokenstring)

	if eerr==nil{
		t.Error("Token must be expired")
	}

}
