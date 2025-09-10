package main

import (
	"html/template"
	"net/http"

	"github.com/zro404/ucpong/game"
	"golang.org/x/net/websocket"
)

var templ = template.Must(template.ParseGlob("templates/**/*.html"))

var rd = game.NewRoomDirectory()

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

func handleJoinForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templ.ExecuteTemplate(w, "toast", map[string]string{
			"message": "Method not allowed",
		})
		return
	}

	gameCode := r.FormValue("code")
	if gameCode == "" {
		templ.ExecuteTemplate(w, "toast", map[string]string{
			"message": "Game code is required",
		})
		return
	}

	_, ok := (*rd)[gameCode]
	if !ok {
		templ.ExecuteTemplate(w, "toast", map[string]string{
			"message": "Game not found",
		})
		return
	}

	w.Header().Set("HX-Redirect", "/game/"+gameCode)
}

func main() {

	fs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	http.HandleFunc("/", serveHome)
	http.Handle("/new", http.RedirectHandler("/game/"+rd.NewRoom(), http.StatusTemporaryRedirect))
	http.HandleFunc("/game/", serveGamePage)
	http.HandleFunc("/join", handleJoinForm)

	http.Handle("/ws", websocket.Handler(rd.HandleNewPlayer))

	registerPartialRoutes()
	http.ListenAndServe(":8000", nil)
}
