package entity

type Reviews struct {
	ID          string `bson:"_id" json:"_id"`
	ApartmentId string `bson:"apartment_id" json:"apartmentId"`
	UserId      string `json:"userId" bson:"user_id"`
	Rating      uint8  `json:"rating" bson:"rating"`
	ReviewId    string `json:"review" bson:"review"`
}

type ReviewsRes struct {
	Rating   uint8  `json:"rating" bson:"rating"`
	Review   string `json:"review" bson:"review"`
	ReviewId string `json:"reviewId" bson:"review_id"`
}
