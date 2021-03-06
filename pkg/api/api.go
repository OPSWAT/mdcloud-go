package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// API struct containing main details
type API struct {
	URL    string
	Token  string
	Client *http.Client
	Limits map[string][]string
	Type   int
}

type ApiError struct {
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

type ApikeyStatus struct {
	QosScan  string `json:"qos_scan"`
	PaidUser int    `json:"paid_user"`
}

// URL for API
const (
	mainURL          = "https://api.metadefender.com/v4"
	shared           = "a93f6ed2dec2a246b69935eefd318273"
	LimitFor         = "X-RateLimit-For"
	LimitInterval    = "X-RateLimit-Interval"
	LimitLimit       = "X-RateLimit-Limit"
	LimitRemainingIn = "X-RateLimit-Reset"
	LimitReset       = "X-RateLimit-Reset"
	LimitUsed        = "X-RateLimit-Used"
)

// NewAPI object
func NewAPI(apikey string) (API, error) {
	api := API{Token: apikey, URL: mainURL, Client: &http.Client{Timeout: 2 * time.Minute}, Limits: make(map[string][]string)}
	if apikey != "" {
		api.Token = apikey
		err := api.setType()
		if err != nil {
			return api, err
		}
	} else {
		api.Token = shared
	}
	return api, nil
}

func (api *API) setType() error {
	var err error
	var apikeyStatus ApikeyStatus
	url := fmt.Sprintf("%s/apikey/%s", api.URL, api.Token)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("apikey", api.Token)
	resp, err := api.Client.Do(req)
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if e := json.Unmarshal(body, &apikeyStatus); e != nil {
			logrus.Fatalln(e)
			return e
		}
	}
	defer resp.Body.Close()
	api.Type = apikeyStatus.PaidUser
	return err
}

func fmtResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
