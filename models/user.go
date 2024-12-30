/*
Model for interacting with database.
Specifically users
*/

package models

import (
	"database/sql"
	"log"

	"github.com/henryngai/Simple_REST_API/database"
	"github.com/henryngai/Simple_REST_API/errors"
)

type User struct {
	Email          string
	HashedPassword string
}

func GetUserByEmail(email string) (User, error) {
	var user User
	query := "SELECT email, hashed_password FROM users WHERE email = ?"
	row := database.DB.QueryRow(query, email)
	err := row.Scan(&user.Email, &user.HashedPassword) // Assigns to struct.
	if err == sql.ErrNoRows {
		log.Println("Error no users found")
		return user, errors.ErrUserNotFound
	}

	return user, nil
}

func RegisterUser(email string, hashedPassword string) error {
	query := "INSERT INTO users (email, hashed_password) VALUES (?, ?)"
	_, err := database.DB.Exec(query, email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
