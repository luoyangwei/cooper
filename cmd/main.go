package main

import (
	"flag"
	"fmt"
	"luoyangwei/cooper"
)

var (
	cooperCmd cooper.Cooper
)

func init() {
	flag.Var(&cooperCmd, "file", "file to process")
}

func main() {
	flag.Parse()

	fmt.Println("files: ", cooperCmd.Files)
	cooperCmd.Execute()
}
