package analysis

import (
  "fmt"
  "lib/model"
  "lib/languages/javascript"
)

type KeywordMap map[string][]model.Keyword

type Analyzer interface {
  ReadValues(file model.File)
  ReadFunctions(file model.File) []model.Keyword
  ReadClasses(file model.File)
  ReadInterfaces(file model.File)
}

func Analysis(p Analyzer, file model.File, keyword KeywordMap) {
  keywords := p.ReadFunctions(file)
  for _, item := range keywords {
    keyword["function"] = append(keyword["function"], item)
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
}
