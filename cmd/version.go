package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mdcloud",
	Long:  "All software has versions. This is mdcloud's",
	Run: func(cmd *cobra.Command, args []string) {
		logger.WithField("version", VERSION).Println("mdcloud " + VERSION)
	},
}
