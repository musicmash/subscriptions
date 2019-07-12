package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	if userName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := db.DbMgr.GetSimpleUserSubscriptions(userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	bytes, err := json.Marshal(&artists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	if userName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists := []int64{}
	if err := json.NewDecoder(r.Body).Decode(&artists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(artists) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := db.DbMgr.SubscribeUser(userName, artists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	if userName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists := []int64{}
	if err := json.NewDecoder(r.Body).Decode(&artists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(artists) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := db.DbMgr.UnSubscribeUser(userName, artists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
