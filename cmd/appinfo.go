package cmd

import (
	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/spf13/cobra"
)

// appinfoCmd represents the appinfo command
var appinfoCmd = &cobra.Command{
	Use:   "appinfo",
	Short: "Appinfo for hash",
	Long:  "Appinfo for hash by sha1",
	Run: func(cmd *cobra.Command, args []string) {
		lookup.AppinfoByHash(API, args)
	},
}

func init() {
	RootCmd.AddCommand(appinfoCmd)
}
