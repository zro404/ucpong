package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/**/*.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "base.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8000", nil)
}
