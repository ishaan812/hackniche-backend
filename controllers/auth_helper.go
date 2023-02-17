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

var tokenValidityDuration = 60 * 24 * time.Minute

func LoginUser(Organizer *database.Organizer) (*database.Organizer, error) {
	var organizer database.Organizer
	err := dbconn.Where("email = ?", Organizer.Email).First(&organizer).Error
	if err != nil {
		fmt.Println("ERROR: sapid does not exist")
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(organizer.Password), []byte(Organizer.Password))
	if err != nil {
		fmt.Println("ERROR: Wrong Password Entered")
		return nil, err
	}
	fmt.Println("INFO: ", Organizer.Email, " logged in")
	Organizer = &organizer
	return &organizer, nil
}

func RegisterUser(Organizer *database.Organizer) error {
	err := dbconn.Where("email = ?", Organizer.Email).First(&Organizer).Error
	if err != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(Organizer.Password), 14)
		if err != nil {
			return err
		} else {
			Organizer.Password = string(bytes)
		}
		dbconn.Create(&Organizer)
		fmt.Println("INFO: New Organizer ", Organizer.Email, " has been registered")
		return nil
	}
	fmt.Println("ERROR: Organizer ", Organizer.Email, " already exists")
	return errors.New("user already exists")
}

func LogoutUser(c *http.Cookie) error {
	fmt.Println("INFO: Logged out")
	c.Expires = time.Now().Add(-1 * time.Hour)
	return nil
}

func CreateJWT(Organizer *database.Organizer) (*http.Cookie, error) {
	expirationTime := time.Now().Add(tokenValidityDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":       Organizer.Name,
		"email":      Organizer.Email,
		"expiresat":  expirationTime,
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
	fmt.Println("INFO: JWT of ", Organizer.Email, " generated with expiration time ", expirationTime)
	return JWTCookie, nil
}

func ValidateJWT(c *http.Cookie, jwtKey string) (*Claims, error) {
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

func RefreshJWT(claims *Claims, jwtKey string) (*http.Cookie, error) {
	expirationTime := time.Now().Add(tokenValidityDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":       claims.Name,
		"email":      claims.Email,
		"expiresat":  expirationTime,
		"department": claims.Department,
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
