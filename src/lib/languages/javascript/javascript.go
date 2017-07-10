package javascript

import (
  "fmt"
  "lib/parser"
  "lib/model"
)

type Javascript struct {

}

func (p Javascript) ReadValues(file model.File) {
  parser.ReadSourceCode(file.Path)
  fmt.Println("Javascript ReadValues")
}

func (p Javascript) ReadFunctions(file model.File) {
  fmt.Println("Javascript ReadFunctions")
}

func (p Javascript) ReadClasses(file model.File) {
  fmt.Println("Javascript ReadClasses")
}

func (p Javascript) ReadInterfaces(file model.File) {
  fmt.Println("Javascript ReadInterfaces")
}
