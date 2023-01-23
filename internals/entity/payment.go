package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID         primitive.ObjectID
	BookingsId primitive.ObjectID
	Method     string
	Amount     uint32
	Status     string
}
