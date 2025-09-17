package main

import (
	"fmt"
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

	_, ok := (*rd).Rooms[gameCode]
	if !ok {
		templ.ExecuteTemplate(w, "toast", map[string]string{
			"message": "Game not found",
		})
		return
	}

	if (*rd).Rooms[gameCode].IsFull() {
		templ.ExecuteTemplate(w, "toast", map[string]string{
			"message": "Room is full",
		})
		return
	}

	w.Header().Set("HX-Redirect", "/game/"+gameCode)
}

func handleRandomGame(w http.ResponseWriter, r *http.Request) {
	gameCode := (*rd).FindGame()
	w.Header().Set("HX-Redirect", "/game/"+gameCode)
}

func main() {

	fs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	http.HandleFunc("/", serveHome)
	http.Handle("/new", http.RedirectHandler("/game/"+rd.NewRoom(), http.StatusTemporaryRedirect))
	http.HandleFunc("/game/", serveGamePage)
	http.HandleFunc("/join", handleJoinForm)
	http.HandleFunc("/random", handleRandomGame)

	http.Handle("/ws", websocket.Handler(rd.HandleNewPlayer))

	registerPartialRoutes()

	fmt.Println("\033[1;97mLive\033[0m @ \033[1;96m http://localhost:8000 \033[0m")
	fmt.Println("press \033[1;97mCTRL + C\033[0m to exit")

	http.ListenAndServe(":8000", nil)
}
