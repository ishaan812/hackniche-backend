package controllers

import (
	"encoding/json"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func AddTeamToHackathon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var HackathonTeams database.HackathonTeams
	json.NewDecoder(r.Body).Decode(&HackathonTeams)
	err := dbconn.Create(&HackathonTeams)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Organizer Created")
	}
}

func GetAllHackathonTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var hackathon []database.HackathonTeams
	// year := r.URL.Query().Get("year")
	// division := r.URL.Query().Get("division")
	// batch := r.URL.Query().Get("batch")
	// department := r.URL.Query().Get("department")
	// if year == "" && division == "" && batch == "" && department == "" {
	// 	dbconn.Find(&hackathon)
	// 	json.NewEncoder(w).Encode(&hackathon)
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
	dbconn.Find(&hackathon)
	json.NewEncoder(w).Encode(&hackathon)
	// }
}

func GetTeamsByHackathon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var hackathonteams []database.HackathonTeams
	err := dbconn.Where("hackathon_id = ?", params["id"]).Find(&hackathonteams).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&hackathonteams)
	}
}

func DeleteHackathonTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var hackathon database.HackathonTeams
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&hackathon).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&hackathon)
		json.NewEncoder(w).Encode("HackathonTeams Deleted")
	}
}

func UpdateHackathonTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var hackathon database.HackathonTeams
	err := dbconn.Where("id = ?", params["id"]).First(&hackathon).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&hackathon)
	dbconn.Save(&hackathon)
	json.NewEncoder(w).Encode(&hackathon)
}
