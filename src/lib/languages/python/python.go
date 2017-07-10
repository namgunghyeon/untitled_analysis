package python

import (
  "fmt"
  "lib/model"
)

type Python struct {

}

func (p Python) ReadValues(file model.File) {
  fmt.Println("Python ReadValues")
}

func (p Python) ReadFunctions(file model.File) {
  fmt.Println("Python ReadFunctions")
}

func (p Python) ReadClasses(file model.File) {
  fmt.Println("Python ReadClasses")
}

func (p Python) ReadInterfaces(file model.File) {
  fmt.Println("Python ReadInterfaces")
}
