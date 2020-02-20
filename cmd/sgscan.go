package cmd

import (
	"github.com/OPSWAT/mdcloud-go/pkg/aws"
	"github.com/OPSWAT/mdcloud-go/pkg/ipscan"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var groups []string

// sgscanCmd represents the sgscan command
var sgscanCmd = &cobra.Command{
	Use:   "sgscan",
	Short: "Scan security groups using IP scan API",
	Long:  "Scan security groups associated with your AWS account based on your main credentials using IP Scan API",
	Run: func(cmd *cobra.Command, args []string) {
		if aws.Session == nil {
			aws.LoadProfile()
			if _, err := aws.Session.Config.Credentials.Get(); err != nil {
				logrus.Fatalln("Couldn't find AWS config under~/.aws/credentials")
			}
		}
		if groups != nil {
			ipscan.ScanSGs(API, groups)
		} else {
			ipscan.ScanSGs(API, nil)
		}
	},
}

func init() {
	RootCmd.AddCommand(sgscanCmd)
	sgscanCmd.PersistentFlags().StringArrayVarP(&groups, "include", "i", nil, "specific security groups to scan")
}
