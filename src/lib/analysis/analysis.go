package analysis

import (
  "lib/model"
  "lib/languages/javascript"
)



type Analyzer interface {
  ReadValues(file model.File) []model.Keyword
  ReadFunctions(file model.File) []model.Keyword
  ReadClasses(file model.File) []model.Keyword
  ReadInterfaces(file model.File) []model.Keyword
}

func Analysis(p Analyzer, file model.File, keyword model.KeywordMap) {
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

  interfaceKeywords := p.ReadInterfaces(file)
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

func Start(scan model.Scan) model.KeywordMap {
  keyword := make(model.KeywordMap)
  for _, file := range scan.Files {
    language := getLanguage(file.Ext)
    Analysis(language, file, keyword)
  }
  return keyword
}
