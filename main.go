package main

import (
	"os"

	"github.com/OPSWAT/mdcloud-go/cmd"
	"github.com/sirupsen/logrus"
)

// VERSION build var
var VERSION string

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	cmd.Execute(VERSION)
}
