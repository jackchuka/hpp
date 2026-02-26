package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
)

var rootCmd = &cobra.Command{
	Use:   "hpp",
	Short: "HotPepper Gourmet API CLI",
	Long:  "Search Japanese restaurants using the HotPepper Gourmet API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&outputFormat, "format", "json", "output format: table or json")
}
