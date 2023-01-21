package routes

import "github.com/gin-gonic/gin"

func ReviewsRoutes(route *gin.Engine) {
	route.GET("/reviews", handlers.GetAllReviews)
	route.GET("/reviews/:id", handlers.GetAReview)

}
