package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/log"
)

func doSearch(w http.ResponseWriter, r *http.Request) {
	artistNames, provided := r.URL.Query()["artist_name"]
	if !provided {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(artistNames[0]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := db.DbMgr.SearchArtists(artistNames[0])
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
