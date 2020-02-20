package cmd

import (
	"github.com/OPSWAT/mdcloud-go/pkg/rescan"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/spf13/cobra"
)

// rescanCmd represents the rescan command
var rescanCmd = &cobra.Command{
	Use:   "rescan [file_id]",
	Short: "Rescan file",
	Long:  "Rescan file by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		utils.VerifyArgsOrRun(args, 0, func() { rescan.ByFileIDs(API, args) }, func() { cmd.Help() })
	},
}

func init() {
	RootCmd.AddCommand(rescanCmd)
}
