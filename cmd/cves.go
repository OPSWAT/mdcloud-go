package cmd

import (
	"github.com/OPSWAT/mdcloud-go/cve"
	"github.com/spf13/cobra"
)

// cvesCmd represents the cve command
var cvesCmd = &cobra.Command{
	Use:   "cves",
	Short: "Available CVE list",
	Long:  "Retrieve list of all CVEs",
	Run: func(cmd *cobra.Command, args []string) {
		cve.List(API)
	},
}

func init() {
	RootCmd.AddCommand(cvesCmd)
}
