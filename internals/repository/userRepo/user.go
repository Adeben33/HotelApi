package userRepo

import (
	"github.com/adeben33/HotelApi/internals/dataBaseStore/postgresDB"
	"github.com/adeben33/HotelApi/internals/entity/userEntity"
)

type userRepo struct {
	DB postgresDB.DatabaseServer
}
type UserRepo interface {
	CreateUser(user *userEntity.CreateUser) (userEntity.CreateUserReq, error)
}

func (u *userRepo) CreateUser(user *userEntity.CreateUser) (userEntity.CreateUserReq, error) {
	DB := u.DB.GetConn()
	_ = DB.Create(&user)

	//Build user Response
	UserResponse := userEntity.CreateUserReq{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		AccountStatus: user.AccountStatus,
		Role:          user.Role,
	}
	return UserResponse, nil
}
