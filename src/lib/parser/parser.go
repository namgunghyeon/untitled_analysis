package parser

import (
    "io/ioutil"
)

func ReadSourceCode(path string) string{
  code, _ := ioutil.ReadFile(path)
  return string(code)
}
