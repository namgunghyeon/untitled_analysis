package main

import (
	"flag"
	"lib/scaner"
	"lib/analysis"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	scanedFile := scaner.Scan(root)
	analysis.Start(scanedFile)
}
