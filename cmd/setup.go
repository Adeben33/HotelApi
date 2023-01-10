package cmd

import "github.com/adeben33/HotelApi/internals/dataBaseStore/postgresDB"

func Setup() {
	postgresDB.ConnectToDb()
}
