package main

import (
	"flag"
	"lib/scanner"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	scanner.Scan(root)
}
