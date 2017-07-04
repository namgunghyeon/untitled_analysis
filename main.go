package main

import (
	"fmt"
	"flag"
	"lib/scanner"
	"lib/parser"
	"lib/model"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	scanner.Scan(root)
	parser.Parser()

	scan := model.Scann{
		Paths: []string{"a", "b", "c"},
	};
	fmt.Printf("%s", scan.Paths)
}
