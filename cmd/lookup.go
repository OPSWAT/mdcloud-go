package cmd

import (
	"strings"

	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/spf13/cobra"
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "lookup file",
	Long:  "lookup file by md5, sha1, sha267",
	Run: func(cmd *cobra.Command, args []string) {
		var ips []string
		var hashes []string
		for _, arg := range args {
			if strings.Contains(arg, ".") {
				ips = append(ips, arg)
			} else {
				hashes = append(hashes, arg)
			}
		}
		if len(ips) > 0 {
			lookup.ByIP(API, ips)
		}
		if len(hashes) > 0 {
			lookup.ByHash(API, hashes)
		}
	},
}

func init() {
	RootCmd.AddCommand(lookupCmd)
}
