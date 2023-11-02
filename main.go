package main

import (
	"log"
	"matthewhope/real-time-forum/ui"
	"net/http"
)

const (
	addr = ":8080"
)

func main() {
	// Define multiplexer
	mux := http.NewServeMux()
	// Define file-system root & serve static files
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", ui.Index)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
