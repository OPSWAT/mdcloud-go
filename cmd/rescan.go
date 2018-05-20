package cmd

import (
	"github.com/OPSWAT/mdcloud-go/rescan"
	"github.com/spf13/cobra"
)

// rescanCmd represents the rescan command
var rescanCmd = &cobra.Command{
	Use:   "rescan",
	Short: "Rescan file",
	Long:  "Rescan file by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		rescan.ByFileIDs(API, args)
	},
}

func init() {
	RootCmd.AddCommand(rescanCmd)
}
