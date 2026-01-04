package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string)(string,error)  {
	//this func make the password into hashed form
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func CheckPassword(hash string,password string )bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err==nil
}