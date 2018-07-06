package cmd

import (
	"strings"

	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/spf13/cobra"
)

var download bool

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup [hash]",
	Short: "Lookup or download file or IP",
	Long:  "Lookup or download file by md5, sha1, sha256 or IP",
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
		utils.VerifyArgsOrRun(args, 0, func() {
			if len(ips) > 0 {
				lookup.ByIP(API, ips)
			}
			if len(hashes) > 0 {
				lookup.ByHash(API, hashes, download)
			}
		}, func() { cmd.Help() })
	},
}

func init() {
	RootCmd.AddCommand(lookupCmd)
	lookupCmd.PersistentFlags().BoolVarP(&download, "download", "d", false, "get download url")
}
