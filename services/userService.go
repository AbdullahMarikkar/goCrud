package services

import (
	"fmt"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/AbdullahMarikkar/goCrud/utils"
)

func AuthorizeUser(email string,password string)(*models.User,error){

	user, err := models.GetUserByEmail(email)

	if(!utils.VerifyPassword(password,user.Password)){
		fmt.Println("Error in AuthorizeUser, Password is Incorrect",err)
		return nil,err
	}

	if err != nil{
		return nil,err
	}

	return user,nil
}