package api

import (
	"fmt"
	"net/http"
)

// GetCVEs lists all CVEs
func (api *API) GetCVEs() (string, error) {
	url := fmt.Sprintf("%s/cve", api.URL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}

// GetCVEDetails returns CVE details or products, vendors, hashes for that cve
func (api *API) GetCVEDetails(CVE, property string) (string, error) {
	apiurl := fmt.Sprintf("%s/cve/%s", api.URL, CVE)
	switch property {
	case "products":
		apiurl += "/products"
	case "vendors":
		apiurl += "/vendors"
	case "hashes":
		apiurl += "/hashes"
	}
	req, err := http.NewRequest(http.MethodGet, apiurl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}
