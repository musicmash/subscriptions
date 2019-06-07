package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	usersNames, provided := r.URL.Query()["user_name"]
	if !provided {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(usersNames[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := db.DbMgr.GetUserSubscriptions(usersNames[0])
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
	usersNames, provided := r.URL.Query()["user_name"]
	if !provided {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(usersNames[0]) == 0 {
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

	subs := make([]*db.Subscription, len(artists))
	for i, artistID := range artists {
		subs[i] = &db.Subscription{UserName: usersNames[0], ArtistID: artistID}
	}

	if err := db.DbMgr.SubscribeUser(subs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	usersNames, provided := r.URL.Query()["user_name"]
	if !provided {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(usersNames[0]) == 0 {
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

	if err := db.DbMgr.UnSubscribeUser(usersNames[0], artists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
