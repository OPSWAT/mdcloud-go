package cve

import (
	"fmt"

	"github.com/OPSWAT/mdcloud-go/api"
)

// List all CVEs
func List(api api.API) {
	fmt.Println(api.GetCVEs())
}

// ByHash lookup
func Lookup(api api.API, CVE, property string) {
	fmt.Println(api.GetCVEDetails(CVE, property))
}
