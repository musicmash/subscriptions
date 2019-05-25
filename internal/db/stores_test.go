package db

import (
	"testing"

	"github.com/musicmash/artists/internal/testutil/vars"
	"github.com/stretchr/testify/assert"
)

func TestDB_StoreType(t *testing.T) {
	setup()
	defer teardown()

	// action
	assert.NoError(t, DbMgr.EnsureStoreExists(vars.StoreDeezer))

	// assert
	assert.True(t, DbMgr.IsStoreExists(vars.StoreDeezer))
}

func TestDB_StoreType_NotExists(t *testing.T) {
	setup()
	defer teardown()

	assert.False(t, DbMgr.IsStoreExists(vars.StoreDeezer))
}
