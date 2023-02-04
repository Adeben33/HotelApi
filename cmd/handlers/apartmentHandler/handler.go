package apartmentHandler

import (
	"context"
	"github.com/adeben33/HotelApi/internals/dataBaseStore/mongoDBConnection"
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jaswdr/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var validate = validator.New()
var apartmentCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "apartment")
var reviewCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "review")

func GetApartment(c *gin.Context) {
	var apartment entity.ApartmentRes
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	apartmentId := c.Param("apartmentId")
	filter := bson.M{"apartment_id": apartmentId}

	findErr := apartmentCollection.FindOne(ctx, filter).Decode(apartment)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	c.JSON(http.StatusOK, apartment)
}

func GetAllApartment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var apartments []entity.ApartmentRes

	filter := bson.M{}
	count, countErr := apartmentCollection.CountDocuments(ctx, filter)
	if countErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": countErr.Error()})
		return
	}
	findOption := options.Find()

	//search querry, this can contain apartment names, amenities and so on

	search := c.Query("search")
	location := c.Query("location")

	if search != " " || location != " " {
		filter = bson.M{
			"$or": []bson.M{
				{"name": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},
				{"apartment_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},
				{"address": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},
				{"address": bson.M{
					"$regex": primitive.Regex{
						Pattern: location,
						Options: "i",
					},
				}},
			},
		}
	}

	//Sort
	sort := c.Query("sort")

	if sort != " " {
		if sort == "asc" {
			findOption.SetSort(
				bson.M{
					"apartment_id": 1,
				})
		} else if sort == "desc" {
			findOption.SetSort(bson.M{
				"apartment_id": -1,
			})
		}
	}

	var perPage int64 = 9
	page, _ := strconv.Atoi(c.Query("page"))

	if page < 1 {
		page = 1
	}

	skippingLimit := int64(page-1) * perPage
	findOption.SetSkip(skippingLimit)
	findOption.SetLimit(perPage)

	cursor, findErr := apartmentCollection.Find(ctx, filter, findOption)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var apartment entity.ApartmentRes
		cursor.Decode(&apartment)
		apartments = append(apartments, apartment)
	}

	searchCount, _ := apartmentCollection.CountDocuments(ctx, filter)

	c.JSON(http.StatusOK, gin.H{
		"data":           apartments,
		"All data count": count,
		"Search Count":   searchCount,
		"page":           page,
	})
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

func FakeApartment(c *gin.Context) {
	//	This is to create a set of apartment
	//ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//defer cancel()
	//var users []entity.User
	fake := faker.New()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	for i := 0; i <= 50; i++ {
		var apartment entity.Apartment
		apartment.ID = primitive.NewObjectID()
		apartment.ApartmentId = apartment.ID.Hex()
		apartment.Name = fake.Person().FirstName() + " Hotel"
		apartment.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		apartment.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		apartment.NumberofRooms = uint8(rand.Intn(6))
		apartment.Price = uint16(rand.Intn(10000))
		apartmentCollection.InsertOne(ctx, apartment)
	}
	defer cancel()
	c.JSON(http.StatusOK, gin.H{"success": "success"})
}

func GetAllApartmentWithAmenities(c *gin.Context) {
	var amenities []string
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if v, ok := c.GetQueryArray("amenities"); ok {
		amenities = v
	}

	//the in operator would be used here
	filter := bson.M{"amenities": bson.M{"$in": amenities}}

	//pagination will be added because the apartment might be too large

	findOptions := options.Find()

	page, _ := strconv.Atoi(c.Param("page"))
	if page < 0 {
		page = 1
	}
	var perPage int64 = 9
	skippingLimit := int64(page-1) * perPage

	findOptions.SetSkip(skippingLimit)
	findOptions.SetLimit(perPage)

	cursor, findErr := apartmentCollection.Find(ctx, filter, findOptions)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	var results []bson.M
	err := cursor.All(ctx, &results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)

}

func GetAllReviews(c *gin.Context) {
	var reviews []entity.ReviewsRes
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	apartmentId := c.Param("apartmentId")

	//	review in the apartment struct contains reviewId related to a specific apartment
	//All reviews have apartmentId in their struct
	filter := bson.M{"apartment_id": apartmentId}

	cursor, findErr := reviewCollection.Find(ctx, filter)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": findErr.Error()})
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var review entity.ReviewsRes
		cursor.Decode(&review)
		reviews = append(reviews, review)
	}

	c.JSON(http.StatusOK, reviews)

}

func GetRentedApartment(c *gin.Context) {
	ctx, cancel := context.
		c.Param("userId")
}
