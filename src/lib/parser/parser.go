package parser

import (
    "fmt"
    "io/ioutil"
)

func ReadSourceCode(path string) string{
  b, err := ioutil.ReadFile(path)
   if err != nil {
       fmt.Print(err)
   }
   str := string(b)
   fmt.Println(str)
   return str
}
