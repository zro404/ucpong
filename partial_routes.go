package main

import "net/http"

func serveJoinInput(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "join-input.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveJoinOptions(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "join-options.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerPartialRoutes() {
	http.HandleFunc("/ui/join-input", serveJoinInput)
	http.HandleFunc("/ui/join-options", serveJoinOptions)
}
