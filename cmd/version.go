package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func getVersion() string {
  return "0.0.1"
}

var versionCmd = &cobra.Command{
  Use: "version",
  Short: "Display version information.",
  Long: `
The version command provides information about the application's version.
  `,
  Run: func(cmd *cobra.Command, args []string) {
    version := getVersion()
    fmt.Printf("Migrate Component version: %v\n", version)
  },
}
