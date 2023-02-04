package bookingHandler

import (
	"context"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetBookings(c *gin.Context) {

}

func CreateBooking(c *gin.Context) {
	//Only one apartment can be booked at a time
	var booking entity.Bookings
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := c.BindJSON(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving data"})
		return
	}
	booking.ApartmentId
}
