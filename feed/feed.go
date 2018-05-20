package feed

import (
	"fmt"
	"log"

	"github.com/OPSWAT/mdcloud-go/api"
)

// Lookup feed by type
func Lookup(api api.API, args []string, page int, engine, fmtType string) {
	if args != nil && len(args) > 0 {
		switch args[0] {
		case "false-positives":
			if res, err := api.GetFalsePositivesFeed(engine, page); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		case "infected":
			if res, err := api.GetInfectedHashesFeed(fmtType, page); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		case "hashes":
			if res, err := api.GetHashesFeed(page); err == nil {
				fmt.Println(res)
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		if res, err := api.GetHashesFeed(page); err == nil {
			fmt.Println(res)
		} else {
			log.Fatalln(err)
		}
	}
}
