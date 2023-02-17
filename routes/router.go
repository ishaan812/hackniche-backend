package routes

import (
	"fmt"
	"log"
	"net/http"
	"service/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/register", controllers.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("GET")

	r.HandleFunc("/participant", controllers.CreateParticipant).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllParticipants", controllers.GetAllParticipants).Methods("GET")
	r.HandleFunc("/participant/{id}", controllers.GetParticipantByID).Methods("GET")
	r.HandleFunc("/participant/{id}", controllers.UpdateParticipant).Methods("PUT")
	r.HandleFunc("/participant/{id}", controllers.DeleteParticipant).Methods("DELETE")

	r.HandleFunc("/hackathon", controllers.CreateHackathon).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllHackathons", controllers.GetAllHackathons).Methods("GET")
	r.HandleFunc("/hackathon/{id}", controllers.GetHackathonByID).Methods("GET")
	r.HandleFunc("/hackathon/{id}", controllers.UpdateHackathon).Methods("PUT")
	r.HandleFunc("/hackathon/{id}", controllers.DeleteHackathon).Methods("DELETE")

	r.HandleFunc("/organizer", controllers.CreateOrganizer).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllOrganizers", controllers.GetAllOrganizers).Methods("GET")
	r.HandleFunc("/organizer/{id}", controllers.GetOrganizerByID).Methods("GET")
	r.HandleFunc("/organizer/{id}", controllers.UpdateOrganizer).Methods("PUT")
	r.HandleFunc("/organizer/{id}", controllers.DeleteOrganizer).Methods("DELETE")

	r.HandleFunc("/team", controllers.CreateTeam).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllTeams", controllers.GetAllTeams).Methods("GET")
	r.HandleFunc("/team/{id}", controllers.GetTeamByID).Methods("GET")
	r.HandleFunc("/team/{id}", controllers.UpdateTeam).Methods("PUT")
	r.HandleFunc("/team/{id}", controllers.DeleteTeam).Methods("DELETE")

	r.HandleFunc("/addTeamToHackathon", controllers.AddTeamToHackathon).Methods("POST", "OPTIONS")
	r.HandleFunc("/getAllHackathonTeams", controllers.GetAllHackathonTeams).Methods("GET")
	r.HandleFunc("/getTeamsbyHackathon/{id}", controllers.GetTeamsByHackathon).Methods("GET")
	r.HandleFunc("/UpdateHackathonTeam/{id}", controllers.UpdateHackathonTeams).Methods("PUT")
	r.HandleFunc("/DeleteHackathonTeam/{id}", controllers.DeleteHackathonTeams).Methods("DELETE")

	fmt.Print("Server running on localhost:9000\n")
	serverErr := http.ListenAndServe("localhost:9000", handlers.CORS(headers, methods, origins)(r))
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
