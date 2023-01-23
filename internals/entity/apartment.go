package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Apartment struct {
	ID          primitive.ObjectID
	Name        string
	Address     AddressInfo
	Amenities   []string
	Images      []string
	Price       uint16
	Review      []primitive.ObjectID
	Bookings    []primitive.ObjectID
	RenterId    primitive.ObjectID
	ApartmentId primitive.ObjectID
}

type AddressInfo struct {
	Street string
	City   string
	State  string
	Zip    string
}
