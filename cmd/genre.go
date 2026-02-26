package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var genreParams api.GenreParams
var genreKeyword string

var genreCmd = &cobra.Command{
	Use:   "genre",
	Short: "List cuisine genres",
	Example: `  hpp genre
  hpp genre --keyword ramen
  hpp genre --code G001,G002`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("keyword") {
			genreParams.Keyword = &genreKeyword
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.GenreResponse
		if err := client.Get("/genre/v1/", genreParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME"})
		for _, g := range resp.Results.Genres {
			tw.Row(g.Code, g.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genreCmd)
	genreCmd.Flags().StringSliceVar(&genreParams.Code, "code", nil, "genre codes")
	genreCmd.Flags().StringVar(&genreKeyword, "keyword", "", "genre name search")
}
