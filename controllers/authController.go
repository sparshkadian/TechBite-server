package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"techBite/models"
)

// func SignUp(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal("Error Reading request body", err)
// 		return
// 	}
// 	err = json.Unmarshal(data, &user)
// 	if err != nil {
// 		log.Fatal("Error unmarshaling data", err)
// 		return
// 	}
// 	newUser := user.UseSignup(w)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newUser)
// }

func Login(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading request body", err)
		return 
	}
	var user models.User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		log.Fatal("Error unmarshaling data", err)
		return
	}
	userDetails := user.UseLogin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDetails)
}