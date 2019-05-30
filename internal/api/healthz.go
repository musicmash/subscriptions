package api

import (
	"net/http"

	"github.com/musicmash/subscriptions/internal/db"
)

func healthz(w http.ResponseWriter, _ *http.Request) {
	if err := db.DbMgr.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
