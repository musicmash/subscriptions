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

func TestAPI_Artists_GetForStore(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureStoreExists(vars.StoreApple))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExistsInStore(1, vars.StoreApple, vars.StoreIDA))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExistsInStore(2, vars.StoreApple, vars.StoreIDB))

	// action
	url := fmt.Sprintf("%v/v1/artists?store=%s", server.URL, vars.StoreApple)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	artists := []*db.ArtistStoreInfo{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&artists))
	assert.Len(t, artists, 2)
	// Skrillex
	assert.Equal(t, int64(1), artists[0].ArtistID)
	assert.Equal(t, vars.StoreApple, artists[0].StoreName)
	assert.Equal(t, vars.StoreIDA, artists[0].StoreID)
	// Architects
	assert.Equal(t, int64(2), artists[1].ArtistID)
	assert.Equal(t, vars.StoreApple, artists[1].StoreName)
	assert.Equal(t, vars.StoreIDB, artists[1].StoreID)
}

func TestAPI_Artists_GetForStore_Empty(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureStoreExists(vars.StoreApple))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistSkrillex}))
	assert.NoError(t, db.DbMgr.EnsureArtistExistsInStore(1, vars.StoreApple, vars.StoreIDA))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: vars.ArtistArchitects}))
	assert.NoError(t, db.DbMgr.EnsureArtistExistsInStore(2, vars.StoreApple, vars.StoreIDB))

	// action
	url := fmt.Sprintf("%v/v1/artists?store=%s", server.URL, vars.StoreDeezer)
	resp, err := http.Get(url)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	artists := []*db.ArtistStoreInfo{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&artists))
	assert.Len(t, artists, 0)
}
