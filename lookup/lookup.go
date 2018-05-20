package lookup

import (
	"fmt"
	"log"

	"github.com/OPSWAT/mdcloud-go/api"
)

// ByHash lookup
func ByHash(api api.API, args []string, download bool) {
	if len(args) > 0 {
		if len(args) == 1 {
			if download {
				fmt.Println(api.GetHashDownloadLink(args[0]))
			} else {
				fmt.Println(api.HashDetails(args[0]))
			}
		} else {
			fmt.Println(api.HashesDetails(args))
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}

// ByIP lookup
func ByIP(api api.API, args []string) {
	if len(args) > 0 {
		if len(args) == 1 {
			fmt.Println(api.IPDetails(args[0]))
		} else {
			fmt.Println(api.IPsDetails(args))
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}

// AppinfoByHash lookup
func AppinfoByHash(api api.API, args []string) {
	if len(args) == 1 {
		fmt.Println(api.HashAppinfo(args[0]))
	} else {
		log.Fatal("Error: args count not valid")
	}
}
