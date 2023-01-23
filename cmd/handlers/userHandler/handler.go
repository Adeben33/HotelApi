package userHandler

import (
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
)

type userRepo struct {
}

func SignUp(c *gin.Context) {
	var userDetails entity.User
	//bind the data from the request
	err := c.ShouldBindJSON(&userDetails)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": " Error while binding"})
		return
	}
	////send to user service
	//err = user.SignUpUser(userDetails)
	//if err != nil {
	//	log.Panic("Unable to sign up user")
	//}
}
