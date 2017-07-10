package java

import (
  "fmt"
  "lib/model"
)

type Java struct {}

func (p Java) ReadValues(file model.File) {
  fmt.Println("Java ReadValues")
}

func (p Java) ReadValues(file model.File) {
  fmt.Println("Java ReadValues")
}

func (p Java) ReadFunctions(file model.File) {
  fmt.Println("Java ReadFunctions")
}

func (p Java) ReadClasses(file model.File) {
  fmt.Println("Java ReadClasses")
}

func (p Java) ReadInterfaces(file model.File) {
  fmt.Println("Java ReadInterfaces")
}
