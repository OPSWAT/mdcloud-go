package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DomainLookup used for rescan body
type DomainLookup struct {
	FQDN []string `json:"fqdn"`
}

// Domain result
type Domain struct {
	Address       string `json:"address"`
	LookupResults struct {
		StartTime  time.Time `json:"start_time"`
		DetectedBy int       `json:"detected_by"`
		Sources    []struct {
			Provider   string    `json:"provider"`
			Assessment string    `json:"assessment"`
			DetectTime string    `json:"detect_time"`
			UpdateTime time.Time `json:"update_time"`
			Status     int       `json:"status"`
		} `json:"sources"`
	} `json:"lookup_results"`
	Error ApiError `json:"error"`
}

type DomainsLookupResult struct {
	Data  []Domain `json: data`
	Error ApiError `json:"error"`
}

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
	payload := &DomainLookup{FQDN: address}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	req.Header.Add("apikey", api.Token)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}
