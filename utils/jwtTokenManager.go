package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


var secretKey = []byte("secret-key")

func CreateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "email": email, 
        "exp": time.Now().Add(time.Minute * 30).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) (bool,error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return false,err
	}
   
	if !token.Valid {
	   fmt.Errorf("invalid token")
	   return false,err
	}
   
	return true,nil
 }