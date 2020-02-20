package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// HashDetails by file_id
func (api *API) HashDetails(hash string) (string, error) {
	url := fmt.Sprintf("%s/hash/%s", api.URL, hash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}

// HashesDetails by file_ids
func (api *API) HashesDetails(hashes []string) (string, error) {
	url := fmt.Sprintf("%s/hash", api.URL)
	payload := struct {
		Hashes []string `json:"hash"`
	}{
		hashes,
	}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}

// HashVulnerabilities by hash
func (api *API) HashVulnerabilities(hash string, limit, offset int) (string, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/vulnerability/%s", api.URL, hash))
	q := url.Query()
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	url.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}
