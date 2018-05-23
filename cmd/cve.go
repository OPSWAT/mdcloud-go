package cmd

import (
	"github.com/OPSWAT/mdcloud-go/cve"
	"github.com/spf13/cobra"
)

var property string

// cveCmd represents the cve command
var cveCmd = &cobra.Command{
	Use:   "cve [CVE]",
	Short: "CVE lookup",
	Long:  "Retrieve CVEs by name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cve.Lookup(API, args[0], property)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(cveCmd)
	cveCmd.PersistentFlags().StringVarP(&property, "type", "t", "", "get list of hashes, products, vendors for CVE")
}
