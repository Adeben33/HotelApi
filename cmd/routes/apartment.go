package routes

import "github.com/gin-gonic/gin"

func ApartmentRoutes(route *gin.Engine) {
	route.GET("/apartments", handlers.GetAllApartment())
	route.GET("/apartments/:id", handlers.GetAparment())
	route.POST("/apartments", handlers.CreateApartment)
	route.PUT("/apartments/:id", handlers.UpdateApartment)
	route.DELETE("/apartment/:id", handlers.DeleteApartment)
	//Retreive a list of apartments with specific amenities
	route.GET("/apartments", handlers.GetAllApartmentWithAmenities())
	//retrive all reviews for a specific apartment
	route.POST("/apartmemts/:id/reviews", handlers.GetAllReviews())
	//	Retrieve a list of apartments rented by a specific user
	route.GET("/apartments", handlers.GetRentedApartment())
	//	Retrieve a list of apartments within a specific price range.
	route.GET("/apartments", handlers.GetApartmentWithPrice())
	//Retrieve a list of apartments with a specific number of bedrooms.
	route.GET("/apartments", handlers.GetApartmentWithNumberofBedrooms())
	//Retrieve a list of apartments that are available on a specific date range.
	route.GET("/apartments", handlers.GetAvailableApartments())
	//Retrieve all bookings for a specific apartment.
	route.GET("/apartments/:id/bookings")
	//Create a new booking for a specific apartment.
	route.POST("/apartments/:id/bookings", handlers.CreateBooking)
	//Retrieve all images for a specific apartment.
	route.GET("/apartments/:id/images", handlers.ApartmentImages())
	//Upload new images for a specific apartment./upload picture for the apartment
	route.POST("/apartments/:id/images", handlers.UploadApartmentImage())
	//	Retrieve a list of apartments with a specific rating range.
	route.GET("/apartments", handlers.GetApartmentWithRatings())
	// Retrieve the average rating for a specific apartment.
	route.GET("/apartments/:id/averageRating", handlers.ApartmentAverageRating())
	//Retrieve the last booking for a specific apartment.
	route.GET("/apartments/:id/lastBooking", handlers.ApartmentLastBooking())
	//Retrieve upcoming bookings for a specific apartment.
	route.GET("/apartments/:id/upcomingBo0king", handlers.ApartmentUpcomingBooking())
	//Retrieve a list of apartments based on a specific location (e.g. by latitude and longitude).
	route.GET("/apartments", handlers.GetApartmentInALocation())
	//Retrieve all amenities for a specific apartment.
	route.GET("/apartments/:id/amenities", handlers.GetApartmentAmenities())
	//Update amenities for a specific apartment.
	route.PUT("/apartments/:id/amenities", handlers.UpdateApartmentAmenities())
	//Retrieve the availability status of a specific apartment eg available, booked etc
	route.GET("/apartments/:id/propertyAvailability", handlers.GetApartmentAmenities())
	//Retrieve a list of apartments based on a specific property type (e.g. house, apartment, etc.).
	//Retrieve the total price for a specific apartment based on a specific date range.
	route.GET("/apartments/:id/totalPrice", handlers.GetApartmenTotalPrice())
	//Retrieve the check-in and check-out time for a specific apartment.
	route.GET("/apartments/:id/checkInCheckOut", handlers.GetApartmenCheckInCheckOut())
	//Update the check-in and check-out time for a specific
	route.PUT("/apartments/:id/checkInCheckOut", handlers.GetApartmenCheckInCheckOut())
	//Retrieve the booking calendar for a specific apartment.
	route.GET("/apartments/:id/bookingCalender", handlers.GetApartmenBookingCalender())
	// Retrieve the location of a specific apartment.
	route.GET("/apartments/:id/propertyLocation", handlers.GetApartmenLocation())
	//Update the location of a specific apartment.
	route.PUT("/apartments/:id/propertyLocation", handlers.UpdateApartmenLocation())

}
