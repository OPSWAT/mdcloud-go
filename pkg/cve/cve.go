package cve

import (
	"github.com/OPSWAT/mdcloud-go/pkg/api"
	logger "github.com/sirupsen/logrus"
)

// List all CVEs
func List(api api.API) {
	if res, err := api.GetCVEs(); err == nil {
		logger.Println(res)
	} else {
		logger.Fatalln(err)
	}
}

// Lookup cve details
func Lookup(api api.API, CVE, property string) {
	if res, err := api.GetCVEDetails(CVE, property); err == nil {
		logger.Println(res)
	} else {
		logger.Fatalln(err)
	}
}
