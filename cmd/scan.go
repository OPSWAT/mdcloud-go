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
		if sanitization {
			requestHeaders = append(requestHeaders, "x-rule=sanitize_docs")
		}
		filescan.Scan(API, args, watcher, requestHeaders)
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().BoolP("watch", "w", watcher, "watches files under a path for changes & sends them to scan")
	scanCmd.PersistentFlags().StringArrayVarP(&requestHeaders, "request-headers", "r", nil, "comma separated additional headers")
	scanCmd.PersistentFlags().BoolP("sanitize", "s", sanitization, "enable sanitization header")
}
