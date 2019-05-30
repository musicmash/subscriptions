package search

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/artists/pkg/api"
)

func Do(provider *api.Provider, artistName string) ([]*Artist, error) {
	url := fmt.Sprintf("%s/search?artist_name=%s", provider.URL, artistName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	artists := []*Artist{}
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}
