package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// IPLookupReq used for rescan body
type IPLookupReq struct {
	Address []string `json:"address"`
}

// IPDetails by file_id
func (api *API) IPDetails(ip string) string {
	req, _ := http.NewRequest("GET", URL+"ip/"+ip, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}

// IPsDetails by file_ids
func (api *API) IPsDetails(address []string) string {
	payload := &IPLookupReq{Address: address}
	j, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", api.URL+"ip", bytes.NewBuffer(j))
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}
