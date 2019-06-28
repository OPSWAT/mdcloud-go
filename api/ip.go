package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// IPLookupReq used for rescan body
type IPLookupReq struct {
	Address []string `json:"address"`
}

// IPDetails by file_id
func (api *API) IPDetails(ip string) (string, error) {
	url := fmt.Sprintf("%s/ip/%s", api.URL, ip)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", api.Authorization)
	if err != nil {
		return "", err
	}
	return fmtResponse(api.Client.Do(req))
}

// IPsDetails by file_ids
func (api *API) IPsDetails(address []string) (string, error) {
	url := fmt.Sprintf("%s/ip", api.URL)
	payload := &IPLookupReq{Address: address}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	req.Header.Add("Authorization", api.Authorization)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}
