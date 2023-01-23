package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers/userHandler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	//signup
	routes.POST("/users/signup", userHandler.SignUp)
	//	login
	//routes.POST("users/login", userHandler.Login())
	////Returns a new token with a renewed expiration time.
	//routes.POST("users/refresh", userHandler.RefreshLogin())

	////	GET user
	//routes.GET("/users/:userId", userHandler.GetUser())
	////	GET all users
	//routes.GET("/users", userHandler.GetAllUsers())
	////Update an existing user by ID.
	//routes.PUT("/users/:userId", userHandler.UpdateUser())
	////Delete an existing user by ID. (Admin Only)
	//routes.DELETE("/users/:userId", userHandler.DeleteUser())
	//// Retrieve all bookings for a specific user.
	//routes.GET("/users/:id/bookings", userHandler.GetUserBookings())
	////GET Retrieve all properties rented by a specific user.
	//routes.GET("/users/:id/renterProperties", userHandler.GetRenterProperties())
	////	Retrieve all properties managed by a specific user.
	//routes.GET("/users/:id/managedProperties", userHandler.GetManagedProperties())
	////Retrieve the role of a specific user.
	//routes.GET("/users/:id/role", userHandler.GetUserRole())
	////Update the role of a specific user. (Admin Only)
	//routes.PUT("/users/:id/role", userHandler.UpdateUserRole())
	////Retrieve the recent bookings for a specific user.
	//routes.GET("/users/:id/recentBookings", userHandler.GetRecentBookings())
	////Retrieve the upcoming bookings for a specific user.
	//routes.GET("/users/:id/UpComingBookings", userHandler.GetUpcomingBookings())
	////	Retrieve the booking history for a specific user.
	//routes.GET("/users/:id/BookingHistory", userHandler.GetBookingHistory())

}
