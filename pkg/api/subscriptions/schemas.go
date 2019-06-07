package subscriptions

type Subscription struct {
	ID       uint64 `json:"id"`
	UserName string `json:"name"`
	ArtistID int64  `json:"artist_id"`
}
