package cmd

import (
	"github.com/OPSWAT/mdcloud-go/filescan"
	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/spf13/cobra"
)

var options *filescan.ScanOptions

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan file or path",
	Long:  "Scan file or path, all folder etc.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.VerifyArgsOrRun(args, 0, func() {
			options.Path = args
			if options.Sanitization == true {
				options.Headers = append(options.Headers, "x-rule=sanitize_docs")
			}
			filescan.Scan(API, *options)
		}, func() { cmd.Help() })
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
	options = &filescan.ScanOptions{}
	scanCmd.PersistentFlags().BoolVarP(&options.Watcher, "watch", "w", false, "watches files under a path for changes & sends them to scan")
	scanCmd.PersistentFlags().BoolVarP(&options.Sanitization, "sanitize", "s", false, "enable sanitization header")
	scanCmd.PersistentFlags().StringArrayVarP(&options.Headers, "request-headers", "r", nil, "comma separated additional headers")
	scanCmd.PersistentFlags().BoolVarP(&options.LookupFile, "lookup", "l", false, "lookup sha1 before scanning one file")
	scanCmd.PersistentFlags().BoolVarP(&options.Poll, "poll", "p", true, "poll for result")
}
