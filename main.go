package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/**/*.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveGamePage(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "game.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	http.HandleFunc("/", serveHome)
	http.Handle("/new", http.RedirectHandler("/game", http.StatusTemporaryRedirect))
	http.HandleFunc("/game", serveGamePage)

	registerPartialRoutes()
	http.ListenAndServe(":8000", nil)
}
