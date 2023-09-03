package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of PintaDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PintaDB v0.1.0")
	},
}
