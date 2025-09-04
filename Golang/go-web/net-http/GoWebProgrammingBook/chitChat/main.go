package main

import (
	"net/http"
	"text/template"
)

func index(writer http.ResponseWriter, request *http.Request) {

	// 1
	// files := []string{
	// 	"templates/layout.html",
	// 	"templates/navbar.html",
	// 	"templates/index.html"}

	// templates := template.Must(template.ParseFiles(files...))
	// threads, err := data.Threads()
	// if err == nil {
	// 	templates.ExecuteTemplate(writer, "layout", threads)
	// }

	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
		public_tmpl_files := []string{"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}
		private_tmpl_files := []string{"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}
		var templates *template.Template
		if err != nil {
			templates = templates.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates.ExecuteTemplate(writer, "layout", threads)
		}
	}
}




func main() {

	mux := http.NewServeMux()

	//static file
	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	//server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
