package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// ==================================================
// Models
// ==================================================
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ==================================================
// Globals
// ==================================================
var secret []byte

// Dummy user (in real-world apps, fetch from DB)
var dummyUser = User{
	Username: "jwt",
	Password: "$2a$10$LZlRjN69lh97u8lP9XAsSOG6UO/E0rA7gKxN0cM8sYqJcF7HPaY9O", // bcrypt hash for "1234"
}

// ==================================================
// Handlers
// ==================================================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	// Check username
	if user.Username != dummyUser.Username {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
		return
	}

	// Generate token
	tokenString, err := CreateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create token"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the protected area"})
}

// ==================================================
// JWT Helpers
// ==================================================
func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(secret)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
}

// ==================================================
// Middleware
// ==================================================
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing or invalid Authorization header"})
			return
		}

		tokenString := authHeader[7:]
		token, err := VerifyToken(tokenString)
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or expired token"})
			return
		}

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

// ==================================================
// Main
// ==================================================
func main() {

	//load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading env")
	}

	secret = []byte(os.Getenv("SECRET"))

	if len(secret) == 0 {
		log.Fatal("SECRET environment variable is not set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.Handle("/protected", AuthMiddleware(http.HandlerFunc(ProtectedHandler))).Methods("GET")

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Could not start the server:", err)
	}
}
