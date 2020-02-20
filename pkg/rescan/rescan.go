package rescan

import (
	"github.com/OPSWAT/mdcloud-go/pkg/api"
	logger "github.com/sirupsen/logrus"
)

// ByFileIDs sends to rescan using file_ids
func ByFileIDs(api api.API, fileIDs []string) {
	if len(fileIDs) > 0 {
		if len(fileIDs) == 1 {
			if res, err := api.RescanFile(fileIDs[0]); err == nil {
				logger.Println(res)
			} else {
				logger.Fatalln(err)
			}
		} else {
			if res, err := api.RescanFiles(fileIDs); err == nil {
				logger.Println(res)
			} else {
				logger.Fatalln(err)
			}
		}
	} else {
		logger.Fatalln("args count not valid")
	}
}
