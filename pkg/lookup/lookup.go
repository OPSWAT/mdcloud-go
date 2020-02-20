package lookup

import (
	"net/url"

	"github.com/OPSWAT/mdcloud-go/pkg/api"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/sirupsen/logrus"
)

// ByHash lookup
func ByHash(api api.API, args []string, download bool) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			if download {
				if res, err := api.GetHashDownloadLink(args[0]); err == nil {
					logrus.Println(res)
				} else {
					logrus.Fatalln(err)
				}
			} else {
				if res, err := api.HashDetails(args[0]); err == nil {
					logrus.Println(res)
				} else {
					logrus.Fatalln(err)
				}
			}
		} else {
			if res, err := api.HashesDetails(args); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		}
	})
}

// ByIP lookup
func ByIP(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			if res, err := api.IPDetails(args[0]); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		} else {
			if res, err := api.IPsDetails(args); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		}
	})
}

// ByDomain lookup
func ByDomain(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			if res, err := api.DomainDetails(args[0]); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		} else {
			if res, err := api.DomainsDetails(args); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		}
	})
}

// ByUrl lookup
func ByUrl(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 0, func() {
		if len(args) == 1 {
			escaped := url.PathEscape(args[0])
			if res, err := api.UrlDetails(escaped); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		} else {
			if res, err := api.UrlsDetails(args); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		}
	})
}

// AppinfoByHash lookup
func AppinfoByHash(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 1, func() {
		if res, err := api.HashAppinfo(args[0]); err == nil {
			logrus.Println(res)
		} else {
			logrus.Fatalln(err)
		}
	})
}

// SanitizedByFileID lookup
func SanitizedByFileID(api api.API, args []string) {
	utils.VerifyArgsOrRun(args, 1, func() {
		if res, err := api.GetSanitizedLink(args[0]); err == nil {
			logrus.Println(res)
		} else {
			logrus.Fatalln(err)
		}
	})
}
