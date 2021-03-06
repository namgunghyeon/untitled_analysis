package scaner
import (
  "path/filepath"
  "os"
  "lib/model"
)

func visitDirs(ignoreDirs []string, scan *model.Scan) filepath.WalkFunc {
  return func(path string, info os.FileInfo, err error) error {
    if err != nil {
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
    extension := filepath.Ext(path)
    file := model.File{
      Path: path,
      Ext: extension,
    }
    scan.Files = append(scan.Files, file)
    return nil
  }
}

func Scan(path string) model.Scan {
  ignoreDirs := []string{".bzr", ".hg", ".git", "node_modules", "build", "test", ".eslintrc"}
  scan := model.Scan{
    Files: []model.File{},
  };

  err := filepath.Walk(path, visitDirs(ignoreDirs, &scan))
  if err != nil {
    return scan
  }
  return scan
}
