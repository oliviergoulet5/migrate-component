package cmd

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/spf13/cobra"
  "github.com/manifoldco/promptui"
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

  createConfig()
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
