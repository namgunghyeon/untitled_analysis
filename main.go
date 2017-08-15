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
	name := flag.Arg(1)
	version := flag.Arg(2)
	language := flag.Arg(3)
	scanedFile := scaner.Scan(root)
	KeywordMap := analysis.Start(scanedFile)
	writer.CockroacWriteToKeywordIndex(KeywordMap)
	writer.CockroacWriteToKeyword(name, version, KeywordMap)
	//writer.WriteToKeywordIndex(KeywordMap)
	//writer.WriteToKeyword(name, version, KeywordMap)
	writer.WriteToProejctInfo(name, version, language)

	//writer.WriteToProejct(name, version, KeywordMap)
	//writer.WriteToJson(name, version, KeywordMap)
}
