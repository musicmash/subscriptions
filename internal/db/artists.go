package db

import (
	"fmt"
	"html"
)

type Artist struct {
	ID     int64  `json:"id"    sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name   string `json:"name"`
	Poster string `json:"poster"`
}

type ArtistStoreInfo struct {
	ID        int64  `json:"-"    gorm:"primary_key"   sql:"AUTO_INCREMENT"`
	ArtistID  int64  `json:"-"`
	StoreName string `json:"name" gorm:"unique_index:idx_art_store_name_id"`
	StoreID   string `json:"id"   gorm:"unique_index:idx_art_store_name_id"`
}

type ArtistMgr interface {
	EnsureArtistExists(artist *Artist) error
	GetAllArtists() ([]*Artist, error)
	SearchArtists(name string) ([]*Artist, error)
}

func (mgr *AppDatabaseMgr) GetAllArtists() ([]*Artist, error) {
	artists := []*Artist{}
	if err := mgr.db.Find(&artists).Error; err != nil {
		return nil, err
	}
	return artists, nil
}

func (mgr *AppDatabaseMgr) SearchArtists(name string) ([]*Artist, error) {
	artists := []*Artist{}
	name = fmt.Sprintf("%%%s%%", html.EscapeString(name))
	if err := mgr.db.Where("name LIKE ?", name).Order("name").Find(&artists).Error; err != nil {
		return nil, err
	}
	return artists, nil
}

func (mgr *AppDatabaseMgr) EnsureArtistExists(artist *Artist) error {
	return mgr.db.Create(artist).Error
}

type ArtistStoreInfoMgr interface {
	GetArtistsForStore(name string) ([]*ArtistStoreInfo, error)
	GetArtistFromStore(id int64, store string) ([]*ArtistStoreInfo, error)
	IsArtistExistsInStore(storeName, storeID string) bool
	EnsureArtistExistsInStore(artistID int64, storeName, storeID string) error
}

func (mgr *AppDatabaseMgr) GetArtistFromStore(id int64, store string) ([]*ArtistStoreInfo, error) {
	artists := []*ArtistStoreInfo{}
	if err := mgr.db.Where("artist_id = ? and store_name = ?", id, store).Find(&artists).Error; err != nil {
		return nil, err
	}
	return artists, nil
}

func (mgr *AppDatabaseMgr) GetArtistsForStore(name string) ([]*ArtistStoreInfo, error) {
	artists := []*ArtistStoreInfo{}
	if err := mgr.db.Where("store_name = ?", name).Find(&artists).Error; err != nil {
		return nil, err
	}
	return artists, nil
}

func (mgr *AppDatabaseMgr) IsArtistExistsInStore(storeName, storeID string) bool {
	info := ArtistStoreInfo{}
	err := mgr.db.Where("store_name = ? and store_id = ?", storeName, storeID).First(&info).Error
	return err == nil
}

func (mgr *AppDatabaseMgr) EnsureArtistExistsInStore(artistID int64, storeName, storeID string) error {
	if !mgr.IsArtistExistsInStore(storeName, storeID) {
		return mgr.db.Create(&ArtistStoreInfo{ArtistID: artistID, StoreName: storeName, StoreID: storeID}).Error
	}
	return nil
}
