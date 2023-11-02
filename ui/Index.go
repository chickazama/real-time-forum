package ui

import (
	"html/template"
	"log"
	"net/http"
)

const (
	path         = "./static/index.go.html"
	templateName = "page"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "page", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
}
