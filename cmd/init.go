package cmd

import (
  "fmt"
  "github.com/oliviergoulet5/migrate-component/internal/config"
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
  
  // If the user does not have a configuration file, create one.
  if (!config.HasConfig()) {
    config.CreateConfig()
  }

  configFile := config.GetConfig()
  config.AppendMigrationToConfig(configFile, &migratingFrom, &migratingTo)

  defer configFile.Close()
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
