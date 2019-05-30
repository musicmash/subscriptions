package search

type Artist struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Poster     string `json:"poster"`
	Popularity int    `json:"popularity"`
	Followers  uint   `json:"followers"`
}
