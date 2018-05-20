package cmd

import (
	"github.com/OPSWAT/mdcloud-go/feed"
	"github.com/spf13/cobra"
)

var page int
var engine string
var fmtType string

// feedCmd represents the feed command
var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "Feed of hashes, infected or false-positives",
	Long:  "The endpoints part of malware sharing program are designed to expose the latest malware identified by OPSWAT, both infected files and possible false positives.",
	Run: func(cmd *cobra.Command, args []string) {
		feed.Lookup(API, args, page, engine, fmtType)
	},
}

func init() {
	RootCmd.AddCommand(feedCmd)
	feedCmd.PersistentFlags().IntVarP(&page, "page", "p", 0, "get specific page")
	feedCmd.PersistentFlags().StringVarP(&engine, "engine", "e", "", "false positives for specific engine")
	feedCmd.PersistentFlags().StringVarP(&fmtType, "type", "t", "", "format type for infected (json, csv, bro)")
}
