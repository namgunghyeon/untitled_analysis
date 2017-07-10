package analysis

import (
  "lib/model"
  "lib/languages/python"
  "lib/languages/javascript"
)

type Analyzer interface {
  ReadValues(file model.File)
  ReadFunctions(file model.File)
  ReadClasses(file model.File)
  ReadInterfaces(file model.File)
}

func Analysis(p Analyzer, file model.File) {
  p.ReadValues(file)
  p.ReadFunctions(file)
  p.ReadClasses(file)
  p.ReadInterfaces(file)
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
