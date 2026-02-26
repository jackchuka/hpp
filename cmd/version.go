package cmd

import (
	"fmt"

	"github.com/jackchuka/hpp/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hpp %s (commit: %s, built: %s)\n", version.Version, version.Commit, version.BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
