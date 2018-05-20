package api

import (
	"net/http"
)

// GetCVEs lists all CVEs
func (api *API) GetCVEs() (string, error) {
	req, _ := http.NewRequest("GET", URL+"cve", nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}

// GetCVEDetails returns CVE details or products, vendors, hashes for that cve
func (api *API) GetCVEDetails(CVE, property string) (string, error) {
	apiurl := URL + "cve/" + CVE
	switch property {
	case "products":
		apiurl += "/products"
	case "vendors":
		apiurl += "/vendors"
	case "hashes":
		apiurl += "/hashes"
	}
	req, _ := http.NewRequest("GET", apiurl, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}
