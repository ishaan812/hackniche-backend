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


	fmt.Print("Server running on localhost:9000\n")
	serverErr := http.ListenAndServe("localhost:9000", handlers.CORS(headers, methods, origins)(r))
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
