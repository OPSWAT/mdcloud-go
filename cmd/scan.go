package cmd

import (
	"github.com/OPSWAT/mdcloud-go/filescan"

	"github.com/spf13/cobra"
)

var path string
var watcher bool

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan file or path",
	Long:  "Scan file or path, all folder etc.",
	Run: func(cmd *cobra.Command, args []string) {
		filescan.Scan(API, path, watcher)
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path")
	scanCmd.PersistentFlags().BoolP("watch", "w", watcher, "watcher")
}
