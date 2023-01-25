package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bookings struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"userId"`
	ApartmentId primitive.ObjectID `bson:"apartment_id" json:"apartmentId"`
	StartDate   time.Time          `bson:"start_date" json:"startDate"`
	EndDate     time.Time          `bson:"end_date" json:"endDate"`
	TotalPrice  uint32             `bson:"total_price" json:"totalPrice"`
	PaymentId   primitive.ObjectID `bson:"payment_id" json:"paymentId"`
}
