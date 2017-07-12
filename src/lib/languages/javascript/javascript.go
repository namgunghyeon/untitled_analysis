package javascript

import (
  "lib/parser"
  "lib/model"
  "strings"
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
