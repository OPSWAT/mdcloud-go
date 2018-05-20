package cmd

import (
	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/spf13/cobra"
)

// appinfoCmd represents the appinfo command
var appinfoCmd = &cobra.Command{
	Use:   "appinfo [hash]",
	Short: "Appinfo for hash",
	Long:  "Appinfo for hash by sha1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			lookup.AppinfoByHash(API, args)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(appinfoCmd)
}
