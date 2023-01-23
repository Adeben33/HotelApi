package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bookings struct {
	ID          primitive.ObjectID
	UserId      primitive.ObjectID
	ApartmentId primitive.ObjectID
	StartDate   time.Time
	EndDate     time.Time
	TotalPrice  uint32
	PaymentId   primitive.ObjectID
}
