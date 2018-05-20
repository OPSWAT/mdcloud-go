package feed

import (
	"fmt"

	"github.com/OPSWAT/mdcloud-go/api"
)

// Lookup feed by type
func Lookup(api api.API, args []string, page int, engine, fmtType string) {
	if args != nil && len(args) > 0 {
		switch args[0] {
		case "false-positives":
			fmt.Println(api.GetFalsePositivesFeed(engine, page))
		case "infected":
			fmt.Println(api.GetInfectedHashesFeed(fmtType, page))
		case "hashes":
			fmt.Println(api.GetHashesFeed(page))
		}
	} else {
		fmt.Println(api.GetHashesFeed(page))
	}
}
