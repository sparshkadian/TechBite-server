package models

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	db "techBite/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID					int         `json:"id"`
	Name 				string  	`json:"name"`
	Email				string  	`json:"email"`
	Password 			string  	`json:"password"`
}

type NewUser struct {
	ID   	 int  	`json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

var dbCon *sql.DB

func init() {
	dbCon = db.ReturnDB()
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) UseSignup(w http.ResponseWriter) *NewUser{	
	rand.NewSource(time.Now().UnixNano())
	pseudoRandomId :=  rand.Intn(10000)
	user.ID = pseudoRandomId

	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if hashError != nil {
		log.Fatal("Error hashing password", hashError)
		return nil
	}
	user.Password = string(hashedPassword)

	_, err := dbCon.Exec("INSERT INTO `blogwebsite`.`users` (id, name, email, password) VALUES(?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("Error Creating new User", err)
		return nil
	}
	
	newUser := NewUser{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	return &newUser
}

func (user *User) UseLogin(w http.ResponseWriter) *NewUser{
	email := user.Email
	password := user.Password
	if len(email) == 0 || len(password) == 0 {
		http.Error(w, "Both Fields are required", http.StatusBadRequest)
		return nil
	}
	err := dbCon.QueryRow("SELECT * FROM `blogwebsite`.`user` WHERE (email = ?)", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No such user exists", http.StatusBadRequest)
			return nil
		} else {
			log.Println("Error querying database:", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return nil
		}
	}
	comparePasswords := CheckPasswordHash(user.Password, password)
	if !comparePasswords {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return nil
	}
	returningUser := NewUser{ID: user.ID, Name: user.Name, Email: user.Email}
	return &returningUser
}