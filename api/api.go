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
	URL           string
	Token         string
	Client        *http.Client
	Authorization string
	Limits        map[string][]string
	Type          int
}

type ApikeyStatus struct {
	Success bool `json:"success"`
	Data    struct {
		QosScan  string `json:"qos_scan"`
		PaidUser int    `json:"paid_user"`
	} `json:"data"`
}

// URL for API
const (
	mainURL          = "https://api.metadefender.com/v3"
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
	api := API{Token: apikey, Authorization: fmt.Sprintf("apikey %s", apikey), URL: mainURL,
		Client: &http.Client{Timeout: 2 * time.Minute}, Limits: make(map[string][]string)}
	if apikey != "" {
		api.Token = apikey
		api.Authorization = fmt.Sprintf("apikey %s", apikey)
		err := api.getAPIType()
		if err != nil {
			return api, err
		}
	} else {
		api.Token = shared
		api.Authorization = fmt.Sprintf("apikey %s", shared)
	}
	return api, nil
}

func (api *API) getAPIType() error {
	var err error
	var apikeyStatus ApikeyStatus
	url := fmt.Sprintf("%s/apikey/%s", api.URL, api.Token)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", api.Authorization)
	resp, err := api.Client.Do(req)
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if e := json.Unmarshal(body, &apikeyStatus); e != nil {
			logrus.Fatalln(e)
			return e
		}
	}
	defer resp.Body.Close()
	api.Type = apikeyStatus.Data.PaidUser
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
