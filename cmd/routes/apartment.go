package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers/apartmentHandler"
	"github.com/gin-gonic/gin"
)

func ApartmentRoutes(route *gin.Engine) {
	route.GET("/apartments", apartmentHandler.GetAllApartment)
	route.GET("/apartments/:apartmentId", apartmentHandler.GetApartment)
	route.POST("/apartments", apartmentHandler.CreateApartment)

	//Fake apartment creator
	route.GET("/apartments/fake", apartmentHandler.FakeApartment)

	//route.PUT("/apartments/:id", apartmentHandler.UpdateApartment)
	//route.DELETE("/apartment/:id", apartmentHandler.DeleteApartment)

	//Retreive a list of apartments with specific amenities (?amenities=amenities1,amenities2,amenities3...)
	route.GET("/apartments", apartmentHandler.GetAllApartmentWithAmenities)

	//retrive all reviews for a specific apartment
	route.GET("/apartments/:apartmentId/reviews", apartmentHandler.GetAllReviews)

	//	Retrieve a list of apartments rented by a specific user
	route.GET("/apartments", apartmentHandler.GetRentedApartment)
	////	Retrieve a list of apartments within a specific price range.
	//route.GET("/apartments", apartmentHandler.GetApartmentWithPrice())
	////Retrieve a list of apartments with a specific number of bedrooms.
	//route.GET("/apartments", apartmentHandler.GetApartmentWithNumberofBedrooms())
	////Retrieve a list of apartments that are available on a specific date range.
	//route.GET("/apartments", apartmentHandler.GetAvailableApartments())
	////Retrieve all bookings for a specific apartment.
	//route.GET("/apartments/:id/bookings", apartmentHandler.GetAllApartmentBookings)
	////Create a new booking for a specific apartment.
	//route.POST("/apartments/:id/bookings", apartmentHandler.CreateBooking)
	////Retrieve all images for a specific apartment.
	//route.GET("/apartments/:id/images", apartmentHandler.ApartmentImages())
	////Upload new images for a specific apartment./upload picture for the apartment
	//route.POST("/apartments/:id/images", apartmentHandler.UploadApartmentImage())
	////	Retrieve a list of apartments with a specific rating range.
	//route.GET("/apartments", apartmentHandler.GetApartmentWithRatings())
	//// Retrieve the average rating for a specific apartment.
	//route.GET("/apartments/:id/averageRating", apartmentHandler.ApartmentAverageRating())
	////Retrieve the last booking for a specific apartment.
	//route.GET("/apartments/:id/lastBooking", apartmentHandler.ApartmentLastBooking())
	////Retrieve upcoming bookings for a specific apartment.
	//route.GET("/apartments/:id/upcomingBo0king", apartmentHandler.ApartmentUpcomingBooking())
	////Retrieve a list of apartments based on a specific location (e.g. by latitude and longitude).
	//route.GET("/apartments", handlers.GetApartmentInALocation())
	////Retrieve all amenities for a specific apartment.
	//route.GET("/apartments/:id/amenities", apartmentHandler.GetApartmentAmenities())
	////Update amenities for a specific apartment.
	//route.PUT("/apartments/:id/amenities", apartmentHandler.UpdateApartmentAmenities())
	////Retrieve the availability status of a specific apartment eg available, booked etc
	//route.GET("/apartments/:id/propertyAvailability", apartmentHandler.GetApartmentAmenities())
	////Retrieve a list of apartments based on a specific property type (e.g. house, apartment, etc.).
	////Retrieve the total price for a specific apartment based on a specific date range.
	//route.GET("/apartments/:id/totalPrice", apartmentHandler.GetApartmenTotalPrice())
	////Retrieve the check-in and check-out time for a specific apartment.
	//route.GET("/apartments/:id/checkInCheckOut", apartmentHandler.GetApartmenCheckInCheckOut())
	////Update the check-in and check-out time for a specific
	//route.PUT("/apartments/:id/checkInCheckOut", apartmentHandler.GetApartmenCheckInCheckOut())
	////Retrieve the booking calendar for a specific apartment.
	//route.GET("/apartments/:id/bookingCalender", apartmentHandler.GetApartmenBookingCalender())
	//// Retrieve the location of a specific apartment.
	//route.GET("/apartments/:id/propertyLocation", apartmentHandler.GetApartmenLocation())
	////Update the location of a specific apartment.
	//route.PUT("/apartments/:id/propertyLocation", apartmentHandler.UpdateApartmenLocation())

}
