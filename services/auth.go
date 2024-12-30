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

// Expected to return the JWT after comparing hashed passwords
func Authenticate(email string, password string) (string, error) {
	// Check if user exists
	user, noUserErr := models.GetUserByEmail(email)
	if noUserErr != nil {
		return "", noUserErr
	}

	// Compares a hashed password with its plain-text equivalent.
	log.Println(user.HashedPassword)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if passwordErr != nil {
		return "", passwordErr
	}

	// TODO Generate JWT
	log.Println("Success, JWT time")
	return "JWT", nil

}

func Register(email string, password string) error {
	_, err := models.GetUserByEmail(email)
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
