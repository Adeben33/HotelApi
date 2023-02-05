package apartmentHandler

import (
	"context"
	"fmt"
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
var bookingCollection = mongoDBConnection.OpenCollection(mongoDBConnection.Client, "booking")

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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	renterId := c.Param("userId")
	filter := bson.M{"renter_id": renterId}

	cursor, findErr := apartmentCollection.Find(ctx, filter)
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

func GetApartmentWithNumberofBedrooms(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var apartments []entity.ApartmentRes

	numberOfRooms, _ := strconv.Atoi(c.Query("numberOfRooms"))
	if numberOfRooms < 1 {
		msg := fmt.Sprintf("Invalid room numbers")
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	page, _ := strconv.Atoi(c.Param("page"))
	if page < 1 {
		page = 1
	}
	filter := bson.M{"numberof_rooms": numberOfRooms}
	findOptions := options.Find()
	var perPage int64 = 9

	skippingLimit := int64(page-1) * perPage
	findOptions.SetSkip(skippingLimit)
	findOptions.SetLimit(perPage)

	cursor, findErr := apartmentCollection.Find(ctx, filter, findOptions)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	count, _ := apartmentCollection.CountDocuments(ctx, filter)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var apartment entity.Apartment
		var apartmentRes entity.ApartmentRes
		cursor.Decode(&apartment)

		apartmentRes.NumberofRooms = apartment.NumberofRooms
		apartmentRes.Price = apartment.Price
		apartmentRes.Images = apartment.Images
		apartmentRes.Name = apartment.Name
		apartmentRes.UpdatedAt = apartment.UpdatedAt
		apartmentRes.CreatedAt = apartment.CreatedAt
		apartmentRes.Address = apartment.Address
		apartmentRes.Amenities = apartment.Amenities
		apartmentRes.Review = apartment.Review

		apartments = append(apartments, apartmentRes)
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":  apartments,
		"count": count,
	})
}

func GetAllApartmentBookings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var apartment entity.Apartment

	apartmentId := c.Param("apartmentId")

	filter := bson.M{"apartment_id": apartmentId}
	findErr := apartmentCollection.FindOne(ctx, filter).Decode(apartment)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}
	var bookings []entity.Bookings
	for _, bookingId := range apartment.BookingsId {
		var booking entity.Bookings
		findErr := bookingCollection.FindOne(ctx, bson.M{"booking_id": bookingId}).Decode(&booking)
		if findErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		}
		bookings = append(bookings, booking)
	}
	c.JSON(http.StatusOK, bookings)

}

func GetApartmentAmenities(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var apartment entity.Apartment

	apartmentId := c.Param("apartmentId")

	filter := bson.M{"apartment_id": apartmentId}
	findErr := apartmentCollection.FindOne(ctx, filter).Decode(apartment)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	c.JSON(http.StatusOK, apartment.Amenities)
}

func UpdateApartmentAmenities(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var amenities []string
	apartmentId := c.Param("apartmentId")

	if err := c.ShouldBindJSON(&amenities); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"apartment_id": apartmentId}

	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters:             nil,
		BypassDocumentValidation: nil,
		Collation:                nil,
		Comment:                  nil,
		Hint:                     nil,
		Upsert:                   &upsert,
		Let:                      nil,
	}

	update := bson.D{{
		"$push", bson.D{{
			"amenities", amenities,
		},
		},
	}}

	updateResult, updateErr := apartmentCollection.UpdateOne(ctx, filter, update, &opts)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResult)
}

func ApartmentAverageRating(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var reviews []entity.Reviews
	var apartment entity.Apartment

	apartmentId := c.Param("apartmentId")

	findErr := apartmentCollection.FindOne(ctx, bson.M{"apartmentId": apartmentId}).Decode(&apartment)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, findErr.Error())
		return
	}

	filter := bson.M{"apartment_id": apartmentId}

	cursor, findErr := reviewCollection.Find(ctx, filter)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, findErr.Error())
		return
	}
	defer cursor.Close(ctx)
	var rating uint8 = 0
	for cursor.Next(ctx) {
		var review entity.Reviews
		cursor.Decode(&review)
		rating += review.Rating
		reviews = append(reviews, review)
	}
	average := rating / uint8(len(reviews))
	c.JSON(http.StatusOK, gin.H{
		"Apartment":      apartment.Name,
		"Average Review": average,
	})
}

func ApartmentLastBooking(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//This will be sorted by the id in the decreasing order and the first will be last booking
	var apartment entity.Apartment
	filter := bson.M{"apartment": apartment}
	findOption := options.Find()
	findOption.SetSort(
		bson.M{
			"_id": -1,
		},
	)
	findOption.SetLimit(1)

	cursor, findErr := bookingCollection.Find(ctx, filter, findOption)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, findErr.Error())
		return
	}
	defer cursor.Close(ctx)
	var bookings []entity.Bookings
	for cursor.Next(ctx) {
		var booking entity.Bookings
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, gin.H{
		"apartment Name": apartment.Name,
		"last Booking":   bookings,
	})
}
