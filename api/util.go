package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func mapRequestBody[T any](r *http.Request) (T, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	var body T
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return body, err
	}

	bodyString := string(bodyBytes)
	log.Println(bodyString)

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return body, err
	}

	return body, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
	}

	jwtKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	claims := &jwt.StandardClaims{}
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", fmt.Errorf("invalid JWT signature")
		}
		return "", fmt.Errorf("could not parse JWT")
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid JWT")
	}

	return claims.Subject, nil
}
