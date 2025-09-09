package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// sessions
// cookies
// CSRF tokens
// username/ password

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// fake data base
var users = map[string]Login{}

func main() {

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/register", HandleRegister).Methods("POST")
	router.HandleFunc("/login", HandleLogin).Methods("POST")
	router.HandleFunc("/protected", HandleProtectedRoute).Methods("POST")
	router.HandleFunc("/logout", HandleLogout).Methods("POST")

	//server
	http.ListenAndServe(":8080", router)

}
