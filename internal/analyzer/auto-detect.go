package analyzer

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
  "strings"
)

func isIgnoredDirectory(ignoreDirs *[]string, fileInfo *os.FileInfo) bool {
  for _, ignoredDir := range *ignoreDirs {
    if (*fileInfo).IsDir() && (*fileInfo).Name() == ignoredDir {
      return true
    }
  }

  return false
}

func AutoDetectComponents() *string {
  projectPath, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  fileExtensions := [3]string{".tsx", ".js", ".jsx"}
  ignoreDirs := []string{"node_modules"}
  var potentialComponents []string

  err = filepath.Walk(projectPath,
    func(path string, fileInfo os.FileInfo, err error) error {
      if err != nil {
        return err
      }

      if (isIgnoredDirectory(&ignoreDirs, &fileInfo)) {
        return filepath.SkipDir
      }

      for _, ignoredDir := range ignoreDirs {
        if fileInfo.IsDir() && fileInfo.Name() == ignoredDir {
          return filepath.SkipDir
        }
      }
      
      for _, ext := range fileExtensions {
        if strings.HasSuffix(fileInfo.Name(), ext) {
          potentialComponents = append(potentialComponents, fileInfo.Name())
        }
      }

      return nil
    })
  
  if err != nil {
    log.Fatal()
  }

  for _, potentialComponent := range potentialComponents {
    fmt.Println(potentialComponent)
  }

  tmp := ""
  return &tmp
}
