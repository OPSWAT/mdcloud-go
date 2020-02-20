package cmd

import (
	"github.com/OPSWAT/mdcloud-go/pkg/feed"
	"github.com/spf13/cobra"
)

var page int
var engine string
var fmtType string

// feedCmd represents the feed command
var feedCmd = &cobra.Command{
	Use:   "feed [type]",
	Short: "Feed of hashes, infected or false-positives",
	Long:  "Feed of hashes, infected or false-positives. The endpoints part of malware sharing program are designed to expose the latest malware identified by OPSWAT, both infected files and possible false positives.",
	Run: func(cmd *cobra.Command, args []string) {
		feed.Lookup(API, args, page, engine, fmtType)
	},
}

func init() {
	RootCmd.AddCommand(feedCmd)
	feedCmd.PersistentFlags().IntVarP(&page, "page", "p", 0, "get specific page")
	feedCmd.PersistentFlags().StringVarP(&engine, "engine", "e", "", "get false positives for specific engine")
	feedCmd.PersistentFlags().StringVarP(&fmtType, "type", "t", "json", "set format type for infected: json, csv or bro")
}
