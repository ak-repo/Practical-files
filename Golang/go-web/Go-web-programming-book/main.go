package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

type World struct{}

func (h *World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World")
}

func main() {

	hello := Hello{}
	world := World{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/world", &world)
	http.Handle("/hello", &hello)

	server.ListenAndServe()

}
