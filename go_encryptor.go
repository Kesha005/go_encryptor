package go_encryptor

import (
	
	"golang.org/x/crypto/bcrypt"
)



func HashPassword(password string)(string,error){

	hashed,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err!=nil{
		
		
		return "",err
	}
	return string(hashed),nil
}


func CheckHash(password,hash string)(bool){

	res := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	if res!=nil{
		return false
	}
	return true

}