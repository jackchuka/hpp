package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

// Area parent command
var areaCmd = &cobra.Command{
	Use:   "area",
	Short: "List geographic areas",
	Long:  "List large, middle, or small geographic areas.",
}

// --- large area ---
var largeAreaParams api.LargeAreaParams
var largeAreaKeyword string

var areaLargeCmd = &cobra.Command{
	Use:   "large",
	Short: "List large areas",
	Example: `  hpp area large
  hpp area large --keyword tokyo`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("keyword") {
			largeAreaParams.Keyword = &largeAreaKeyword
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.LargeAreaResponse
		if err := client.Get("/large_area/v1/", largeAreaParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME", "SERVICE AREA"})
		for _, a := range resp.Results.LargeAreas {
			tw.Row(a.Code, a.Name, a.ServiceArea.Name)
		}
		tw.Flush()
		return nil
	},
}

// --- middle area ---
var middleAreaParams api.MiddleAreaParams
var middleAreaKeyword string
var middleAreaStart, middleAreaCount int

var areaMiddleCmd = &cobra.Command{
	Use:   "middle",
	Short: "List middle areas",
	Example: `  hpp area middle --large-area Z011
  hpp area middle --keyword shinjuku`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("keyword") {
			middleAreaParams.Keyword = &middleAreaKeyword
		}
		if cmd.Flags().Changed("start") {
			middleAreaParams.Start = &middleAreaStart
		}
		if cmd.Flags().Changed("count") {
			middleAreaParams.Count = &middleAreaCount
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.MiddleAreaResponse
		if err := client.Get("/middle_area/v1/", middleAreaParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME", "LARGE AREA"})
		for _, a := range resp.Results.MiddleAreas {
			tw.Row(a.Code, a.Name, a.LargeArea.Name)
		}
		tw.Flush()
		return nil
	},
}

// --- small area ---
var smallAreaParams api.SmallAreaParams
var smallAreaKeyword string
var smallAreaStart, smallAreaCount int

var areaSmallCmd = &cobra.Command{
	Use:   "small",
	Short: "List small areas",
	Example: `  hpp area small --middle-area Y005
  hpp area small --keyword ginza`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("keyword") {
			smallAreaParams.Keyword = &smallAreaKeyword
		}
		if cmd.Flags().Changed("start") {
			smallAreaParams.Start = &smallAreaStart
		}
		if cmd.Flags().Changed("count") {
			smallAreaParams.Count = &smallAreaCount
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.SmallAreaResponse
		if err := client.Get("/small_area/v1/", smallAreaParams, &resp); err != nil {
			return err
		}
		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}
		tw := output.NewTableWriter(os.Stdout, []string{"CODE", "NAME", "MIDDLE AREA"})
		for _, a := range resp.Results.SmallAreas {
			tw.Row(a.Code, a.Name, a.MiddleArea.Name)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(areaCmd)
	areaCmd.AddCommand(areaLargeCmd)
	areaCmd.AddCommand(areaMiddleCmd)
	areaCmd.AddCommand(areaSmallCmd)

	// large area flags
	areaLargeCmd.Flags().StringSliceVar(&largeAreaParams.LargeArea, "code", nil, "large area codes")
	areaLargeCmd.Flags().StringVar(&largeAreaKeyword, "keyword", "", "area name search")

	// middle area flags
	areaMiddleCmd.Flags().StringSliceVar(&middleAreaParams.MiddleArea, "code", nil, "middle area codes")
	areaMiddleCmd.Flags().StringSliceVar(&middleAreaParams.LargeArea, "large-area", nil, "filter by large area codes")
	areaMiddleCmd.Flags().StringVar(&middleAreaKeyword, "keyword", "", "area name search")
	areaMiddleCmd.Flags().IntVar(&middleAreaStart, "start", 0, "result start position")
	areaMiddleCmd.Flags().IntVar(&middleAreaCount, "count", 0, "results per page")

	// small area flags
	areaSmallCmd.Flags().StringSliceVar(&smallAreaParams.SmallArea, "code", nil, "small area codes")
	areaSmallCmd.Flags().StringSliceVar(&smallAreaParams.MiddleArea, "middle-area", nil, "filter by middle area codes")
	areaSmallCmd.Flags().StringVar(&smallAreaKeyword, "keyword", "", "area name search")
	areaSmallCmd.Flags().IntVar(&smallAreaStart, "start", 0, "result start position")
	areaSmallCmd.Flags().IntVar(&smallAreaCount, "count", 0, "results per page")
}
