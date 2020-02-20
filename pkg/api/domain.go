package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) DomainDetails(domain string) (string, error) {
	url := fmt.Sprintf("%s/domain/%s", api.URL, domain)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("apikey", api.Token)
	if err != nil {
		return "", err
	}
	return fmtResponse(api.Client.Do(req))
}

func (api *API) DomainsDetails(address []string) (string, error) {
	url := fmt.Sprintf("%s/url", api.URL)
	payload := struct {
		FQDN []string `json:"fqdn"`
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
