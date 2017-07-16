package main

import (
	"flag"
	"lib/scaner"
	"lib/analysis"
	"lib/writer"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	projectName := flag.Arg(1)
	scanedFile := scaner.Scan(root)
	KeywordMap := analysis.Start(scanedFile)
	writer.WriteToJson(projectName, KeywordMap)

}
