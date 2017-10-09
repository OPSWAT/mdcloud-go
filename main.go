package main

import (
	"github.com/OPSWAT/mdcloud-go/aws"
	"github.com/OPSWAT/mdcloud-go/cmd"
)

// VERSION build var
var VERSION string

func main() {
	aws.LoadProfile()
	cmd.Execute(VERSION)
}
