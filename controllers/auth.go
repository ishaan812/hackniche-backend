package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"service/database"
)

// AUTHENTICATION
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Organizer database.Organizer
	json.NewDecoder(r.Body).Decode(&Organizer)
	faculty, err := LoginUser(&Organizer)
	fmt.Println(Organizer)
	if err != nil {
		json.NewEncoder(w).Encode("AuthError")
	} else {
		JWTCookie, err := CreateJWT(&Organizer)
		if err != nil {
			fmt.Println("Error while creating JWT.")
			json.NewEncoder(w).Encode("JWTError")
		} else {
			http.SetCookie(w, JWTCookie)
			json.NewEncoder(w).Encode(faculty)
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Organizer database.Organizer
	json.NewDecoder(r.Body).Decode(&Organizer)
	err := RegisterUser(&Organizer)
	if err != nil {
		fmt.Println("Error while registering user")
		json.NewEncoder(w).Encode("RegisterError")
	} else {
		fmt.Println("User registered successfully")
		json.NewEncoder(w).Encode("RegisterSuccess")
	}
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	claims, err := ValidateJWT(c, jwtKey)
	fmt.Println(claims)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	JWTCookie, err := RefreshJWT(claims, jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		http.SetCookie(w, JWTCookie)
		json.NewEncoder(w).Encode("JWTRefreshed")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	LogoutUser(c)
	http.SetCookie(w, c)
}
