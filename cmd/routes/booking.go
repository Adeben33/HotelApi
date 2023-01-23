package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers/bookingHandler"
	"github.com/gin-gonic/gin"
)

func BookingRoutes(route *gin.Engine) {
	route.GET("/bookings", bookingHandler.GetBookings())
	//route.GET("/bookings/:id", bookingHandler.GetBookingbyId())
	//route.POST("/bookings", bookingHandler.CreateBookings())
	//route.PUT("/bookings/:id", bookingHandler.UpdateBooking())
	//route.DELETE("/bookings/:id",bookingHandler.DeleteBooking)
	////Retrieve the apartment associated with a specific booking.
	//route.GET("/bookings/:id/apartment", bookingHandler.Getapartment())
	////Retrieve the user associated with a specific booking.
	//route.GET("/bookings/:id/apartment", bookingHandler.GetapartmentUser())
	////Retrieve the payment associated with a specific booking.
	//route.GET("/bookings/:id/payment", bookingHandler.GetapartmentPayment())
	////	Retrieve the review associated with a specific booking.
	//route.GET("/bookings/:id/review", bookingHandler.GetapartmentReview())
	////Cancel a specific booking.
	//route.GET("/bookings/:id/review", bookingHandler.CancelABooking())
	//// Confirm a specific booking.
	//route.GET("/bookings/:id/confirm", bookingHandler.ConfirmABooking())
	//
	////Retrieve all bookings for a specific user.
	//route.GET("/bookings/user/:userId", bookingHandler.GetBookingsForaUser())
	////Retrieve all bookings for a specific apartment.
	//route.GET("/bookings/apartment/:apartmentId", bookingHandler.GetBookingsForanApartment())
	////Retrieve all bookings for a specific date.
	//route.GET("/bookings/date/:date", bookingHandler.GetBookingsForaDay())
	////Retrieve all upcoming bookings.
	//route.GET("/bookings/upcoming", bookingHandler.GetUpcomingBookings())
	//
	////Retrieve all past bookings.
	//route.GET("/bookings/past", bookingHandler.GetPastBookings())
	//
	////Retrieve all cancelled bookings.
	//route.GET("/bookings/cancelled", bookingHandler.GetCancelledBookings())
	//
	////Retrieve all confirmed bookings.
	//route.GET("/bookings/confirmed", bookingHandler.GetConfirmedBookings())
	//
	////Retrieve all checkedin bookings.
	//route.GET("/bookings/checkedin", bookingHandler.GetCheckedinBookings())
	//
	////Retrieve all checkedout bookings.
	//route.GET("/bookings/checkedout", bookingHandler.GetCheckedoutBookings())
	//
	////Retrieve the invoice for a specific booking.
	//route.GET("/bookings/:id/invoice", bookingHandler.GetBookingsInvoice())
	//
	////Send a reminder for a specific booking.
	//route.GET("/bookings/:id/reminder", bookingHandler.SendBookingReminder())

	//

}
