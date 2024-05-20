package routes

import (
	"techBite/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	a := r.PathPrefix("/api/auth").Subrouter()
	a.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	a.HandleFunc("/login", controllers.Login).Methods("POST")
}