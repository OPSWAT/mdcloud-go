package cmd

import (
	"github.com/OPSWAT/mdcloud-go/ipscan"

	"github.com/spf13/cobra"
)

var groups []string

// sgscanCmd represents the sgscan command
var sgscanCmd = &cobra.Command{
	Use:   "sgscan",
	Short: "Scan security groups using IP Scan API",
	Long:  `Scan security groups associated with your AWS account based on your main credentials using IP Scan API`,
	Run: func(cmd *cobra.Command, args []string) {
		ipscan.Apikey = cmd.Parent().PersistentFlags().Lookup("apikey").Value.String()
		if groups != nil {
			ipscan.ScanSGs(groups)
		} else {
			ipscan.ScanSGs(nil)
		}
	},
}

func init() {
	RootCmd.AddCommand(sgscanCmd)
	sgscanCmd.PersistentFlags().StringArrayVarP(&groups, "include", "i", nil, "specific security groups to scan")
}
