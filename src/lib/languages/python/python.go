package python

import "fmt"

type Python struct {

}

func (p Python) ReadValues() {
  fmt.Println("Python ReadValues")
}

func (p Python) ReadFunctions() {
  fmt.Println("Python ReadFunctions")
}

func (p Python) ReadClasses() {
  fmt.Println("Python ReadClasses")
}

func (p Python) ReadInterfaces() {
  fmt.Println("Python ReadInterfaces")
}

func ReadValues2() {
  fmt.Println("Python ReadValues2")
}
