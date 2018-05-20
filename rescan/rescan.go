package rescan

import (
	"fmt"
	"log"

	"github.com/OPSWAT/mdcloud-go/api"
)

// ByFileIDs sends to rescan using file_ids
func ByFileIDs(api api.API, fileIDs []string) {
	if len(fileIDs) > 0 {
		if len(fileIDs) == 1 {
			if res, err := api.RescanFile(fileIDs[0]); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		} else {
			if res, err := api.RescanFiles(fileIDs); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}
