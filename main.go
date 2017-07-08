package main

import (
	"flag"
	"lib/scaner"
	"lib/parser"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	scanned := scaner.Scan(root)
	parser.Parse(scanned)
}
