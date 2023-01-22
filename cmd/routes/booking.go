package routes

import "github.com/gin-gonic/gin"

func BookingRoutes(route *gin.Engine) {
	route.GET("/bookings", handlers.GetBookings())
	route.GET("/bookings/:id", handlers.GetBookingbyId())
	route.POST("/bookings", handlers.CreateBookings())
	route.PUT("/bookings/:id", handlers.UpdateBooking())
	route.DELETE("/bookings/:id", handlers.DeleteBooking)
	//Retrieve the apartment associated with a specific booking.
	route.GET("/bookings/:id/apartment", handlers.Getapartment())
	//Retrieve the user associated with a specific booking.
	route.GET("/bookings/:id/apartment", handlers.GetapartmentUser())
	//Retrieve the payment associated with a specific booking.
	route.GET("/bookings/:id/payment", handlers.GetapartmentPayment())
	//	Retrieve the review associated with a specific booking.
	route.GET("/bookings/:id/review", handlers.GetapartmentReview())
	//Cancel a specific booking.
	route.GET("/bookings/:id/review", handlers.CancelABooking())
	// Confirm a specific booking.
	route.GET("/bookings/:id/confirm", handlers.ConfirmABooking())

	//Retrieve all bookings for a specific user.
	route.GET("/bookings/user/:userId", handlers.GetBookingsForaUser())
	//Retrieve all bookings for a specific apartment.
	route.GET("/bookings/apartment/:apartmentId", handlers.GetBookingsForanApartment())
	//Retrieve all bookings for a specific date.
	route.GET("/bookings/date/:date", handlers.GetBookingsForaDay())
	//Retrieve all upcoming bookings.
	route.GET("/bookings/upcoming", handlers.GetUpcomingBookings())

	//Retrieve all past bookings.
	route.GET("/bookings/past", handlers.GetPastBookings())

	//Retrieve all cancelled bookings.
	route.GET("/bookings/cancelled", handlers.GetCancelledBookings())

	//Retrieve all confirmed bookings.
	route.GET("/bookings/confirmed", handlers.GetConfirmedBookings())

	//Retrieve all checkedin bookings.
	route.GET("/bookings/checkedin", handlers.GetCheckedinBookings())

	//Retrieve all checkedout bookings.
	route.GET("/bookings/checkedout", handlers.GetCheckedoutBookings())

	//Retrieve the invoice for a specific booking.
	route.GET("/bookings/:id/invoice", handlers.GetBookingsInvoice())

	//Send a reminder for a specific booking.
	route.GET("/bookings/:id/reminder", handlers.SendBookingReminder())

	//

}

