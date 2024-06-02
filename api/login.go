package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (s *ApiServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		body, err := mapRequestBody[loginBody](r)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			return fmt.Errorf("error logging")
		}
		user, err := s.store.GetUserByName(body.UserName)
		if err != nil {
			log.Printf("Error logging: %v", err)
			return fmt.Errorf("error logging")
		}

		if !user.HashedPassword.Valid {
			err := s.createFirstPassword(body, user.Id)
			if err != nil {
				return fmt.Errorf("error logging")
			}
			return nil
		}
		log.Println("Checking password")
		if CheckPasswordHash(body.Password, user.HashedPassword.String) {
			log.Println("Password match")

			jwtToken, err := GenerateJWT(strconv.Itoa(user.Id))
			if err != nil {
				log.Println(err.Error())
				return fmt.Errorf("error logging")
			}
			return WriteJson(w, http.StatusAccepted, jwtRespone{UserName: user.Name, JwtToken: jwtToken})
		}

		return fmt.Errorf("error logging")
	}
	return fmt.Errorf("method not supported")
}

func (s *ApiServer) createFirstPassword(body loginBody, userId int) error {
	log.Println("First login setting up password")
	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		return err
	}
	err = s.store.SetUserPassword(userId, hashedPassword)
	return err
}
