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
			fmt.Println(api.RescanFile(fileIDs[0]))
		} else {
			fmt.Println(api.RescanFiles(fileIDs))
		}
	} else {
		log.Fatal("Error: args count not valid")
	}
}
