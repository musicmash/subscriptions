package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("user_name")
	if userName != "" {
		getUserSubscriptions(userName, w)
		return
	}

	rawArtists := r.URL.Query().Get("artists")
	if rawArtists != "" {
		getArtistsSubscriptions(parseArtists(rawArtists), w)
		return
	}

	// empty filters provided
	w.WriteHeader(http.StatusBadRequest)
}

func parseArtists(rawArtists string) []int64 {
	artists := []int64{}
	for _, rawArtist := range strings.Split(rawArtists, ",") {
		if rawArtist == "" {
			continue
		}

		artistID, err := strconv.ParseInt(rawArtist, 10, 64)
		if err != nil {
			continue
		}

		artists = append(artists, artistID)
	}
	return artists
}

func getUserSubscriptions(userName string, w http.ResponseWriter) {
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

func getArtistsSubscriptions(artists []int64, w http.ResponseWriter) {
	subscriptions, err := db.DbMgr.GetArtistsSubscriptions(artists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	bytes, err := json.Marshal(&subscriptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("user_name")
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
	userName := r.Header.Get("user_name")
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
