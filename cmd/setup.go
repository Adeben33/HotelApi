package cmd

import (
	"github.com/adeben33/HotelApi/cmd/routes"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/postgresDB"
	"github.com/adeben33/HotelApi/internals/entity/userEntity"
	"log"
)

func Setup() {
	//Database
	db, err := postgresDB.ConnectToDB()
	if err != nil {
		log.Panic("Unable to connect to Db")
	}
	db.GetConn().AutoMigrate(&userEntity.CreateUser{}, &userEntity.Apartment{})

	//route
	routes.RouteSetup()
}
