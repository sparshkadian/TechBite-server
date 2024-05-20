package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"techBite/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error Reading request body", err)
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal("Error unmarshaling data", err)
		return
	}

	newUser := user.UseSignup(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return 
	}
	
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		http.Error(w, "Error unmarshaling data", http.StatusBadRequest)
		return
	}

	userDetails := user.UseLogin(w)
	if userDetails == nil {
        http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDetails)
}