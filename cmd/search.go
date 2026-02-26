package cmd

import (
	"fmt"
	"os"

	"github.com/jackchuka/hpp/internal/api"
	"github.com/jackchuka/hpp/internal/output"
	"github.com/spf13/cobra"
)

// Local variables for cobra flag binding (pointer fields need indirection)
var (
	searchKeyword          string
	searchName             string
	searchNameKana         string
	searchNameAny          string
	searchTel              string
	searchAddress          string
	searchLat              float64
	searchLng              float64
	searchRange            int
	searchDatum            string
	searchLargeServiceArea string
	searchPartyCapacity    int
	searchKtaiCoupon       int
	searchType             string
	searchOrder            int
	searchStart            int
	searchCount            int
)

var searchParams api.GourmetSearchParams

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search restaurants",
	Long:  "Search restaurants using the HotPepper Gourmet API with various filters.",
	Example: `  hpp search --keyword "ramen" --area Z011
  hpp search --lat 35.6812 --lng 139.7671 --range 3
  hpp search --keyword "izakaya" --wifi --private-room --english`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Populate pointer fields only when flags were explicitly set
		if cmd.Flags().Changed("keyword") {
			searchParams.Keyword = &searchKeyword
		}
		if cmd.Flags().Changed("name") {
			searchParams.Name = &searchName
		}
		if cmd.Flags().Changed("name-kana") {
			searchParams.NameKana = &searchNameKana
		}
		if cmd.Flags().Changed("name-any") {
			searchParams.NameAny = &searchNameAny
		}
		if cmd.Flags().Changed("tel") {
			searchParams.Tel = &searchTel
		}
		if cmd.Flags().Changed("address") {
			searchParams.Address = &searchAddress
		}
		if cmd.Flags().Changed("lat") {
			searchParams.Lat = &searchLat
		}
		if cmd.Flags().Changed("lng") {
			searchParams.Lng = &searchLng
		}
		if cmd.Flags().Changed("range") {
			searchParams.Range = &searchRange
		}
		if cmd.Flags().Changed("datum") {
			searchParams.Datum = &searchDatum
		}
		if cmd.Flags().Changed("large-service-area") {
			searchParams.LargeServiceArea = &searchLargeServiceArea
		}
		if cmd.Flags().Changed("party-capacity") {
			searchParams.PartyCapacity = &searchPartyCapacity
		}
		if cmd.Flags().Changed("ktai-coupon") {
			searchParams.KtaiCoupon = &searchKtaiCoupon
		}
		if cmd.Flags().Changed("type") {
			searchParams.Type = &searchType
		}
		if cmd.Flags().Changed("order") {
			searchParams.Order = &searchOrder
		}
		if cmd.Flags().Changed("start") {
			searchParams.Start = &searchStart
		}
		if cmd.Flags().Changed("count") {
			searchParams.Count = &searchCount
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("HOTPEPPER_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("HOTPEPPER_API_KEY environment variable is required")
		}
		client := api.NewClient(apiKey)
		var resp api.GourmetResponse
		if err := client.Get("/gourmet/v1/", searchParams, &resp); err != nil {
			return err
		}

		if outputFormat == "json" {
			return output.WriteJSON(os.Stdout, resp)
		}

		fmt.Fprintf(os.Stderr, "Found %d results (showing %s)\n\n",
			resp.Results.ResultsAvailable, resp.Results.ResultsReturned)

		tw := output.NewTableWriter(os.Stdout, []string{"NAME", "GENRE", "AREA", "ACCESS", "BUDGET", "URL"})
		for _, s := range resp.Results.Shops {
			tw.Row(s.Name, s.Genre.Name, s.MiddleArea.Name, s.Access, s.Budget.Average, s.URLs.PC)
		}
		tw.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	f := searchCmd.Flags()

	// Text search
	f.StringVar(&searchKeyword, "keyword", "", "free text search")
	f.StringVar(&searchName, "name", "", "shop name (partial match)")
	f.StringVar(&searchNameKana, "name-kana", "", "shop name in kana")
	f.StringVar(&searchNameAny, "name-any", "", "shop name or kana")
	f.StringVar(&searchTel, "tel", "", "phone number (digits only)")
	f.StringVar(&searchAddress, "address", "", "address (partial match)")

	// Location
	f.Float64Var(&searchLat, "lat", 0, "latitude")
	f.Float64Var(&searchLng, "lng", 0, "longitude")
	f.IntVar(&searchRange, "range", 0, "search range: 1=300m 2=500m 3=1km 4=2km 5=3km")
	f.StringVar(&searchDatum, "datum", "", "geodetic system: world or tokyo")

	// Area filters
	f.StringVar(&searchLargeServiceArea, "large-service-area", "", "large service area code")
	f.StringSliceVar(&searchParams.ServiceArea, "service-area", nil, "service area codes")
	f.StringSliceVar(&searchParams.LargeArea, "area", nil, "large area codes")
	f.StringSliceVar(&searchParams.MiddleArea, "middle-area", nil, "middle area codes")
	f.StringSliceVar(&searchParams.SmallArea, "small-area", nil, "small area codes")

	// Category
	f.StringSliceVar(&searchParams.Genre, "genre", nil, "genre codes")
	f.StringSliceVar(&searchParams.Budget, "budget", nil, "budget codes")
	f.StringSliceVar(&searchParams.CreditCardFilter, "credit-card", nil, "credit card codes")
	f.StringSliceVar(&searchParams.Special, "special", nil, "special codes (AND)")
	f.StringSliceVar(&searchParams.SpecialOr, "special-or", nil, "special codes (OR)")
	f.StringSliceVar(&searchParams.SpecialCategory, "special-category", nil, "special category codes (AND)")
	f.StringSliceVar(&searchParams.SpecialCategoryOr, "special-category-or", nil, "special category codes (OR)")

	// Capacity
	f.IntVar(&searchPartyCapacity, "party-capacity", 0, "min banquet capacity")

	// Boolean filters
	f.BoolVar(&searchParams.WiFi, "wifi", false, "has WiFi")
	f.BoolVar(&searchParams.Wedding, "wedding", false, "wedding/party inquiry")
	f.BoolVar(&searchParams.Course, "course", false, "has courses")
	f.BoolVar(&searchParams.FreeDrink, "free-drink", false, "all-you-can-drink")
	f.BoolVar(&searchParams.FreeFood, "free-food", false, "all-you-can-eat")
	f.BoolVar(&searchParams.PrivateRoom, "private-room", false, "has private rooms")
	f.BoolVar(&searchParams.Horigotatsu, "horigotatsu", false, "sunken kotatsu seating")
	f.BoolVar(&searchParams.Tatami, "tatami", false, "tatami seating")
	f.BoolVar(&searchParams.Cocktail, "cocktail", false, "cocktail selection")
	f.BoolVar(&searchParams.Shochu, "shochu", false, "shochu selection")
	f.BoolVar(&searchParams.Sake, "sake", false, "sake selection")
	f.BoolVar(&searchParams.Wine, "wine", false, "wine selection")
	f.BoolVar(&searchParams.Card, "card", false, "accepts cards")
	f.BoolVar(&searchParams.NonSmoking, "non-smoking", false, "non-smoking seats")
	f.BoolVar(&searchParams.Charter, "charter", false, "private rental")
	f.BoolVar(&searchParams.Ktai, "ktai", false, "mobile phone OK")
	f.BoolVar(&searchParams.Parking, "parking", false, "has parking")
	f.BoolVar(&searchParams.BarrierFree, "barrier-free", false, "barrier-free access")
	f.BoolVar(&searchParams.Sommelier, "sommelier", false, "has sommelier")
	f.BoolVar(&searchParams.NightView, "night-view", false, "scenic night view")
	f.BoolVar(&searchParams.OpenAir, "open-air", false, "open-air seating")
	f.BoolVar(&searchParams.Show, "show", false, "live/show")
	f.BoolVar(&searchParams.Equipment, "equipment", false, "entertainment equipment")
	f.BoolVar(&searchParams.Karaoke, "karaoke", false, "has karaoke")
	f.BoolVar(&searchParams.Band, "band", false, "band performance OK")
	f.BoolVar(&searchParams.TV, "tv", false, "has TV/projector")
	f.BoolVar(&searchParams.Lunch, "lunch", false, "lunch service")
	f.BoolVar(&searchParams.Midnight, "midnight", false, "open after 11pm")
	f.BoolVar(&searchParams.MidnightMeal, "midnight-meal", false, "food after 11pm")
	f.BoolVar(&searchParams.English, "english", false, "English menu")
	f.BoolVar(&searchParams.Pet, "pet", false, "pet allowed")
	f.BoolVar(&searchParams.Child, "child", false, "children welcome")

	// Mobile coupon
	f.IntVar(&searchKtaiCoupon, "ktai-coupon", -1, "mobile coupon: 0=with 1=without")

	// Output control
	f.StringVar(&searchType, "type", "", "response type: lite, credit_card, special")
	f.IntVar(&searchOrder, "order", 0, "sort: 1=name 2=genre 3=area 4=recommended")
	f.IntVar(&searchStart, "start", 0, "result start position")
	f.IntVar(&searchCount, "count", 0, "results per page (max 100)")
}
