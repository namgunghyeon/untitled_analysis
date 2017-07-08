package javascript

import "fmt"

type Javascript struct {

}

func (p Javascript) ReadValues() {
  fmt.Println("Javascript ReadValues")
}

func (p Javascript) ReadFunctions() {
  fmt.Println("Javascript ReadFunctions")
}

func (p Javascript) ReadClasses() {
  fmt.Println("Javascript ReadClasses")
}

func (p Javascript) ReadInterfaces() {
  fmt.Println("Javascript ReadInterfaces")
}
