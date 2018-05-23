package rescan

import (
	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/sirupsen/logrus"
)

// ByFileIDs sends to rescan using file_ids
func ByFileIDs(api api.API, fileIDs []string) {
	if len(fileIDs) > 0 {
		if len(fileIDs) == 1 {
			if res, err := api.RescanFile(fileIDs[0]); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		} else {
			if res, err := api.RescanFiles(fileIDs); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
			}
		}
	} else {
		logrus.Fatalln("args count not valid")
	}
}
