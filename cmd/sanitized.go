package cmd

import (
	"github.com/OPSWAT/mdcloud-go/lookup"
	"github.com/spf13/cobra"
)

// sanitizedCmd represents the lookup command
var sanitizedCmd = &cobra.Command{
	Use:   "sanitized",
	Short: "Sanitized result by file_id",
	Long:  "Sanitized result by file_id",
	Run: func(cmd *cobra.Command, args []string) {
		lookup.SanitizedByFileID(API, args)
	},
}

func init() {
	RootCmd.AddCommand(sanitizedCmd)
}
