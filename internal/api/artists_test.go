package api

import (
	"testing"

	"github.com/musicmash/artists/internal/db"
	"github.com/musicmash/artists/internal/testutil/vars"
	"github.com/musicmash/artists/pkg/api/artists"
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
	artists, err := artists.Get(client, vars.StoreApple)

	// assert
	assert.NoError(t, err)
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
	artists, err := artists.Get(client, vars.StoreDeezer)

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 0)
}
