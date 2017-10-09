package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var apikey string

// VERSION for build
var VERSION string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mdcloud-go",
	Short: "Metadefender Cloud API wrapper",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if apikey == "" {
			if result, ok := os.LookupEnv("MDCLOUD_APIKEY"); ok {
				apikey = result
			} else {
				println("Apikey not set, please specify token while calling or set the environment variable")
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(version string) {
	VERSION = version
	if err := RootCmd.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&apikey, "apikey", "a", "", "apikey token (default is MDCLOUD_APIKEY env variable)")
}
