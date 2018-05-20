package cve

import (
	"fmt"
	"log"

	"github.com/OPSWAT/mdcloud-go/api"
)

// List all CVEs
func List(api api.API) {
	if res, err := api.GetCVEs(); err == nil {
		fmt.Println(res)
	} else {
		log.Fatalln(err)
	}
}

// Lookup cve details
func Lookup(api api.API, CVE, property string) {
	if res, err := api.GetCVEDetails(CVE, property); err == nil {
		fmt.Println(res)
	} else {
		log.Fatalln(err)
	}
}
