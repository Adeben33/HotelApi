package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Apartment struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	Name        string               `json:"name" bson:"name" validate:"required"`
	Address     AddressInfo          `json:"address" bson:"address" validate:"omitempty"`
	Amenities   []string             `json:"amenities" bson:"amenities" validate:"omitempty"`
	Images      []string             `json:"images" bson:"images" validate:"omitempty"`
	Price       uint16               `json:"price" bson:"price" validate:"required,number"`
	Review      []primitive.ObjectID `json:"review" bson:"review" validate:"omitempty"`
	Bookings    []primitive.ObjectID `json:"bookings" bson:"bookings" validate:"omitempty"`
	RenterId    primitive.ObjectID   `json:"renterId" bson:"renter_id"`
	ApartmentId primitive.ObjectID   `json:"apartmentId" bson:"apartment_id"`
}

type AddressInfo struct {
	Street string `json:"street" bson:"street" validate:"required"`
	City   string `json:"city" bson:"city" validate:"required"`
	State  string `json:"state" bson:"state" validate:"required"`
	Zip    string `json:"zip" bson:"zip" validate:"omitempty"`
}
