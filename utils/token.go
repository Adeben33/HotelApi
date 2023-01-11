package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Token struct {
	Email string
	Id    string
	jwt.StandardClaims
}

type TokenSrv interface {
	CreateToken(id, email string) (string, error)
	ValidateToken(token string) (*Token, error)
}
type tokensrv struct {
	SecretKey string
}

func (t *tokensrv) CreateToken(id, email string) (string, error) {
	var err error
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email

	tokenString, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return tokenString, err
}

func (t *tokensrv) ValidateToken(tokenUrl string) (*Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenUrl,
		&Token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.SecretKey), nil
		},
	)

	claims, ok := token.Claims.(*Token)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}

	return claims, err

}

func NewTokenSrv(secret string) TokenSrv {
	return &tokensrv{secret}
}
