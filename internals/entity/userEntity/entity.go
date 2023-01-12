package userEntity

type CreateUser struct {
	UserId        uint   `json:"user_id" gorm:"primary_key""`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Email         string `json:"email" validate:"email" gorm:""`
	Phone         string `json:"phone" validate:"e164"`
	Password      string `json:"password" validate:"required,min=6"`
	Gender        string `json:"gender" validate:"required, oneof='male' 'female' Others"`
	DateOfBirth   string `json:"dateOfBirth" validate:"omitempty"`
	AccountStatus string `json:"accountStatus" validate:"required oneof='active' 'paused' 'deactivated'"`
	DateCreated   string `json:"date_created"`
	role          string `json:"role" validate:"required"`
}

type Apartment struct {
	ApartmentId     int    `json:"apartment_id" gorm:"primary_key"`
	ApartmentName   string `json:"apartment_name" validate:"required"`
	ApartmentStatus string `json:"apartment_status" validate:"required oneof='available' 'booked' 'paused'"`
	DaysBookedfor   int    `json:"days_bookedfor" validate:"omitempty""`
	Occupant        string `json:"user_id" validate:"required"`
	CreatedBy       string `json:"created_by" validate:"required"`
	Picture         string `json:"picture" validate:"omitempty"`
}
