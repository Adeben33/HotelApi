package utils

import (
	"context"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var apartmentCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "apartment")
var bookingCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "booking")

func CheckAvailability(apartmentId string, bookingStartDate, bookingEndDate time.Time) bool {
	var apartment entity.Apartment
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	findingErr := apartmentCollection.FindOne(ctx, bson.D{{"apartment_id", apartmentId}}).Decode(&apartment)
	if findingErr != nil {
		panic(findingErr)
	}

	//check if a booking already exist, if not not data is assigned yet

	if len(apartment.BookingsId) == 0 {
		return true
	}

	// Iterate through existing bookings and check for overlaps

	for _, bookingsId := range apartment.BookingsId {
		var booking entity.Bookings
		_ = bookingCollection.FindOne(ctx, bson.D{{"bookings_id", bookingsId}}).Decode(&booking)

		//the apartment has a list of bookingId to check the booking for the specific apartment
		if (bookingStartDate.After(booking.StartDate) && bookingStartDate.Before(booking.EndDate)) ||
			(bookingEndDate.After(booking.StartDate) && bookingEndDate.Before(booking.EndDate)) ||
			(bookingStartDate.Before(booking.StartDate) && bookingEndDate.After(booking.EndDate)) {
			return false
		}
	}

	return true
}

