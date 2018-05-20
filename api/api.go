package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// API struct containing main details
type API struct {
	URL    string
	Token  string
	Client http.Client
}

// URL for API
const URL = "https://api.metadefender.com/v3/"

// NewAPI object
func NewAPI(apikey string) API {
	return API{Token: apikey, URL: URL, Client: http.Client{
		Timeout: 300 * time.Second,
	}}
}

func fmtResponse(resp *http.Response, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
