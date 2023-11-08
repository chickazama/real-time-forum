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
	serveStaticFiles(mux)
	setupAPIHandlers(mux)
	setupUIHandlers(mux)
	err := http.ListenAndServeTLS(addr, "server.crt", "server.key", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

func setupAPIHandlers(mux *http.ServeMux) {
	mux.Handle("/websocket", ws.NewWebSocketHandler(repository))
	mux.Handle("/api/user", api.NewUserHandler(repository))
	mux.Handle("/api/users", api.NewUsersHandler(repository))
	mux.Handle("/api/messages", api.NewChatHandler(repository))
	mux.Handle("/api/posts", api.NewPostsHandler(repository))
	mux.Handle("/api/comments", api.NewCommentsHandler(repository))
}

func setupUIHandlers(mux *http.ServeMux) {
	mux.Handle("/signup", ui.NewSignupHandler(repository))
	mux.Handle("/login", ui.NewLoginHandler(repository))
	mux.Handle("/logout", ui.NewLogoutHandler(repository))
	mux.HandleFunc("/", ui.Index)
}
