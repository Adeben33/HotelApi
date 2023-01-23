package routes

import (
	"github.com/adeben33/HotelApi/cmd/handlers/reviewHandler"
	"github.com/gin-gonic/gin"
)

func ReviewsRoutes(route *gin.Engine) {
	route.GET("/reviews", reviewHandler.GetAllReviews)
	//route.GET("/reviews/:id", reviewHandler.GetAReview)
	//route.POST("/reviews", reviewHandler.CreateReview)
	//route.PUT("/reviews/:id", reviewHandler.UpdateReview())
	//route.DELETE("/reviews/:id", reviewHandler.DeleteReview())
	////	Retrieve a list of reviews for a specific apartment.
	//route.GET("/reviews", reviewHandler.GetAnApartmentReview)
	////Retrieve a list of reviews written by a specific user.
	//route.GET("/reviews", reviewHandler.GetAReview)
	////Retrieve a list of reviews with a specific rating.
	//route.GET("/reviews", reviewHandler.GetASpecificReviewWithRating())

}
