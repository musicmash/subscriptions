package subscriptions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/subscriptions/pkg/api"
)

const HeaderUserName = "user_name"

func Get(provider *api.Provider, userName string) ([]*Subscription, error) {
	url := fmt.Sprintf("%s/subscriptions", provider.URL)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("user_name", userName)

	resp, err := provider.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	subscriptions := []*Subscription{}
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

	url := fmt.Sprintf("%s/subscriptions", provider.URL)
	request, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("user_name", userName)

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

	url := fmt.Sprintf("%s/subscriptions", provider.URL)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("user_name", userName)

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
