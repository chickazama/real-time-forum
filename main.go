package main

import (
	"log"
	"matthewhope/real-time-forum/api"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/ui"
	"matthewhope/real-time-forum/ws"
	"net/http"
)

const (
	addr = ":8080"
)

func init() {
	dal.Init()
	ws.Setup()
}

func main() {
	// Define multiplexer
	mux := http.NewServeMux()
	setupHandlers(mux)
	// Define file-system root & serve static files
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/signup", ui.Signup)
	mux.HandleFunc("/login", ui.Login)
	mux.HandleFunc("/logout", ui.Logout)

	mux.HandleFunc("/", ui.Index)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setupHandlers(mux *http.ServeMux) {
	repo := dal.NewDummyRepository()
	mux.Handle("/websocket", ws.NewWebSocketHandler(repo))
	mux.Handle("/api/user", api.NewUserHandler(repo))
	mux.Handle("/api/users", api.NewUsersHandler(repo))
	mux.Handle("/api/messages", api.NewChatHandler(repo))
	mux.Handle("/api/posts", api.NewPostsHandler(repo))
	mux.Handle("/api/comments", api.NewCommentsHandler(repo))
}
