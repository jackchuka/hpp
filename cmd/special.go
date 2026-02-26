package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var specialCmd = &cobra.Command{
	Use:   "special",
	Short: "List specials and features",
}

var specialParams api.SpecialParams

var specialListCmd = &cobra.Command{
	Use:   "list",
	Short: "List specials/features",
	Example: `  hpp special list
  hpp special list --category SPC0`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.SpecialResponse
		if err := client.Get("/special/v1/", specialParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME", "CATEGORY"})
		for _, s := range resp.Results.Specials {
			tw.Row(s.Code, s.Name, s.SpecialCategory.Name)
		}
		tw.Flush()
		return nil
	},
}

var specialCategoryParams api.SpecialCategoryParams

var specialCategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "List special categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.SpecialCategoryResponse
		if err := client.Get("/special_category/v1/", specialCategoryParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME"})
		for _, c := range resp.Results.SpecialCategories {
			tw.Row(c.Code, c.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(specialCmd)
	specialCmd.AddCommand(specialListCmd)
	specialCmd.AddCommand(specialCategoryCmd)

	specialListCmd.Flags().StringSliceVar(&specialParams.Special, "code", nil, "special codes")
	specialListCmd.Flags().StringSliceVar(&specialParams.SpecialCategory, "category", nil, "filter by category codes")

	specialCategoryCmd.Flags().StringSliceVar(&specialCategoryParams.SpecialCategory, "code", nil, "category codes")
}
