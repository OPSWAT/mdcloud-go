package cmd

import (
	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/spf13/cobra"
)

// sanitizedCmd represents the lookup command
var sanitizedCmd = &cobra.Command{
	Use:   "sanitized [file_id]",
	Short: "Sanitized result by file_id",
	Long:  "Sanitized result by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			lookup.SanitizedByFileID(API, args)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(sanitizedCmd)
}
