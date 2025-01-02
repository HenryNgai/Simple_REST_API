/*
Services for handlers to use
*/
package services

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/henryngai/Simple_REST_API/errors"
	"github.com/henryngai/Simple_REST_API/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	SecretKey string
}

// Constructor like
func NewAuthService(secretKey string) *AuthService {
	return &AuthService{SecretKey: secretKey}
}

// Expected to return the JWT after comparing hashed passwords
func (service *AuthService) Authenticate(email string, password string) (string, error) {
	// Check if user exists
	user, noUserErr := models.GetUserByEmail(email)
	if noUserErr != nil {
		return "", noUserErr
	}

	// Compares a hashed password with its plain-text equivalent.
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if passwordErr != nil {
		return "", passwordErr
	}

	// Generate JWT and return
	JWT, JWTError := service.GenerateWebToken(email)
	if JWTError != nil {
		return "", JWTError
	}
	return JWT, nil

}

func (service *AuthService) Register(email string, password string) error {
	_, err := models.GetUserByEmail(email)
	if err == nil { // We received no error, meaning we already have this user
		return errors.ErrUserAlreadyyExists
	}

	// Salt and hash
	hashedPassword, hashPasswordErr := service.hashPassword(password)
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
func (service *AuthService) hashPassword(password string) (string, error) {
	// bcrypt automatically generates salt internally
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil

}

// Generate JWT Token
func (service *AuthService) GenerateWebToken(email string) (string, error) {
	// Define claims
	claims := jwt.MapClaims{
		"email": email,                            // Custom
		"exp":   time.Now().Add(time.Hour).Unix(), // 1 hour expiration
		"iat":   time.Now().Unix(),                // Time issued
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString([]byte(service.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

// Generate secret key of length for JWT
func GenerateSecretKey(length int) string {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Println("Failed to generate secret key of length " + string(length))
		return ""
	}
	return hex.EncodeToString(key)
}

// ValidateJWT validates the JWT token and returns claims if valid
// interface{} allows for any type
func (service *AuthService) ValidateJWT(tokenString string) (map[string]interface{}, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrSigningMethod
		}
		return []byte(service.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Optional: Check expiration claim
		if exp, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)
			if expTime.Before(time.Now()) {
				return nil, errors.ErrExpToken
			}
		}
		// Return claims as a map
		return claims, nil
	}

	return nil, errors.ErrInvalidToken
}
