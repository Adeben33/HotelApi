package utils

import (
	"github.com/adeben33/HotelApi/internals/entity"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type signinDetails struct {
	Email      string
	FirstName  string
	LastName   string
	Role       string
	Authorized bool
	UserId     string
	jwt.StandardClaims
}

func CreateToken(user entity.User, secretKey string) (string, string, error) {
	var err error
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	claims["authorized"] = true
	claims["userId"] = user.UserId
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["firstName"] = user.FirstName
	claims["lastName"] = user.LastName

	refreshToken := jwt.New(jwt.SigningMethodHS512)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["exp"] = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return tokenString, refreshTokenString, err

	//claims := &signinDetails{
	//	Email:      user.Email,
	//	FirstName:  user.FirstName,
	//	LastName:   user.LastName,
	//	Role:       user.Role,
	//	Authorized: true,
	//	UserId:     user.UserId,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	//	},
	//}
	//refreshClaims := &signinDetails{
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	//	},
	//}
	//token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETKEY))
	//refershToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims).SignedString([]byte(SECRETKEY))
	//if err != nil {
	//	log.Panic(err)
	//}
	//return token, refershToken, nil
}

//func ValidateToken(tokenUrl,secretKey string) (*Token, error) {
//	token, err := jwt.ParseWithClaims(
//		tokenUrl,
//		&Token{},
//		func(token *jwt.Token) (interface{}, error) {
//			return []byte(SecretKey), nil
//		},
//	)
//
//	claims, ok := token.Claims.(*Token)
//	if !ok {
//		return nil, err
//	}
//
//	if claims.ExpiresAt < time.Now().Local().Unix() {
//		return nil, err
//	}
//
//	return claims, err
//
//}

//func NewTokenSrv(secret string) TokenSrv {
//	return &tokensrv{secret}
//}
