package api

import (
	"io/ioutil"
	"log"
	"net"
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
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}}
}

// FmtResponse to string from resp
func FmtResponse(resp *http.Response, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
