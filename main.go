package main

import (
	"log"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/AbdullahMarikkar/goCrud/routers"
)

// TODO : Add JWT Access Token authentication and authorization, add access token in http only cookie
// TODO : Middleware implementation for authorization
// TODO : Refresh token mechanism with persistent table and IP Tracking

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	routers.Init()

}

func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}
