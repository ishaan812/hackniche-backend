package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var team database.Team
	json.NewDecoder(r.Body).Decode(&team)
	err := dbconn.Create(&team)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Team Created")
	}
}

func GetAllTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var team []database.Team
	// year := r.URL.Query().Get("year")
	// division := r.URL.Query().Get("division")
	// batch := r.URL.Query().Get("batch")
	// department := r.URL.Query().Get("department")
	// if year == "" && division == "" && batch == "" && department == "" {
	// 	dbconn.Find(&team)
	// 	json.NewEncoder(w).Encode(&team)
	// } else {
	// 	querystring := ""
	// 	if year != "" {
	// 		querystring += "year = " + year
	// 	}
	// 	if division != "" {
	// 		if querystring != "" {
	// 			querystring = querystring + " AND "
	// 		}
	// 		querystring = querystring + "division = '" + division + "'"
	// 	}
	// 	if batch != "" {
	// 		if querystring != "" {
	// 			querystring = querystring + " AND "
	// 		}
	// 		querystring = querystring + "batch = " + batch
	// 	}
	// 	if department != "" {
	// 		if querystring != "" {
	// 			querystring = querystring + " AND "
	// 		}
	// 		querystring = querystring + "department = '" + department + "'"
	// 	}
	// fmt.Println(querystring)
	dbconn.Find(&team)
	json.NewEncoder(w).Encode(&team)
	// }
}

func GetTeamByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var team database.Team
	err := dbconn.Where("id = ?", params["id"]).First(&team).Error
	fmt.Println(team)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&team)
	}
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var team database.Team
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&team).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&team)
		json.NewEncoder(w).Encode("Team Deleted")
	}
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var team database.Team
	err := dbconn.Where("id = ?", params["id"]).First(&team).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&team)
	dbconn.Save(&team)
	json.NewEncoder(w).Encode(&team)
}
