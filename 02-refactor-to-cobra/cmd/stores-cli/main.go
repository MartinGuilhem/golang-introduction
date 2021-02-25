package main

import (
	"github.com/spf13/cobra"
	"golang-introduction/02-refactor-to-cobra/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "stores-cli"}
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
