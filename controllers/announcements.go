package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service/database"

	"github.com/gorilla/mux"
)

func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Announcement database.Announcement
	json.NewDecoder(r.Body).Decode(&Announcement)
	err := dbconn.Create(&Announcement)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error)
	} else {
		json.NewEncoder(w).Encode("Announcement Created")
	}
}

func GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Announcement []database.Announcement
	// year := r.URL.Query().Get("year")
	// division := r.URL.Query().Get("division")
	// batch := r.URL.Query().Get("batch")
	// department := r.URL.Query().Get("department")
	// if year == "" && division == "" && batch == "" && department == "" {
	// 	dbconn.Find(&Announcement)
	// 	json.NewEncoder(w).Encode(&Announcement)
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
	dbconn.Find(&Announcement)
	json.NewEncoder(w).Encode(&Announcement)
	// }
}

func GetAnnouncementByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Announcement database.Announcement
	err := dbconn.Where("id = ?", params["id"]).First(&Announcement).Error
	fmt.Println(Announcement)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		json.NewEncoder(w).Encode(&Announcement)
	}
}

func DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var Announcement database.Announcement
	params := mux.Vars(r)
	err := dbconn.Where("id = ?", params["id"]).First(&Announcement).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	} else {
		dbconn.Delete(&Announcement)
		json.NewEncoder(w).Encode("Announcement Deleted")
	}
}

func UpdateAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Announcement database.Announcement
	err := dbconn.Where("id = ?", params["id"]).First(&Announcement).Error
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	json.NewDecoder(r.Body).Decode(&Announcement)
	dbconn.Save(&Announcement)
	json.NewEncoder(w).Encode(&Announcement)
}
