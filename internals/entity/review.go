package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reviews struct {
	ID          primitive.ObjectID
	ApartmentId primitive.ObjectID
	UserId      primitive.ObjectID
	Rating      uint8
	Review      string
}
