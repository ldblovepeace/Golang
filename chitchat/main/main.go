package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ldblovepeace/chitchat/data"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index2)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}

		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	} else {
		fmt.Println(err)
	}
}
