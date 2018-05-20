package cmd

import (
	"github.com/OPSWAT/mdcloud-go/cve"
	"github.com/spf13/cobra"
)

var property string

// cveCmd represents the cve command
var cveCmd = &cobra.Command{
	Use:   "cve",
	Short: "CVE lookup",
	Long:  "Retrieve CVEs by name",
	Run: func(cmd *cobra.Command, args []string) {
		cve.Lookup(API, args[0], property)
	},
}

func init() {
	RootCmd.AddCommand(cveCmd)
	cveCmd.PersistentFlags().StringVarP(&property, "type", "t", "", "list hashes, products, vendors for CVE")
}
