package analysis

import (
  "fmt"
  "lib/model"
  "lib/languages/javascript"
)

type KeywordMap map[string][]model.Keyword

type Analyzer interface {
  ReadValues(file model.File) []model.Keyword
  ReadFunctions(file model.File) []model.Keyword
  ReadClasses(file model.File) []model.Keyword
  ReadInterfaces(file model.File) []model.Keyword
}

func Analysis(p Analyzer, file model.File, keyword KeywordMap) {
  functionKeywords := p.ReadFunctions(file)
  for _, item := range functionKeywords {
    keyword["function"] = append(keyword["function"], item)
  }

  valueKeywords := p.ReadValues(file)
  for _, item := range valueKeywords {
    keyword["value"] = append(keyword["value"], item)
  }

  classKeywords := p.ReadClasses(file)
  for _, item := range classKeywords {
    keyword["class"] = append(keyword["class"], item)
  }

  interfaceKeywords := p.ReadClasses(file)
  for _, item := range interfaceKeywords {
    keyword["interface"] = append(keyword["interface"], item)
  }
}

func getLanguage(extension string) Analyzer {
  switch extension {
  case ".go":
    var javascriptTest javascript.Javascript
    return javascriptTest
  case ".java":
    var javascriptTest javascript.Javascript
    return javascriptTest
  case ".py":
    var javascriptTest javascript.Javascript
    return javascriptTest
  case ".js":
    var javascriptTest javascript.Javascript
    return javascriptTest
  default:
    var javascriptTest javascript.Javascript
    return javascriptTest
  }
}

func Start(scan model.Scan) {
  keyword := make(KeywordMap)
  for _, file := range scan.Files {
    language := getLanguage(file.Ext)
    Analysis(language, file, keyword)
  }
  fmt.Println()
  fmt.Println("function", len(keyword["function"]))
  fmt.Println("value", len(keyword["value"]))
  fmt.Println("class", len(keyword["class"]))
  fmt.Println("interface", len(keyword["interface"]))
}
