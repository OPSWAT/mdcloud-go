package cmd

import (
	"github.com/OPSWAT/mdcloud-go/pkg/lookup"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/spf13/cobra"
)

// sanitizedCmd represents the lookup command
var sanitizedCmd = &cobra.Command{
	Use:   "sanitized [file_id]",
	Short: "Sanitized result by file_id",
	Long:  "Sanitized result by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		utils.VerifyArgsOrRun(args, 0, func() { lookup.SanitizedByFileID(API, args) }, func() { cmd.Help() })
	},
}

func init() {
	RootCmd.AddCommand(sanitizedCmd)
}
