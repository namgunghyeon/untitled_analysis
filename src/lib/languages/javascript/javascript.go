package javascript

import (
  "fmt"
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
  return len(line) - 1
}

func getFuncLine(line string, startIndex int, endIndex int) string {
  funcCode := string(line[startIndex:endIndex])
  funcCodes := strings.Split(funcCode," ")
  if len(funcCodes) > 1 && string(funcCodes[1]) != "" {
    return funcCode
  }
  return ""
}

func (p Javascript) ReadValues(file model.File) {
}

func (p Javascript) ReadFunctions(file model.File) {
  code := parser.ReadSourceCode(file.Path)
  codes := strings.Split(code,"\n")
  for i := 0; i < len(codes); i++ {
    line := codes[i]
    funcStartIndex := findFuncStartIndex(line)
    if funcStartIndex > 0 {
      funcEndIndex := findFuncEndIndex(line, funcStartIndex)
      funcCode := getFuncLine(line, funcStartIndex, funcEndIndex)
      if funcCode != "" {
        fmt.Println("funcCode", funcCode)
      }
    }
  }
}

func (p Javascript) ReadClasses(file model.File) {
  //fmt.Println("Javascript ReadClasses")
}

func (p Javascript) ReadInterfaces(file model.File) {
  //fmt.Println("Javascript ReadInterfaces")
}
