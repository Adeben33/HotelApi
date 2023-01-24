package main

import (
	"github.com/adeben33/HotelApi/cmd"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
)

func init() {
	mongoDBConnection.MongoDBConnection()
}

func main() {
	cmd.Setup()
}
