/*
Model for interacting with database.
Specifically users
*/

package models

import (
	"log"

	"github.com/henryngai/Simple_REST_API/errors"
)

func DoesUserExist(email string) error {
	log.Println("Test - User not found ", email)
	return errors.ErrUserNotFound
}

func RegisterUser(email string, password string) error {
	log.Println("Registering user in model success")
	return nil
}
