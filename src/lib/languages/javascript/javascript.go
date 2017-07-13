package javascript

import (
  "fmt"
  "lib/parser"
  "lib/model"
  "strings"
  "regexp"
)

type Javascript struct {

}

func findFuncStartIndex(line string) int{
  return strings.Index(line, "function")
}

func findFuncEndIndex(line string, startIndex int) int{
  for j := startIndex; j < len(line); j++ {
    if string(line[j]) == "(" {
      return j
    }
  }
  return startIndex
}

func getFuncLine(line string, startIndex int, endIndex int) string {
  funcCode := string(line[startIndex:endIndex])
  funcCodes := strings.Split(funcCode, " ")
  if len(funcCodes) > 1 && string(funcCodes[1]) != "" {
    return funcCode
  }
  return ""
}

func isFuncLine(funcCode string) bool{
  return len(strings.Split(funcCode, " ")) == 2
}
func getFuncName(funcCode string) string {
  return string(strings.Split(funcCode, " ")[1])
}

func (p Javascript) ReadValues(file model.File) {
  code := parser.ReadSourceCode(file.Path)
  codes := strings.Split(code, "\n")
  valueTypes := []string{"var", "const", "let"}
  valueEnds := []string{"=", ","}
  for i := 0; i < len(codes); i++ {
    line := codes[i]
    for j := 0; j < len(valueTypes); j++ {
      valueType := valueTypes[j]
      valueStartIndex := strings.Index(line, valueType)

      for k := 0; k < len(valueEnds); k++ {
        valueEnd := valueEnds[k]
        valueEndIndex := strings.Index(line, valueEnd)

        if valueStartIndex > 0 && valueEndIndex > 0{
          fmt.Println("valueStartIndex", valueStartIndex, "line", line)
          match, _ := regexp.MatchString(`([A-Za-z0-9\_]+)[ \t]{0,3}\=[^<>!]`, line)
          fmt.Println(match)
        }
      }
    }
  }
}

func (p Javascript) ReadFunctions(file model.File) []model.Keyword {
  keywords := []model.Keyword{}
  code := parser.ReadSourceCode(file.Path)
  codes := strings.Split(code, "\n")
  for i := 0; i < len(codes); i++ {
    line := codes[i]
    funcStartIndex := findFuncStartIndex(line)
    if funcStartIndex > 0 {
      funcEndIndex := findFuncEndIndex(line, funcStartIndex)
      funcCode := getFuncLine(line, funcStartIndex, funcEndIndex)
      if isFuncLine(funcCode)  {
        keyword := model.Keyword {
          Type: "function",
          Name: getFuncName(funcCode),
        }
        keywords = append(keywords, keyword)
      }
    }
  }
  return keywords
}

func (p Javascript) ReadClasses(file model.File) {
  //fmt.Println("Javascript ReadClasses")
}

func (p Javascript) ReadInterfaces(file model.File) {
  //fmt.Println("Javascript ReadInterfaces")
}
