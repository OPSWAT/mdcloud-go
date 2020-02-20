package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// IPDetails by file_id
func (api *API) IPDetails(ip string) (string, error) {
	url := fmt.Sprintf("%s/ip/%s", api.URL, ip)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("apikey", api.Token)
	if err != nil {
		return "", err
	}
	return fmtResponse(api.Client.Do(req))
}

// IPsDetails by file_ids
func (api *API) IPsDetails(address []string) (string, error) {
	url := fmt.Sprintf("%s/ip", api.URL)
	payload := struct {
		Address []string `json:"address"`
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
