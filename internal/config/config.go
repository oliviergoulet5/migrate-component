package config 

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "github.com/oliviergoulet5/migrate-component/internal/models"
  "gopkg.in/yaml.v3"
)

// Retrieve the configuration file from the user's home directory.
// Returns:
//  The configuration file.
func GetConfigFile() *os.File {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("Error:", err)
    panic(err)
  }

  configFilePath := filepath.Join(homeDir, ".config/migrate-component.yml")

  file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    panic(err)
  }
  
  return file
}

// Check whether or not the user has a configuration file in their home
// directory.
// Returns:
//  A boolean value indicating if user has a configuration file.
func HasConfigFile() bool {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("Error:", err)
    panic(err)
  }
  
  configFilePath := filepath.Join(homeDir, ".config/migrate-component.yml")

  configInfo, err := os.Stat(configFilePath)
  if err != nil {
    panic(err)
  }

  return configInfo.IsDir()
}

// Creates a configuration file under ~/.config.
func CreateConfigFile() {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  configDir := filepath.Join(homeDir, ".config")
  configDirInfo, err := os.Stat(configDir)
  if err != nil {
    fmt.Println("Error, could not find config dir")
    return
  }

  if !configDirInfo.IsDir() {
    // create it
    if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
      fmt.Println("Error when creating ~/.config")
      return
    }
  }
  
  configFilePath := filepath.Join(homeDir, ".config/migrate-component.yml")
  file, err := os.Create(configFilePath) 
  if err != nil {
    fmt.Println("Error when creating migrate-component.yml in config directory")
    return
  }

  defer file.Close()
}

// Appends a migration entry to the configuration file. It includes the project
// path, the migration from and to.
// Parameters:
//  configFile: The configuration file.
//  from: The component library that the user is migrating away from.
//  to: The component library that the user is migrating towards.
func AppendMigrationToConfigFile(configFile *os.File, from *string, to *string) {
  cwd, err := os.Getwd()
  if err != nil {
    return
  }

  fileContent, err := ioutil.ReadAll(configFile)
  if err != nil {
    panic(err)
  }

  migration := models.Migration{
    cwd,
    *from,
    *to,
  }

  var configYml models.Config
  if err := yaml.Unmarshal([]byte(fileContent), &configYml); err != nil {
    panic(err)
  }

  configYml.Migrations = append(configYml.Migrations, migration)

  yml, err := yaml.Marshal(configYml)
  if err != nil {
    panic(err)
  }

  _, err = configFile.Write(yml)
  if err != nil {
    panic(err)
  }
}

