package api

import (
	"fmt"
	"net/http"
)

// HashAppinfo by file_id
func (api *API) HashAppinfo(hash string) (string, error) {
	url := fmt.Sprintf("%s/appinfo/%s", api.URL, hash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}
