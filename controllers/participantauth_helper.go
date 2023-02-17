package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"service/database"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserParticipation(Participant *database.Participant) (*database.Participant, error) {
	var organizer database.Participant
	err := dbconn.Where("email = ?", Participant.Email).First(&organizer).Error
	if err != nil {
		fmt.Println("ERROR: sapid does not exist")
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(organizer.Password), []byte(Participant.Password))
	if err != nil {
		fmt.Println("ERROR: Wrong Password Entered")
		return nil, err
	}
	fmt.Println("INFO: ", Participant.Email, " logged in")
	Participant = &organizer
	return &organizer, nil
}

func RegisterUserParticipation(Participant *database.Participant) error {
	err := dbconn.Where("email = ?", Participant.Email).First(&Participant).Error
	if err != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(Participant.Password), 14)
		if err != nil {
			return err
		} else {
			Participant.Password = string(bytes)
		}
		dbconn.Create(&Participant)
		fmt.Println("INFO: New Participant ", Participant.Email, " has been registered")
		return nil
	}
	fmt.Println("ERROR: Participant ", Participant.Email, " already exists")
	return errors.New("user already exists")
}

func LogoutUserParticipation(c *http.Cookie) error {
	fmt.Println("INFO: Logged out")
	c.Expires = time.Now().Add(-1 * time.Hour)
	return nil
}

func CreateJWTParticipation(Participant *database.Participant) (*http.Cookie, error) {
	expirationTime := time.Now().Add(tokenValidityDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      Participant.Name,
		"email":     Participant.Email,
		"expiresat": expirationTime,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		fmt.Println("ERROR: Error in JWT Key")
		return nil, err
	}
	JWTCookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	fmt.Println("INFO: JWT of ", Participant.Email, " generated with expiration time ", expirationTime)
	return JWTCookie, nil
}

func ValidateJWTParticipation(c *http.Cookie, jwtKey string) (*Claims, error) {
	var claims Claims
	tknStr := c.Value
	tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		fmt.Println("ERROR: JWT of ", claims.Email, " is invalid")
		return nil, err
	}
	fmt.Println("INFO: JWT of ", claims.Email, " validated")
	return &claims, nil
}

func RefreshJWTParticipation(claims *Claims, jwtKey string) (*http.Cookie, error) {
	expirationTime := time.Now().Add(tokenValidityDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      claims.Name,
		"email":     claims.Email,
		"expiresat": expirationTime,
	})
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		fmt.Println("ERROR: Error while creating JWT.")
		return nil, err
	}
	JWTCookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	fmt.Println("INFO: JWT of ", claims.Email, " refreshed to ", expirationTime)
	return JWTCookie, nil
}
