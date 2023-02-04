package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Apartment struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name" bson:"name" validate:"required"`
	Address       AddressInfo        `json:"address" bson:"address" validate:"omitempty"`
	NumberofRooms uint8              `bson:"numberof_rooms" json:"numberofRooms" validate:"omitempty"`
	Amenities     []string           `json:"amenities" bson:"amenities" validate:"omitempty"`
	Images        []string           `json:"images" bson:"images" validate:"omitempty"`
	Price         uint16             `json:"price" bson:"price" validate:"required,number"`
	Review        []string           `json:"review" bson:"review" validate:"omitempty"`
	CreatedAt     time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updated_at"`
	BookingsId    []string           `json:"bookingsId" bson:"bookings_id" validate:"omitempty"`
	RenterId      string             `json:"renterId" bson:"renter_id"`
	ApartmentId   string             `json:"apartmentId" bson:"apartment_id"`
}

type AddressInfo struct {
	Street string `json:"street" bson:"street" validate:"required"`
	City   string `json:"city" bson:"city" validate:"required"`
	State  string `json:"state" bson:"state" validate:"required"`
	Zip    string `json:"zip" bson:"zip" validate:"omitempty"`
}

type ApartmentRes struct {
	Name          string      `json:"name" bson:"name" validate:"required"`
	Address       AddressInfo `json:"address" bson:"address" validate:"omitempty"`
	NumberofRooms uint8       `bson:"numberof_rooms" json:"numberofRooms" validate:"omitempty"`
	Amenities     []string    `json:"amenities" bson:"amenities" validate:"omitempty"`
	Images        []string    `json:"images" bson:"images" validate:"omitempty"`
	Price         uint16      `json:"price" bson:"price" validate:"required,number"`
	Review        []string    `json:"review" bson:"review" validate:"omitempty"`
	CreatedAt     time.Time   `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time   `json:"updatedAt" bson:"updated_at"`
}
