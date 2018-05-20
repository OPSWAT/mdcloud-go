package cmd

import (
	"github.com/OPSWAT/mdcloud-go/filescan"

	"github.com/spf13/cobra"
)

var path string
var watcher bool
var requestHeaders []string
var sanitization bool

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan file or path",
	Long:  "Scan file or path, all folder etc.",
	Run: func(cmd *cobra.Command, args []string) {
		if sanitization == true {
			requestHeaders = append(requestHeaders, "x-rule=sanitize_docs")
		}
		filescan.Scan(API, args, watcher, requestHeaders)
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().BoolVarP(&watcher, "watch", "w", false, "watches files under a path for changes & sends them to scan")
	scanCmd.PersistentFlags().BoolVarP(&sanitization, "sanitize", "s", false, "enable sanitization header")
	scanCmd.PersistentFlags().StringArrayVarP(&requestHeaders, "request-headers", "r", nil, "comma separated additional headers")
}
