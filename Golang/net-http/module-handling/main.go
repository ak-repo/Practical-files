package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handler

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("<h1>Hello</h1>"))
}

func main() {

	//router
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	//server
	log.Fatal(http.ListenAndServe(":8080", router))
}
