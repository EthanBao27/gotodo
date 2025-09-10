/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package main

import (
	"fmt"
	"os"

	"github.com/ethanbao27/gotodo/cmd"
)

var (
	version = "dev"
	_      = "none" // commit variable is injected by goreleaser
	_      = "unknown" // date variable is injected by goreleaser
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("gotodo %s\n", version)
		os.Exit(0)
	}
	
	// on first running of gotodo, init the auto-complection command
	cmd.Execute()
}
