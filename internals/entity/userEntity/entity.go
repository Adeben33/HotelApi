package userEntity

type CreateUser struct {
	UserId        string `json:"user_id"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Email         string `json:"email" validate:"email"`
	Phone         string `json:"phone" validate:"e164"`
	Password      string `json:"password" validate:"required,min=6"`
	Gender        string `json:"gender" validate:"required, oneof='male' 'female' Others"`
	DateOfBirth   string `json:"dateOfBirth" validate:"omitempty"`
	AccountStatus string `json:"accountStatus" validate:"required oneof='active' 'paused' 'deactivated'"`
	BookedApartment
}

//type createUser struct {
//	UserId        string `json:"user_id"`
//	FirstName     string `json:"first_name" validate:"required"`
//	LastName      string `json:"last_name"  validate:"required"`
//	Email         string `json:"email" validate:"email"`
//	Phone         string `json:"phone" validate:"required"`
//	Password      string `json:"password" validate:"required,min=6"`
//	Gender        string `json:"gender"  validate:"required,oneof='Male' 'Female' 'Others'"`
//	DateOfBirth   string `json:"date_of_birth" validate:"required"`
//	AccountStatus string `json:"account_status"`
//	PaymentStatus string `json:"payment_status"`
//	DateCreated   string `json:"date_created"`
//}
