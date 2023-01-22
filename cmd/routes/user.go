package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	//signup
	routes.POST("/users/signup", handlers.Signup())
	//	login
	routes.POST("users/login", handlers.Login())
	//Returns a new token with a renewed expiration time.
	routes.POST("users/refresh", handlers.RefreshLogin())

	//	GET user
	routes.GET("/users/:userId", handlers.GetUser())
	//	GET all users
	routes.GET("/users", handlers.GetAllUsers())
	//Update an existing user by ID.
	routes.PUT("/users/:userId", handlers.UpdateUser())
	//Delete an existing user by ID. (Admin Only)
	routes.DELETE("/users/:userId", handlers.DeleteUser())
	// Retrieve all bookings for a specific user.
	routes.GET("/users/:id/bookings", handlers.GetUserBookings())
	//GET Retrieve all properties rented by a specific user.
	routes.GET("/users/:id/renterProperties", handlers.GetRenterProperties())
	//	Retrieve all properties managed by a specific user.
	routes.GET("/users/:id/managedProperties", handlers.GetManagedProperties())
	//Retrieve the role of a specific user.
	routes.GET("/users/:id/role", handlers.GetUserRole())
	//Update the role of a specific user. (Admin Only)
	routes.PUT("/users/:id/role", handlers.UpdateUserRole())
	//Retrieve the recent bookings for a specific user.
	routes.GET("/users/:id/recentBookings", handlers.GetRecentBookings())
	//Retrieve the upcoming bookings for a specific user.
	routes.GET("/users/:id/UpComingBookings", handlers.GetUpcomingBookings())
	//	Retrieve the booking history for a specific user.
	routes.GET("/users/:id/BookingHistory", handlers.GetBookingHistory())

}
