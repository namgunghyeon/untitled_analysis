package parser

import (
  "fmt"
  "lib/model"
  "lib/languages/java"
)

func Parser(scan model.Scan) string {
  fmt.Printf("%s", scan)
  java.ReadValues()
  return "Parser"
}
