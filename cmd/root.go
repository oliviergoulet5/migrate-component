package cmd

import (
  "os"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "migrate-component",
  Short: "A program to do migrations on your React components.",
}

func Execute() {
  err := rootCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  rootCmd.AddCommand(versionCmd)
  rootCmd.AddCommand(initCmd)
}
