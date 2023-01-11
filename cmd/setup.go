package cmd

import (
	"github.com/adeben33/HotelApi/cmd/routes"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/postgresDB"
)

func Setup() {
	postgresDB.ConnectToDb()
	routes.RouteSetup()
}
