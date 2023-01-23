package userHandler

import (
	"context"
	"fmt"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/adeben33/HotelApi/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var validate = validator.New()

var userCollection *mongo.Collection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "user")

func SignUp(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	cancel()
	//bind the incoming data
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in binding data"})
		return
	}

	//validate the data
	validateErr := validate.Struct(user)
	if validateErr != nil {
		msg := fmt.Sprintf("Error in validating data")
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	//	check if the email has been in the database
	userCount1, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User counting failed"})
		return
	}
	//	check if the phone number has been in the data
	userCount2, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User counting failed"})
		return
	}

	if userCount1 > 0 || userCount2 > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exist in the database"})
	}

	user.ID = primitive.NewObjectID()
	user.UpdatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	harsh, err := utils.HarshPassword(user.Password)
	if err != nil {
		log.Panic(err)
		return
	}
	user.Password = harsh
	user.UserId = user.ID.Hex()

	result, InsertErr := userCollection.InsertOne(ctx, user)
	if InsertErr != nil {
		msg := fmt.Sprintf("user Item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, result)
}
