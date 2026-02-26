package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "List dinner budget ranges",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.BudgetResponse
		if err := client.Get("/budget/v1/", nil, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME"})
		for _, b := range resp.Results.Budgets {
			tw.Row(b.Code, b.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(budgetCmd)
}
