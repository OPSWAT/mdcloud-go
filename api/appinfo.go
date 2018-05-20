package api

import (
	"fmt"
	"net/http"
)

// HashAppinfo by file_id
func (api *API) HashAppinfo(hash string) (string, error) {
	url := fmt.Sprintf("%s/appinfo/%s", api.URL, hash)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}
