package parser

import (
  "fmt"
  "lib/model"
  "lib/languages"
  "lib/languages/python"
  "lib/languages/javascript"
  "path/filepath"
)

func getLanguage(extension string) string {
  switch extension {
  case ".go":
    return "go"
  case ".java":
    return "java"
  case ".py":
    return "python"
  case ".js":
    return "javascript"
  default:
    return ""
  }
}

func getLanguageParser(language string) {

}

func Parse(scan model.Scan) string {
  for _, path := range scan.Paths {
    extension := filepath.Ext(path)
    language := getLanguage(extension)
    if language != "" {
      fmt.Println(language)
    }
    var pythonTest python.Python
    languages.InTheForest(pythonTest)

    var javascriptTest javascript.Javascript
    languages.InTheForest(javascriptTest)
  }
  return "Parser"
}
