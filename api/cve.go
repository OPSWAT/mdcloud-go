package api

import (
	"net/http"
)

// GetCVEs lists all CVEs
func (api *API) GetCVEs() string {
	req, _ := http.NewRequest("GET", URL+"cve", nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// GetCVEDetails returns CVE details or products, vendors, hashes for that cve
func (api *API) GetCVEDetails(CVE, property string) string {
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
	return FmtResponse(api.Client.Do(req))
}
