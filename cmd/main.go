package main

import (
	"fmt"
	"log"
	"net/http"
	"techBite/routes"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r)
	
	fmt.Println("Server started at port:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server", err)
	}
}