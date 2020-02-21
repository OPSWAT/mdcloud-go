package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/OPSWAT/mdcloud-go/pkg/api"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	logstash "github.com/bshuster-repo/logrus-logstash-hook"
	gelf "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prettyf "github.com/x-cray/logrus-prefixed-formatter"
)

// VERSION for build
var VERSION string

// API objects
var API api.API
var apikey string
var formatter string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mdcloud",
	Short: "Metadefender Cloud API wrapper",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch formatter {
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "gelf":
			hostname, _ := os.Hostname()
			logrus.SetFormatter(gelf.NewGelf(hostname))
		case "logstash":
			logrus.SetFormatter(&logstash.LogstashFormatter{})
		case "text":
			logrus.SetFormatter(&prettyf.TextFormatter{})
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(version string) {
	VERSION = version
	var apikeyErr error
	if apikey == "" {
		var ok bool
		if apikey, ok = os.LookupEnv("MDCLOUD_APIKEY"); ok && apikey != "" {
			API, apikeyErr = api.NewAPI(apikey)
		} else {
			API, apikeyErr = api.NewAPI("")
		}
	} else {
		API, apikeyErr = api.NewAPI(apikey)
	}

	if apikeyErr != nil {
		logrus.Fatalln(apikeyErr)
	}

	if err := RootCmd.Execute(); err != nil {
		logrus.Fatalln(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&apikey, "apikey", "a", "", "set apikey token (default is MDCLOUD_APIKEY env variable)")
	RootCmd.PersistentFlags().StringVarP(&formatter, "formatter", "f", "text", "set formatter type to json, text, logstash, gelf or raw")
}

// Response placeholder
type Response struct{}

// Format parses the message, returns raw json, no log
func (f *Response) Format(entry *logrus.Entry) ([]byte, error) {
	var js map[string]interface{}
	if json.Unmarshal([]byte(entry.Message), &js) != nil && utils.IsLetter(string(entry.Message[0])) {
		return nil, nil
	}
	result := fmt.Sprintf("%s\n", entry.Message)
	return []byte(result), nil
}
