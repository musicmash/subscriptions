package db

import (
	"testing"

	"github.com/musicmash/subscriptions/internal/testutil/vars"
	"github.com/stretchr/testify/assert"
)

func TestDB_Subscriptions_SubscribeAndGet(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))

	// action
	subs, err := DbMgr.GetUserSubscriptions(vars.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.Len(t, subs, 1)
	assert.Equal(t, vars.UserObjque, subs[0].UserName)
	assert.Equal(t, int64(vars.StoreIDQ), subs[0].ArtistID)
}

func TestDB_Subscriptions_Get_ForAnotherUser(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))

	// action
	subs, err := DbMgr.GetUserSubscriptions(vars.UserBot)

	// assert
	assert.NoError(t, err)
	assert.Len(t, subs, 0)
}

func TestDB_Subscriptions_SubscribeAndGetSimple(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))

	// action
	subs, err := DbMgr.GetSimpleUserSubscriptions(vars.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.Len(t, subs, 1)
	assert.Equal(t, int64(vars.StoreIDQ), subs[0])
}

func TestDB_Subscriptions_Get_ForAnotherUserSimple(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))

	// action
	subs, err := DbMgr.GetSimpleUserSubscriptions(vars.UserBot)

	// assert
	assert.NoError(t, err)
	assert.Len(t, subs, 0)
}

func TestDB_Subscriptions_UnSubscribe(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.SubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ}))
	subs, err := DbMgr.GetUserSubscriptions(vars.UserObjque)
	assert.NoError(t, err)
	assert.Equal(t, vars.UserObjque, subs[0].UserName)
	assert.Equal(t, int64(vars.StoreIDQ), subs[0].ArtistID)

	// action
	err = DbMgr.UnSubscribeUser(vars.UserObjque, []int64{vars.StoreIDQ})

	// assert
	assert.NoError(t, err)
	subs, err = DbMgr.GetUserSubscriptions(vars.UserObjque)
	assert.NoError(t, err)
	assert.Len(t, subs, 0)
}
