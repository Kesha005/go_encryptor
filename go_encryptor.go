package go_encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)

}
func decode(s string) []byte {
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

	return encode(cipherText), nil
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
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
