package apartmentHandler

import (
	"context"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	count1, err := apartmentCollection.CountDocuments(ctx, bson.M{"name": apartment.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count2, err := apartmentCollection.CountDocuments(ctx, bson.M{"address": apartment.Address})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if count1 > 0 || count2 > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Apartment already exist in the database"})
		return
	}
	apartment.ID = primitive.NewObjectID()
	apartment.ApartmentId = apartment.ID.Hex()
	apartment.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	apartment.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = apartmentCollection.InsertOne(ctx, apartment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, apartment)
}
