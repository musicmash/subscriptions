package api

import (
	"testing"

	"github.com/musicmash/subscriptions/internal/db"
	"github.com/musicmash/subscriptions/internal/testutil/vars"
	"github.com/musicmash/subscriptions/pkg/api/subscriptions"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Subscriptions_Get(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDQ}}))
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserBot, ArtistID: vars.StoreIDW}}))

	// action
	artists, err := subscriptions.Get(client, vars.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, vars.UserObjque, artists[0].UserName)
	assert.Equal(t, int64(vars.StoreIDQ), artists[0].ArtistID)
}

func TestAPI_Subscriptions_Get_UserWithoutSubscriptions(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDQ}}))
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDW}}))

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
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDQ}}))
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDW}}))

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
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDQ}}))
	assert.NoError(t, db.DbMgr.SubscribeUser([]*db.Subscription{{UserName: vars.UserObjque, ArtistID: vars.StoreIDW}}))

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