package parser

import (
  "fmt"
  "lib/model"
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
  default:
    return ""
  }
}

func Parser(scan model.Scan) string {
  for _, path := range scan.Paths {
    extension := filepath.Ext(path)
    language := getLanguage(extension)
    if language != "" {
      fmt.Println(language)
    }
  }
  return "Parser"
}
