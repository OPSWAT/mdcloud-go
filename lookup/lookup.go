package lookup

import (
	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/OPSWAT/mdcloud-go/utils"
	logger "github.com/sirupsen/logrus"
)

// ByHash lookup
func ByHash(api api.API, args []string, download bool) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			if download {
				if res, err := api.GetHashDownloadLink(args[0]); err == nil {
					logger.Println(res)
				} else {
					logger.Fatalln(err)
				}
			} else {
				if res, err := api.HashDetails(args[0]); err == nil {
					logger.Println(res)
				} else {
					logger.Fatalln(err)
				}
			}
		} else {
			if res, err := api.HashesDetails(args); err == nil {
				logger.Println(res)
			} else {
				logger.Fatalln(err)
			}
		}
	})
}

// ByIP lookup
func ByIP(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			if res, err := api.IPDetails(args[0]); err == nil {
				logger.Println(res)
			} else {
				logger.Fatalln(err)
			}
		} else {
			if res, err := api.IPsDetails(args); err == nil {
				logger.Println(res)
			} else {
				logger.Fatalln(err)
			}
		}
	})
}

// AppinfoByHash lookup
func AppinfoByHash(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 1, func() {
		if res, err := api.HashAppinfo(args[0]); err == nil {
			logger.Println(res)
		} else {
			logger.Fatalln(err)
		}
	})
}

// SanitizedByFileID lookup
func SanitizedByFileID(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 1, func() {
		if res, err := api.GetSanitizedLink(args[0]); err == nil {
			logger.Println(res)
		} else {
			logger.Fatalln(err)
		}
	})
}
