package config 

import (
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
  "github.com/oliviergoulet5/migrate-component/internal/models"
  "gopkg.in/yaml.v3"
)

func getConfigFilePath() string {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  }

  configFilePath := filepath.Join(homeDir, ".config/migrate-component.yml")

  return configFilePath
}

func GetConfig() *models.Config {
  configFilePath := getConfigFilePath()

  configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0644)
  if err != nil {
    log.Fatal(err)
  }

  content, err := ioutil.ReadAll(configFile)
  if err != nil {
    log.Fatal(err)
  }
  
  var configuration models.Config
  if err := yaml.Unmarshal([]byte(content), &configuration); err != nil {
    log.Fatal(err)
  }

  defer configFile.Close()

  return &configuration
}

// Check whether or not the user has a configuration file in their home
// directory.
// Returns:
//  A boolean value indicating if user has a configuration file.
func HasConfigFile() bool {
  configInfo, err := os.Stat(getConfigFilePath())
  if err != nil {
    log.Fatal(err)
  }

  return configInfo.IsDir()
}

// Creates a configuration file under ~/.config.
func CreateConfigFile() {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  }

  configDir := filepath.Join(homeDir, ".config")
  configDirInfo, err := os.Stat(configDir)
  if err != nil {
    log.Fatal(err)
  }

  if !configDirInfo.IsDir() {
    // create it
    if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
      log.Fatal(err)
    }
  }
  
  configFilePath := filepath.Join(homeDir, ".config/migrate-component.yml")
  _, err = os.Stat(configFilePath)
  // If an error does not exist, assume the file exists. Since it exists, there
  // is nothing to create. Return early.
  if err == nil {
    return
  }

  file, err := os.Create(configFilePath) 
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()
}

// Check whether a project is already being migrated.
// Parameters:
//  config: The configuration.
//  migrationPath: The path which the new migration will be.
// Returns:
//  A boolean representing if the migration already exists.
func checkMigrationAlreadyExists(configuration *models.Config, migrationPath *string) bool {
  for _, migration := range configuration.Migrations {
    if migration.ProjectPath == *migrationPath {
      return true
    }
  }

  return false
}

// Appends a migration entry to the configuration file. It includes the project
// path, the migration from and to.
// Parameters:
//  configFile: The configuration file.
//  from: The component library that the user is migrating away from.
//  to: The component library that the user is migrating towards.
func AppendMigrationToConfigFile(from *string, to *string) {
  configuration := GetConfig()

  cwd, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  if checkMigrationAlreadyExists(configuration, &cwd) {
    log.Fatal("A migration is already in progress for the current working directory.")
  }

  migration := models.Migration{
    cwd,
    *from,
    *to,
    make(map[string]bool),
  }

  configuration.Migrations = append(configuration.Migrations, migration)

  yml, err := yaml.Marshal(configuration)
  if err != nil {
    log.Fatal(err)
  }
  
  configFile, err := os.OpenFile(getConfigFilePath(), os.O_TRUNC|os.O_WRONLY, 0644)
  _, err = configFile.Write(yml)
  if err != nil {
    log.Fatal(err)
  }

  defer configFile.Close()
}

