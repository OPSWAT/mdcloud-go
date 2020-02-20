package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) UrlDetails(Url string) (string, error) {
	url := fmt.Sprintf("%s/url/%s", api.URL, Url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("apikey", api.Token)
	if err != nil {
		return "", err
	}
	return fmtResponse(api.Client.Do(req))
}

func (api *API) UrlsDetails(address []string) (string, error) {
	url := fmt.Sprintf("%s/url", api.URL)
	payload := struct {
		URL []string `json:"url"`
	}{
		address,
	}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	req.Header.Add("apikey", api.Token)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}
