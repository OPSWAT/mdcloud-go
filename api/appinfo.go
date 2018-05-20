package api

import "net/http"

// HashAppinfo by file_id
func (api *API) HashAppinfo(hash string) string {
	req, _ := http.NewRequest("GET", URL+"appinfo/"+hash, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}
