/*
Services for handlers to use
*/
package services

import (
	"log"

	"github.com/henryngai/Simple_REST_API/errors"
	"github.com/henryngai/Simple_REST_API/models"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(email string, password string) {
	log.Println("Asked for service authentication")
}

func Register(email string, password string) error {
	log.Println("Asked for service registration")
	err := models.DoesUserExist(email)
	if err == nil { // We received no error, meaning we already have this user
		return errors.ErrUserAlreadyyExists
	}

	// Salt and hash
	hashedPassword, hashPasswordErr := hashPassword(password)
	if hashPasswordErr != nil {
		return hashPasswordErr
	}

	// Call to dependency layer to register user to db
	registerErr := models.RegisterUser(email, hashedPassword)
	if registerErr != nil {
		return registerErr
	}

	log.Println("Successfully registered user")
	return nil

}

// Function used to hash passwords
func hashPassword(password string) (string, error) {
	// bcrypt automatically generates salt internally
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil

}
