package artists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/artists/pkg/api"
)

func Get(provider *api.Provider, storeName string) ([]*Info, error) {
	url := fmt.Sprintf("%s/artists?store=%s", provider.URL, storeName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	artists := []*Info{}
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}
