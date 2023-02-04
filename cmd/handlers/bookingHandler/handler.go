package bookingHandler

import (
	"context"
	"fmt"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/adeben33/HotelApi/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var validate = validator.New()
var apartmentCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "apartment")
var bookingCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "booking")

func GetBookings(c *gin.Context) {
	var bookings []entity.Bookings
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//	get the booking
	cursor, findErr := bookingCollection.Find(ctx, bson.M{})
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding bookings"})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var booking entity.Bookings
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, bookings)
}

func GetBookingById(c *gin.Context) {
	var booking entity.Bookings
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//booking id
	bookingId := c.Param("bookingId")

	findErr := bookingCollection.FindOne(ctx, bson.M{"booking_id": bookingId}).Decode(&booking)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, findErr.Error())
		return
	}
	c.JSON(http.StatusOK, booking)
}

func CreateBooking(c *gin.Context) {
	//Only one apartment can be booked at a time
	var booking entity.Bookings
	//var apartment entity.Apartment
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := c.ShouldBindJSON(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving data"})
		return
	}
	if booking.ApartmentId == " " {
		msg := fmt.Sprintf("Warning! There is no apartment to create the necessaery booking")
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	//	check if the apartment is already booked and also how many days is it booked for
	//using the apartmentID in the booking struct to find the apartment details

	if !utils.CheckAvailability(booking.ApartmentId, booking.StartDate, booking.EndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dates are not available"})
		return
	}
	//
	booking.ID = primitive.NewObjectID()
	booking.BookingsId = booking.ID.Hex()

	result, insertEr := bookingCollection.InsertOne(ctx, booking)
	if insertEr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertEr.Error()})
		return
	}
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"apartment_id", booking.ApartmentId}}
	update := bson.D{
		{"$push", bson.D{
			{"bookings_id", booking.BookingsId},
		}},
	}
	_, updateErr := apartmentCollection.UpdateOne(ctx, filter, update, opts)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
