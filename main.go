package main

import (
	"html/template"
	"net/http"
)

var templ *template.Template

func serveHome(w http.ResponseWriter, r *http.Request) {
	templ = template.Must(template.ParseGlob("templates/layouts/*.html"))
	templ = template.Must(templ.ParseGlob("templates/partials/*.html"))
	templ = template.Must(templ.ParseGlob("templates/pages/*.html"))

	err := templ.ExecuteTemplate(w, "base.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8000", nil)
}
