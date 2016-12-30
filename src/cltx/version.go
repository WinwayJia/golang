package main

import (
	"flag"
	"fmt"
	"os"
)

var VersionStr = "unknown"

func main() {
	version := flag.Bool("v", false, "Show Version")

	flag.Parse()
	if *version {
		fmt.Println(VersionStr)
		os.Exit(0)
	}
}
