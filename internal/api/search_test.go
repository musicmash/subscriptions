package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/testutil/vars"
	"github.com/musicmash/subscriptions/pkg/api/search"
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
	artists, err := search.Do(client, vars.ArtistArchitects[0:4])

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, vars.ArtistArchitects, artists[0].Name)
}

func TestAPI_Search_NotFound(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistRitaOra}))

	// action
	artists, err := search.Do(client, vars.ArtistSkrillex)

	// assert
	assert.NoError(t, err)
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
	artists, err := search.Do(client, "")

	// assert
	assert.Error(t, err)
	assert.Nil(t, artists)
}
