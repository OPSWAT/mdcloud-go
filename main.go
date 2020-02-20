package main

import (
	"os"

	"github.com/OPSWAT/mdcloud-go/cmd"
	logger "github.com/sirupsen/logrus"
)

// VERSION build var
var VERSION string

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.InfoLevel)
}

func main() {
	cmd.Execute(VERSION)
}
