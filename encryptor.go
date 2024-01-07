package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)


func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
	
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text string) (string, error) {
	
var bytes = []byte(os.Getenv("IV_16_KEY"))

// This should be in an env file in production
var MySecret string = os.Getenv("SECRET_KEY")

	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string) (string, error) {
	
var bytes = []byte(os.Getenv("IV_16_KEY"))

// This should be in an env file in production
var MySecret string = os.Getenv("SECRET_KEY")
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
func main() {
	StringToEncrypt := "Encrypting this string"
	godotenv.Load(".env")
	fmt.Println(StringToEncrypt)
	encText, err := Encrypt(StringToEncrypt)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	// To decrypt the original StringToEncrypt
	decText, err := Decrypt(encText)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)
}
