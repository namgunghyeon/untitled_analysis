package scanner
import (
  "path/filepath"
  "os"
  "fmt"
)

func ignorePath(path string, ignore string) {}

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  return nil
}

func Scan(path string) string {
  err := filepath.Walk(path, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
  return path
}
