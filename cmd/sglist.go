package cmd

import (
	"github.com/OPSWAT/mdcloud-go/ipscan"

	"github.com/spf13/cobra"
)

var grps []string

// sglistCmd represents the sglist command
var sglistCmd = &cobra.Command{
	Use:   "sglist",
	Short: "List security groups IPs",
	Long:  "List IPs from security groups associated with your AWS account based on your main credentials",
	Run: func(cmd *cobra.Command, args []string) {
		ipscan.Apikey = cmd.Parent().PersistentFlags().Lookup("apikey").Value.String()
		if grps != nil {
			ipscan.ListIPs(grps)
		} else {
			ipscan.ListIPs(nil)
		}
	},
}

func init() {
	RootCmd.AddCommand(sglistCmd)
	sglistCmd.PersistentFlags().StringArrayVarP(&grps, "include", "i", nil, "specific security groups to list")
}
