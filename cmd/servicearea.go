package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var serviceAreaCmd = &cobra.Command{
	Use:   "service-area",
	Short: "List service areas",
}

var serviceAreaLargeCmd = &cobra.Command{
	Use:   "large",
	Short: "List large service areas",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.LargeServiceAreaResponse
		if err := client.Get("/large_service_area/v1/", nil, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME"})
		for _, a := range resp.Results.LargeServiceAreas {
			tw.Row(a.Code, a.Name)
		}
		tw.Flush()
		return nil
	},
}

var serviceAreaListCmd = &cobra.Command{
	Use:   "list",
	Short: "List service areas",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.ServiceAreaResponse
		if err := client.Get("/service_area/v1/", nil, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME", "LARGE SERVICE AREA"})
		for _, a := range resp.Results.ServiceAreas {
			tw.Row(a.Code, a.Name, a.LargeServiceArea.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serviceAreaCmd)
	serviceAreaCmd.AddCommand(serviceAreaLargeCmd)
	serviceAreaCmd.AddCommand(serviceAreaListCmd)
}
