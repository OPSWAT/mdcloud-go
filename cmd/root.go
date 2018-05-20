package cmd

import (
	"log"

	"github.com/OPSWAT/mdcloud-go/api"

	"github.com/spf13/cobra"
)

// VERSION for build
var VERSION string

// API objects
var API api.API
var apikey string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mdcloud",
	Short: "Metadefender Cloud API wrapper",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(api api.API, version string) {
	VERSION = version
	API = api
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&apikey, "apikey", "a", "", "apikey token (default is MDCLOUD_APIKEY env variable)")
}
