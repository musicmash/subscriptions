package subscriptions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/subscriptions/pkg/api"
)

func Get(provider *api.Provider, userName string) ([]int64, error) {
	url := fmt.Sprintf("%s/subscriptions?user_name=%s", provider.URL, userName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	subscriptions := []int64{}
	if err := json.NewDecoder(resp.Body).Decode(&subscriptions); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func Delete(provider *api.Provider, userName string, artists []int64) error {
	body, err := json.Marshal(&artists)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/subscriptions?user_name=%s", provider.URL, userName)
	request, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := provider.Client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("got %d status code", resp.StatusCode)
	}
	return nil
}

func Create(provider *api.Provider, userName string, artists []int64) error {
	body, err := json.Marshal(&artists)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/subscriptions?user_name=%s", provider.URL, userName)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := provider.Client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("got %d status code", resp.StatusCode)
	}
	return nil
}
