package main

import (
	"html/template"
	"net/http"

	"github.com/zro404/ucpong/game"
	"golang.org/x/net/websocket"
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
	rd := game.NewRoomDirectory()

	fs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	http.HandleFunc("/", serveHome)
	http.Handle("/new", http.RedirectHandler("/game/"+rd.NewRoom(), http.StatusTemporaryRedirect))
	http.HandleFunc("/game/", serveGamePage)

	http.Handle("/ws", websocket.Handler(rd.HandleNewPlayer))

	registerPartialRoutes()
	http.ListenAndServe(":8000", nil)
}
