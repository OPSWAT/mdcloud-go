package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// API struct containing main details
type API struct {
	URL           string
	Token         string
	Client        *http.Client
	Authorization string
}

// URL for API
const mainURL = "https://api.metadefender.com/v3"

// NewAPI object
func NewAPI(apikey string) API {
	return API{Token: apikey, Authorization: fmt.Sprintf("apikey %s", apikey), URL: mainURL,
		Client: &http.Client{Timeout: 2 * time.Minute}}
}

func fmtResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
