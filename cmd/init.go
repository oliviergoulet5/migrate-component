package cmd

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "github.com/oliviergoulet5/migrate-component/internal/models"
  "github.com/spf13/cobra"
  "github.com/manifoldco/promptui"
  "gopkg.in/yaml.v3"
)

func runInit() {
  prompt := promptui.Select{
    Label: "Select what UI component library you are migrating from",
    Items: []string{
      "MaterialUI",
      "Other"},
  }

  fmt.Printf("Migrating from:\n")
  _, migratingFrom, err := prompt.Run()
  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }


  prompt = promptui.Select{
    Label: "Select what UI component library you are migrating to",
    Items: []string{
      "Other"},
    }

  fmt.Printf("Migrating to:\n")
  _, migratingTo, err := prompt.Run()
  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }

  _ = migratingFrom
  _ = migratingTo
  
  // If the user does not have a configuration file, create one.
  if (!hasConfig()) {
    createConfig()
  }

  config := getConfig()
  appendMigrationToConfig(config, &migratingFrom, &migratingTo)

  defer config.Close()
}

func getConfig() *os.File {
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

func hasConfig() bool {
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

func createConfig() {
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

func appendMigrationToConfig(configFile *os.File, from *string, to *string) {
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

var initCmd = &cobra.Command{
  Use: "init",
  Short: "Initialize a migration for a project.",
  Long: `
    Initialize a migration for a project.
  `,
  Run: func(cmd *cobra.Command, args []string) {
    runInit()
  },
}
