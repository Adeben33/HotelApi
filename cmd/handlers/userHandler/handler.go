package userHandler

import (
	"context"
	"fmt"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/adeben33/HotelApi/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jaswdr/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var validate = validator.New()
var secretKey = "Adeniyi"

var userCollection *mongo.Collection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "user")

func SignUp(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
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

	//check if the email has been in the database
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
		return
	}

	user.ID = primitive.NewObjectID()
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
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
		//msg := fmt.Sprintf("user Item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": InsertErr.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//	Get the bind data
	var user entity.UserLogin
	var userDbDetails entity.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//	validate the data

	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login input not correctly inputed"})
		return
	}
	//	verify if email is in the database
	findErr := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&userDbDetails)
	if findErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findErr.Error()})
		return
	}

	//	verify Password
	err := utils.VerifyPassword(userDbDetails.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	//Check if the user already has a token then the user is alread logged in
	if userDbDetails.Token != " " {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user already logged in"})
		return
	}
	if userDbDetails.RefreshToken != " " {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user already logged in"})
		return
	}
	//	generate Token
	token, refreshToken, err := utils.CreateToken(userDbDetails, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create token"})
		return
	}

	//	Add token to the database
	userDbDetails.Token = token
	userDbDetails.RefreshToken = refreshToken
	//	return User has logged in
	c.JSON(http.StatusOK, userDbDetails)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user entity.User

	filter := bson.M{"userId": userId}
	findErr := userCollection.FindOne(ctx, filter).Decode(&user)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func FakeUsers(c *gin.Context) {
	//	This is to create a set of users with password
	//ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//defer cancel()
	//var users []entity.User
	role := c.Param("role")
	password := c.Param("password")
	fake := faker.New()
	for i := 0; i <= 50; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user entity.User
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		user.FirstName = fake.Person().FirstName()
		user.LastName = fake.Person().LastName()
		user.Role = role
		user.Email = gofakeit.Email()
		user.Phone = fake.Phone().E164Number()
		user.Password, _ = utils.HarshPassword(password)
		userCollection.InsertOne(ctx, user)
	}
	c.JSON(http.StatusOK, gin.H{"success": "success"})
}

func GetAllUsers(c *gin.Context) {
	var users []entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{}
	findOptions := options.Find()

	//This is used to search in the database with respect to the given search details
	if search := c.Query("search"); search != " " {
		filter = bson.M{
			"$or": []bson.M{{
				"firstName": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
			},
				{
					"lastName": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
				{
					"Address": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
				{
					"email": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
				{
					"userId": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
			},
		}
	}

	//Sorting
	if sort := c.Query("sort"); sort != " " {
		if sort == "asc" {
			findOptions.SetSort(bson.M{
				"userId": 1,
			})
		} else if sort = c.Query("sort"); sort == "desc" {
			findOptions.SetSort(bson.M{
				"userId": -1,
			})
		}
	}

	cursor, err := userCollection.Find(ctx, filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting data"})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user entity.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}
