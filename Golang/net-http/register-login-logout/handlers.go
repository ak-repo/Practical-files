package main

import (
	"fmt"
	"net/http"
	"time"
)

// register
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type","application/json")

	if r.Method != http.MethodPost {
		status := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", status)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// checking len
	if len(password) < 8 || len(username) < 4 {
		status := http.StatusNotAcceptable
		http.Error(w, "Invalid username/password", status)
		return
	}

	// checking already exist
	if _, ok := users[username]; ok {
		status := http.StatusConflict
		http.Error(w, "User already exist", status)
		return
	}

	// adding into DB
	hashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}
	fmt.Fprintln(w, "User registered sucessfully ")

}

// login
func HandleLogin(w http.ResponseWriter, r *http.Request) {

	//checking method
	if r.Method != http.MethodPost {
		status := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request", status)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	//checking pass/username matching in DB
	user, ok := users[username]
	if !ok || !checkPasswordHash(password, user.HashedPassword) {
		status := http.StatusUnauthorized
		http.Error(w, "Invalid username or password", status)
		return
	}

	//generate session and CSRF token
	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	// set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	//set CSRF cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	//store session and csrf token into DB
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	fmt.Fprintln(w, "Login success")

}

// protected route
func HandleProtectedRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		status := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request", status)
	}

	//autherising user
	if err := Authorize(r); err != nil {
		status := http.StatusUnauthorized
		http.Error(w, "Unauthorized", status)
		return
	}

	username := r.FormValue("username")
	fmt.Fprintln(w, "Welcome user: %s", username)

}

//logout

func HandleLogout(w http.ResponseWriter, r *http.Request) {

	if err := Authorize(r); err != nil {
		status := http.StatusUnauthorized
		http.Error(w, "Unauthorised", status)
		return
	}

	// clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	//clear the tokens from DB
	username := r.FormValue("username")
	user := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user

	fmt.Fprintln(w, "Logout successfully")

}
