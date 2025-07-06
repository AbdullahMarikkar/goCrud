package models

import (
	"fmt"

	_ "modernc.org/sqlite"
)

type RefreshToken struct{
	Id int `json:"id"`
	Token string `json:"token"`
	Email string `json:"email"`
	Ip string `json:"ip"`
	Created_At string `json:"createdAt"`
}

type CreateRefreshTokenDto struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Ip string `json:"ip"`
}

func CreateRefreshToken(refreshToken CreateRefreshTokenDto)(*RefreshToken,error){
	result , err := DB.Exec(`INSERT INTO refresh_tokens (token, email, ip) VALUES ( ?, ?, ? )`, refreshToken.Token, refreshToken.Email,refreshToken.Ip)

	var newToken RefreshToken

	if err != nil {
		fmt.Println("Error in CreateRefreshToken",err)
		return nil, err
	}

	id, err := result.LastInsertId()

	newToken.Id = int(id)
	newToken.Token = refreshToken.Token
	newToken.Email = refreshToken.Email
	newToken.Ip = refreshToken.Ip

	return &newToken,nil
}

// TODO : Delete Refresh Token By ID
// TODO : Get Refresh Tokens By email