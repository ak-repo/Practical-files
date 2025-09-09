package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {

		coo1 := http.Cookie{
			Name:     "Cookie_one",
			Value:    "Value_one",
			HttpOnly: true,
		}
		coo2 := http.Cookie{
			Name:     "cookie_two",
			Value:    "value_two",
			HttpOnly: true,
		}

		http.SetCookie(w, &coo1)
		http.SetCookie(w, &coo2)

	})

	router.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		h := r.Header["Cookie"]

		fmt.Fprintln(w, h)

	})

	http.ListenAndServe(":8090", router)
}
