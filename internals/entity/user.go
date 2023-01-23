package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID
	Email            string
	Password         string
	Role             string
	FirstName        string
	LastName         string
	Address          []UserAddressInfo
	Phone            string
	Bookings         []primitive.ObjectID
	RenterProperties []primitive.ObjectID
}

type UserAddressInfo struct {
	Street string
	City   string
	State  string
	Zip    string
}
