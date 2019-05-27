package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/musicmash/artists/internal/db"
	"github.com/musicmash/artists/internal/testutil/vars"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Search(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistRitaOra}))

	// action
	url := fmt.Sprintf("%v/v1/search?artist_name=arch", server.URL)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	artists := []*db.Artist{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&artists))
	assert.Len(t, artists, 1)
	assert.Equal(t, vars.ArtistArchitects, artists[0].Name)
}

func TestAPI_Search_NotFound(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistRitaOra}))

	// action
	url := fmt.Sprintf("%v/v1/search?artist_name=xxx", server.URL)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	artists := []*db.Artist{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&artists))
	assert.Len(t, artists, 0)
}

func TestAPI_Search_NameNotProvided(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistRitaOra}))

	// action
	url := fmt.Sprintf("%v/v1/search", server.URL)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}


func TestAPI_Search_NameIsEmpty(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistRitaOra}))

	// action
	url := fmt.Sprintf("%v/v1/search?artist_name=", server.URL)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
