package routes

import "github.com/gin-gonic/gin"

func ReviewsRoutes(route *gin.Engine) {
	route.GET("/reviews", handlers.GetAllReviews)
	route.GET("/reviews/:id", handlers.GetAReview)
	route.POST("/reviews", handlers.CreateReview)
	route.PUT("/reviews/:id", handlers.UpdateReview())
	route.DELETE("/reviews/:id", handlers.DeleteReview())
	//	Retrieve a list of reviews for a specific apartment.
	route.GET("/reviews", handlers.GetAnApartmentReview)
	//Retrieve a list of reviews written by a specific user.
	route.GET("/reviews", handlers.GetAReview)
	//Retrieve a list of reviews with a specific rating.
	route.GET("/reviews", handlers.GetASpecificReviewWithRating())

}
