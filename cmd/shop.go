package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

var (
	shopKeyword string
	shopTel     string
	shopStart   int
	shopCount   int
)

var shopParams api.ShopSearchParams

var shopCmd = &cobra.Command{
	Use:   "shop",
	Short: "Search shops by name or phone",
	Long:  "Search restaurants by name or phone number using the HotPepper Shop API.",
	Example: `  hpp shop --keyword "居酒屋"
  hpp shop --tel 0312345678`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("keyword") {
			shopParams.Keyword = &shopKeyword
		}
		if cmd.Flags().Changed("tel") {
			shopParams.Tel = &shopTel
		}
		if cmd.Flags().Changed("start") {
			shopParams.Start = &shopStart
		}
		if cmd.Flags().Changed("count") {
			shopParams.Count = &shopCount
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.ShopSearchResponse
		if err := client.Get("/shop/v1/", shopParams, &resp); err != nil {
			return err
		}

		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}

		fmt.Fprintf(os.Stderr, "Found %d results (showing %s)\n\n",
			resp.Results.ResultsAvailable, resp.Results.ResultsReturned)

		tw := output.NewTableWriter(os.Stdout, []string{"ID", "NAME", "GENRE", "ADDRESS", "URL"})
		for _, s := range resp.Results.Shops {
			tw.Row(s.ID, s.Name, s.Genre.Name, s.Address, s.URLs.PC)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(shopCmd)
	f := shopCmd.Flags()

	f.StringVar(&shopKeyword, "keyword", "", "shop name/kana/address search")
	f.StringVar(&shopTel, "tel", "", "phone number (exact match, digits only)")
	f.IntVar(&shopStart, "start", 0, "result start position")
	f.IntVar(&shopCount, "count", 0, "results per page (max 30)")
}
