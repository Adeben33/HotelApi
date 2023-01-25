package apartmentHandler

import (
	"context"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

var validate = validator.New()
var apartmentCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "apartment")

func GetAllApartment(c *gin.Context) {
}

func CreateApartment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var apartment entity.Apartment
	if err := c.BindJSON(&apartment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while bindind data"})
		return
	}

}
