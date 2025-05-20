package main

import (
	"log"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/AbdullahMarikkar/goCrud/routers"
)

// TODO : Create User Service Endpoints

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
