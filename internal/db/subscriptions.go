package db

type Subscription struct {
	ID       uint64 `json:"id"        gorm:"primary_key"     sql:"AUTO_INCREMENT"`
	UserName string `json:"name"      gorm:"unique_index:idx_user_name_artist_id"`
	ArtistID int64  `json:"artist_id" gorm:"unique_index:idx_user_name_artist_id"`
}

type SubscriptionMgr interface {
	GetUserSubscriptions(userName string) ([]*Subscription, error)
	SubscribeUser(subscriptions []*Subscription) error
	UnSubscribeUser(userName string, artists []int64) error
}

func (mgr *AppDatabaseMgr) GetUserSubscriptions(userName string) ([]*Subscription, error) {
	subs := []*Subscription{}
	err := mgr.db.Where("user_name = ?", userName).Find(&subs).Error
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (mgr *AppDatabaseMgr) SubscribeUser(subscriptions []*Subscription) error {
	tx := mgr.Begin()
	for _, sub := range subscriptions {
		if err := tx.db.Create(&sub).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (mgr *AppDatabaseMgr) UnSubscribeUser(userName string, artists []int64) error {
	const query = "delete from subscriptions where user_name = ? and artist_id in (?)"
	return mgr.db.Exec(query, userName, artists).Error
}
