package api

import (
	"net/http/httptest"

	"github.com/musicmash/subscriptions/internal/db"
	apilib "github.com/musicmash/subscriptions/pkg/api"
)

var (
	server *httptest.Server
	client *apilib.Provider
)

func setup() {
	db.DbMgr = db.NewFakeDatabaseMgr()
	server = httptest.NewServer(getMux())
	client = apilib.NewProvider(server.URL, 1)
}

func teardown() {
	_ = db.DbMgr.Close()
	server.Close()
}
