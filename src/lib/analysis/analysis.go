package analysis

import (
  "lib/model"
  "lib/languages/python"
  "lib/languages/javascript"
)

type Analyzer interface {
  ReadValues()
  ReadFunctions()
  ReadClasses()
  ReadInterfaces()
}

func Analysis(p Analyzer, file model.File) {
  p.ReadValues()
  p.ReadFunctions()
  p.ReadClasses()
  p.ReadInterfaces()
}

func getLanguage(extension string) Analyzer {
  switch extension {
  case ".go":
    var pythonTest python.Python
    return pythonTest
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
  for _, file := range scan.Files {
    language := getLanguage(file.Ext)
    Analysis(language, file)
  }
}
