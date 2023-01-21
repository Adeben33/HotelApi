package user

import (
	"github.com/adeben33/HotelApi/internals/entity/userEntity"
	"github.com/adeben33/HotelApi/internals/repository/userRepo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func SignUpUser(user userEntity.CreateUser) error {
	//	check if the user already exists

	//	harsh the password
	harsh, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Panic("Unable to Harsh Password")
	}
	user.Password = string(harsh)
	user.DateCreated = time.Now()
	//	Send it to the database
	var repo userRepo.UserRepo
	_, err = repo.CreateUser(&user)
	//	Send the response
	return nil
}
