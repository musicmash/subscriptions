package api

import (
	"testing"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/testutil/vars"
	"github.com/musicmash/subscriptions/pkg/api/subscriptions"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Subscriptions_Get_User(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserBot, []int64{vars.StoreIDW}))

	// action
	artists, err := subscriptions.Get(client, vars.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, int64(vars.StoreIDQ), artists[0])
}

func TestAPI_Subscriptions_Get_Artists(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserBot, []int64{vars.StoreIDQ, vars.StoreIDW}))

	// action
	artists, err := subscriptions.GetArtistsSubscriptions(client, []int64{vars.StoreIDW})

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, int64(vars.StoreIDW), artists[0].ArtistID)
	assert.Equal(t, vars.UserBot, artists[0].UserName)
}

func TestAPI_Subscriptions_Get_UserWithoutSubscriptions(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDW}))

	// action
	artists, err := subscriptions.Get(client, vars.UserBot)

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 0)
}

func TestAPI_Subscriptions_Get_BadUserName(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDW}))

	// action
	artists, err := subscriptions.Get(client, "")

	// assert
	assert.Error(t, err)
	assert.Nil(t, artists)
}

func TestAPI_Subscriptions_Unsubscribe(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	assert.NoError(t, db.DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDW}))

	// action
	err := subscriptions.Delete(client, vars.UserObjque, []int64{vars.StoreIDQ, vars.StoreIDW})

	// assert
	assert.Nil(t, err)
	subs, err := db.DbMgr.GetUserSubscriptions(vars.UserObjque)
	assert.NoError(t, err)
	assert.Len(t, subs, 0)
}

func TestAPI_Subscriptions_Unsubscribe_EmptyUser(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := subscriptions.Delete(client, "", []int64{})

	// assert
	assert.Error(t, err)
}

func TestAPI_Subscriptions_Unsubscribe_EmptyArtists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := subscriptions.Delete(client, vars.UserObjque, []int64{})

	// assert
	assert.Error(t, err)
}

func TestAPI_Subscriptions_Subscribe(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := subscriptions.Create(client, vars.UserObjque, []int64{vars.StoreIDQ, vars.StoreIDW})

	// assert
	assert.Nil(t, err)
	subs, err := db.DbMgr.GetUserSubscriptions(vars.UserObjque)
	assert.NoError(t, err)
	assert.Len(t, subs, 2)
}
