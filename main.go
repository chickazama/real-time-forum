package main

import (
	"log"
	"matthewhope/real-time-forum/api"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/repo"
	"matthewhope/real-time-forum/ui"
	"matthewhope/real-time-forum/ws"
	"net/http"
)

var (
	repository repo.IRepository
)

const (
	addr = ":8080"
)

func init() {
	dal.Init()
	repository = repo.NewSQLiteRepository()
	ws.Setup(repository)
}

func main() {
	// Define multiplexer
	mux := http.NewServeMux()
	// Define file-system root & serve static files
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	setupHandlers(mux)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/login", ui.Login)
	mux.HandleFunc("/logout", ui.Logout)
	mux.HandleFunc("/", ui.Index)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setupHandlers(mux *http.ServeMux) {
	mux.Handle("/websocket", ws.NewWebSocketHandler(repository))
	mux.Handle("/api/user", api.NewUserHandler(repository))
	mux.Handle("/api/users", api.NewUsersHandler(repository))
	mux.Handle("/api/messages", api.NewChatHandler(repository))
	mux.Handle("/api/posts", api.NewPostsHandler(repository))
	mux.Handle("/api/comments", api.NewCommentsHandler(repository))
	mux.Handle("/signup", ui.NewSignupHandler(repository))
}
