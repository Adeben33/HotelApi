package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID               primitive.ObjectID   `bson:"_id" json:"id"`
	Email            string               `bson:"email" json:"email" validate:"required,email"`
	Password         string               `json:"password" bson:"password" validate:"required,min=6"",`
	Role             string               `json:"role" bson:"role" validate:"required,oneof='admin''user''renter'"`
	FirstName        string               `json:"firstName" bson:"first_name" validate:"required"`
	LastName         string               `json:"lastName" bson:"last_name" validate:"required"`
	Address          []UserAddressInfo    `json:"address" bson:"address" validate:"omitempty"`
	Phone            string               `json:"phone" bson:"phone" validate:"omitempty"`
	Bookings         []primitive.ObjectID `json:"bookings" bson:"bookings" validate:"omitempty"`
	RenterProperties []primitive.ObjectID `json:"renterProperties" bson:"renter_properties" validate:"omitempty"`
	Token            string               `bson:"token" json:"token" validate:"omitempty"`
	RefreshToken     string               `bson:"refresh_token" json:"RefreshToken" validate:"omitempty"`
	CreatedAt        time.Time            `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time            `json:"updatedAt" bson:"updated_at"`
	UserId           string               `json:"userId" bson:"userId"`
}

type UpdateUserStruct struct {
	ID               primitive.ObjectID   `bson:"_id" json:"id"`
	Email            string               `bson:"email" json:"email" validate:"email"`
	Password         string               `json:"password" bson:"password" validate:"omitempty",`
	Role             string               `json:"role" bson:"role" validate:"omitempty"`
	FirstName        string               `json:"firstName" bson:"first_name" validate:"omitempty"`
	LastName         string               `json:"lastName" bson:"last_name" validate:"omitempty"`
	Address          []UserAddressInfo    `json:"address" bson:"address" validate:"omitempty"`
	Phone            string               `json:"phone" bson:"phone" validate:"omitempty"`
	Bookings         []primitive.ObjectID `json:"bookings" bson:"bookings" validate:"omitempty"`
	RenterProperties []primitive.ObjectID `json:"renterProperties" bson:"renter_properties" validate:"omitempty"`
	Token            string               `bson:"token" json:"token" validate:"omitempty"`
	RefreshToken     string               `bson:"refresh_token" json:"RefreshToken" validate:"omitempty"`
	CreatedAt        time.Time            `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time            `json:"updatedAt" bson:"updated_at"`
	UserId           string               `json:"userId" bson:"userId"`
}

type UserAddressInfo struct {
	Street string `json:"street" bson:"street"`
	City   string `json:"city" bson:"city"`
	State  string `json:"state" bson:"state"`
	Zip    string `json:"zip" bson:"zip"`
}

type UserLogin struct {
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=6"",`
}
