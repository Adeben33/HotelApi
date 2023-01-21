package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	//signup
	routes.POST("/signup", handlers.Signup())
	//	login
	routes.POST("/login", handlers.Login())
	//	GET user
	routes.GET("/users/:userId", handlers.GetUser())
	//	GET all users
	routes.GET("/users", handlers.GetAllUsers())

}
