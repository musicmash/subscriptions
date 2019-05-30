package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func getArtists(w http.ResponseWriter, r *http.Request) {
	stores, provided := r.URL.Query()["store"]
	if !provided {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(stores[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := db.DbMgr.GetArtistsForStore(stores[0])
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
