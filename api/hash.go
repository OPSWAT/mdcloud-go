package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// HashLookupReq used for rescan body
type HashLookupReq struct {
	Hashes []string `json:"hash"`
}

// HashDetails by file_id
func (api *API) HashDetails(hash string) string {
	req, _ := http.NewRequest("GET", URL+"hash/"+hash, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// HashesDetails by file_ids
func (api *API) HashesDetails(hashes []string) string {
	payload := &HashLookupReq{Hashes: hashes}
	j, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", api.URL+"hash", bytes.NewBuffer(j))
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("content-type", "application/json")
	return FmtResponse(api.Client.Do(req))
}

// HashVulnerabilities by file_ids
func (api *API) HashVulnerabilities(hash string, limit, offset int) string {
	url, _ := url.Parse(URL + "vulnerability/" + hash)
	q := url.Query()
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}
