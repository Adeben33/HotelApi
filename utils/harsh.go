package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HarshPassword(password string) (string, error) {
	harsh, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	return string(harsh), nil
}

func VerifyPassword(harsedPassword, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(harsedPassword), []byte(loginPassword))
	if err != nil {
		msg := "Incorrect Password"
		log.Panic(msg)
		return err
	}
	return nil
}
