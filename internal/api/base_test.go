package api

import (
	"net/http/httptest"

	"github.com/musicmash/artists/internal/db"
)

var (
	server *httptest.Server
)

func setup() {
	db.DbMgr = db.NewFakeDatabaseMgr()
	server = httptest.NewServer(getMux())
}

func teardown() {
	_ = db.DbMgr.Close()
	server.Close()
}
