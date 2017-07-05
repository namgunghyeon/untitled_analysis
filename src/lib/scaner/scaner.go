package scaner
import (
  "path/filepath"
  "os"
  "log"
  "lib/model"
)

func visitDirs(ignoreDirs []string, scan *model.Scan) filepath.WalkFunc {
  return func(path string, info os.FileInfo, err error) error {
    if err != nil {
      log.Print(err)
      return nil
    }
    if info.IsDir() {
      dir := filepath.Base(path)
      for _, d := range ignoreDirs {
        if d == dir {
          return filepath.SkipDir
        }
      }
    }
    scan.Paths = append(scan.Paths, path)
    return nil
  }
}

func Scan(path string) model.Scan {
  ignoreDirs := []string{".bzr", ".hg", ".git"}
  scan := model.Scan{
    Paths: []string{},
  };
  err := filepath.Walk(path, visitDirs(ignoreDirs, &scan))
  if err != nil {
    log.Print(err)
    return scan
  }
  return scan
}
