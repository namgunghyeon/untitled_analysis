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

func findValueStartIndex(line string, valueType string) int{
  return strings.Index(line, valueType)
}

func findValueEndIndex(line string, startIndex int, valueEnds []string) int{
  for k := 0; k < len(valueEnds); k++ {
    valueEnd := valueEnds[k]
    valueEndIndex := strings.Index(line, valueEnd)
    if valueEndIndex > startIndex {
      return valueEndIndex
    }
  }
  return startIndex
}

func getValueLine(line string, startIndex int, endIndex int) []string{
  typeCode := string(line[startIndex:endIndex])
  tpyeCodes := strings.Split(typeCode, " ")
  removedType := tpyeCodes[1: len(tpyeCodes)]
  return removedType
}

func findTypeValue(line string, valueTypes []string, valueEnds []string) []string{
  values := []string{}
  for j := 0; j < len(valueTypes); j++ {
    valueType := valueTypes[j]
    valueStartIndex := findValueStartIndex(line, valueType)
    valueEndIndex := findValueEndIndex(line, valueStartIndex, valueEnds)
    if valueStartIndex > 0 && valueEndIndex > 0{
      values = getValueLine(line, valueStartIndex, valueEndIndex)
    }
  }
  return values
}

func (p Javascript) ReadValues(file model.File) []model.Keyword{
  keywords := []model.Keyword{}
  code := parser.ReadSourceCode(file.Path)
  codes := strings.Split(code, "\n")
  valueTypes := []string{"var", "const", "let"}
  valueEnds := []string{"=", ","}
  for i := 0; i < len(codes); i++ {
    line := codes[i]
    values := findTypeValue(line, valueTypes, valueEnds)
    if len(values) > 0 {
      for j := 0; j < len(values); j++ {
        if values[j] != "" {
          keyword := model.Keyword {
            Type: "value",
            Path: file.Path,
            Name: values[j],
          }
          keywords = append(keywords, keyword)
        }
      }
    }
  }
  return keywords
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
          Path: file.Path,
          Name: getFuncName(funcCode),
        }
        keywords = append(keywords, keyword)
      }
    }
  }
  return keywords
}


func findClassStartIndex(line string) int{
  return strings.Index(line, "class")
}

func findClassEnIndex(line string, startIndex int) int{
  for j := startIndex; j < len(line); j++ {
    if string(line[j]) == "{" {
      return j
    }
  }
  return startIndex
}

func getClassLine(line string, startIndex int, endIndex int) string {
  classCode := string(line[startIndex:endIndex])
  classCodes := strings.Split(classCode, " ")
  if len(classCode) > 1 && string(classCodes[1]) != "" {
    return classCode
  }
  return ""
}

func isClassLine(classCode string) bool{
  return len(strings.Split(classCode, " ")) >= 2
}

func getClassName(classCode string) string {
  return string(strings.Split(classCode, " ")[1])
}

func (p Javascript) ReadClasses(file model.File) []model.Keyword {
  keywords := []model.Keyword{}
  code := parser.ReadSourceCode(file.Path)
  codes := strings.Split(code, "\n")
  for i := 0; i < len(codes); i++ {
    line := strings.Trim(codes[i], "")
    classStartIndex := findClassStartIndex(line)
    if classStartIndex == 0 {
      classEndIndex := findClassEnIndex(line, classStartIndex)
      classCode := getClassLine(line, classStartIndex, classEndIndex)
      if isClassLine(classCode) {
        keyword := model.Keyword {
          Type: "class",
          Path: file.Path,
          Name: getClassName(classCode),
        }
        keywords = append(keywords, keyword)
      }
    }
  }
  return keywords
}

func (p Javascript) ReadInterfaces(file model.File) []model.Keyword {
  keywords := []model.Keyword{}
  return keywords
}
