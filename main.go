package main

import (
	"fmt"
	"os"

	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/OPSWAT/mdcloud-go/aws"
	"github.com/OPSWAT/mdcloud-go/cmd"
)

// VERSION build var
var VERSION string

// API main
var API api.API

func main() {
	if result, ok := os.LookupEnv("MDCLOUD_APIKEY"); ok {
		API = api.NewAPI(result)
	} else {
		fmt.Println("Apikey not set, please specify token while calling or set the environment variable")
		os.Exit(1)
	}
	aws.LoadProfile()
	cmd.Execute(API, VERSION)
}
