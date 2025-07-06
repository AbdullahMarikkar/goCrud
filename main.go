package main

import (
	"log"

	"github.com/AbdullahMarikkar/goCrud/models"
	"github.com/AbdullahMarikkar/goCrud/routers"
)

// TODO : Refresh token mechanism with persistent table and IP Tracking
// TODO : In Refresh Token Mechanism, orchestrate validation of refresh token by IP, expiry and blacklist(and creation of new RT and AT )
// TODO : delete old RT and add it to Blacklist after Creation of new RT
// TODO : If RT expired send unauthorized error and request for login
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
