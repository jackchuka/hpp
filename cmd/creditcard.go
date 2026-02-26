package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var creditcardCmd = &cobra.Command{
	Use:   "creditcard",
	Short: "List accepted credit card types",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.CreditCardResponse
		if err := client.Get("/credit_card/v1/", nil, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME"})
		for _, c := range resp.Results.CreditCards {
			tw.Row(c.Code, c.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(creditcardCmd)
}
