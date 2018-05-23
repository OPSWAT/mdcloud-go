package api

import (
	"fmt"
	"net/http"
)

// GetCVEs lists all CVEs
func (api *API) GetCVEs() (string, error) {
	url := fmt.Sprintf("%s/cve", api.URL)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", api.Authorization)
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
	req, _ := http.NewRequest(http.MethodGet, apiurl, nil)
	req.Header.Add("Authorization", api.Authorization)
	return fmtResponse(api.Client.Do(req))
}
