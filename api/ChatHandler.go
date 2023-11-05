package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	Repo repo.IRepository
}

func NewChatHandler(r repo.IRepository) *ChatHandler {
	return &ChatHandler{Repo: r}
}

func (h *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	targetID, err := strconv.Atoi(r.PostFormValue("targetID"))
	if err != nil {
		log.Println("here")
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	query := r.URL.Query()
	fmt.Println(len(query))
	for k, v := range query {
		fmt.Println(k)
		fmt.Println(v)
	}
	limit := math.MaxInt
	offset := 0
	limitVal := query.Get("limit")
	if limitVal != "" {
		var err error
		limit, err = strconv.Atoi(limitVal)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
	}
	offsetVal := query.Get("offset")
	if offsetVal != "" {
		var err error
		offset, err = strconv.Atoi(offsetVal)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
	}
	messages, err := h.Repo.GetLimitedMessagesBySenderAndTargetIDs(userID, targetID, limit, offset)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&messages)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
}
