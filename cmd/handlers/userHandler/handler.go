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
	"strconv"
	"time"
)

var validate = validator.New()
var secretKey = "Adeniyi"

var userCollection *mongo.Collection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "user")
var bookingCollection *mongo.Collection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "bookings")

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
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		userCollection.InsertOne(ctx, user)
	}
	c.JSON(http.StatusOK, gin.H{"success": "success"})
}

func GetAllUsers(c *gin.Context) {
	var users []entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{}
	count, _ := userCollection.CountDocuments(ctx, filter)
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

	var perPage int64 = 9
	page, _ := strconv.Atoi(c.Query("page"))

	if page < 1 {
		page = 1
	}
	skipingLimit := (int64(page - 1)) * perPage
	findOptions = findOptions.SetLimit(perPage)
	findOptions = findOptions.SetSkip(skipingLimit)

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
	searchCount, _ := userCollection.CountDocuments(ctx, filter)

	c.JSON(http.StatusOK, gin.H{
		"data":           users,
		"All data count": count,
		"Search Count":   searchCount,
		"page":           page,
	})
}

func UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user entity.UpdateUserStruct
	userId := c.Param("userId")

	//var databaseUser entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error binding data"})
		return
	}
	validateErr := validate.Struct(&user)
	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in the inputted data"})
		return
	}
	//	get the user from the data base
	filter := bson.M{"userId": userId}
	//Check if the user is in the database
	count, _ := userCollection.CountDocuments(ctx, filter)
	if count < 0 {
		fmt.Println(count)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The user does not exist in the database"})
		return
	}

	//	Replacing each value
	var updatedUser primitive.D

	//	checks
	if user.FirstName != " " {
		updatedUser = append(updatedUser, bson.E{"first_name", user.FirstName})
	}

	if user.LastName != " " {
		updatedUser = append(updatedUser, bson.E{"last_name", user.LastName})
	}

	if user.Email != " " {
		updatedUser = append(updatedUser, bson.E{"email", user.Email})
	}

	if user.Phone != " " {
		updatedUser = append(updatedUser, bson.E{"phone", user.Phone})
	}

	if user.Address != nil {
		updatedUser = append(updatedUser, bson.E{"Address", user.Address})
	}

	if user.RenterProperties != nil {
		updatedUser = append(updatedUser, bson.E{"renter_properties", user.RenterProperties})
	}

	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updatedUser = append(updatedUser, bson.E{"updatedAt", user.UpdatedAt})

	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updatedUser}}, &opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{"userId": userId}
	deleteResult, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deleteResult)
}

func GetUserBookings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	userId := c.Param("userId")
	var user entity.User
	//	get the user from the database
	filter := bson.M{"userId": userId}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding user"})
	}
	//getting the user bookingId
	//the booking Id is a slice
	var userBookings []entity.Bookings
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	for _, bookingId := range user.Bookings {
		var userBooking entity.Bookings
		filter := bson.M{"bookingId": bookingId}
		err = bookingCollection.FindOne(ctx, filter).Decode(&userBooking)
		if err != nil {
			panic(err)
		}
		userBookings = append(userBookings, userBooking)
	}
	defer cancel()
	c.JSON(http.StatusOK, userBookings)
}
