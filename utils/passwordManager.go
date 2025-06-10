package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
	if err!=nil{
		fmt.Println("Error",err)
	} 

	return string(hash)
}

func VerifyPassword(password string,hashedPassword string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))

	if err != nil {
		fmt.Println("Error",err)
		return false
	}

	return true
}