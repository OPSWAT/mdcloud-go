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
				if res, err := api.GetHashDownloadLink(args[0]); err == nil {
					fmt.Println(res)
				} else {
					log.Fatalln(err)
				}
			} else {
				if res, err := api.HashDetails(args[0]); err == nil {
					fmt.Println(res)
				} else {
					log.Fatalln(err)
				}
			}
		} else {
			if res, err := api.HashesDetails(args); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}

// ByIP lookup
func ByIP(api api.API, args []string) {
	if len(args) > 0 {
		if len(args) == 1 {
			if res, err := api.IPDetails(args[0]); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		} else {
			if res, err := api.IPsDetails(args); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}

// AppinfoByHash lookup
func AppinfoByHash(api api.API, args []string) {
	if len(args) == 1 {
		if res, err := api.HashAppinfo(args[0]); err == nil {
			fmt.Println(res)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}

// SanitizedByFileID lookup
func SanitizedByFileID(api api.API, args []string) {
	if len(args) == 1 {
		if res, err := api.GetSanitizedLink(args[0]); err == nil {
			fmt.Println(res)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}
