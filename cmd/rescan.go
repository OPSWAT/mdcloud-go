package cmd

import (
	"github.com/OPSWAT/mdcloud-go/rescan"
	"github.com/spf13/cobra"
)

// rescanCmd represents the rescan command
var rescanCmd = &cobra.Command{
	Use:   "rescan [file_id]",
	Short: "Rescan file",
	Long:  "Rescan file by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			rescan.ByFileIDs(API, args)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(rescanCmd)
}
