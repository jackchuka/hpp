package api

import (
	"encoding/json"
	"strconv"
)

// FlexInt handles JSON fields that may be int or empty string.
type FlexInt int

func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if string(b) == `""` || string(b) == "null" {
		*fi = 0
		return nil
	}
	var n int
	if err := json.Unmarshal(b, &n); err != nil {
		// Try parsing as quoted number string like "24"
		var s string
		if err2 := json.Unmarshal(b, &s); err2 != nil {
			return err
		}
		parsed, err3 := strconv.Atoi(s)
		if err3 != nil {
			*fi = 0
			return nil
		}
		*fi = FlexInt(parsed)
		return nil
	}
	*fi = FlexInt(n)
	return nil
}

// GourmetResponse is the top-level response from /gourmet/v1/
type GourmetResponse struct {
	Results GourmetResults `json:"results"`
}

type GourmetResults struct {
	APIVersion       string `json:"api_version"`
	ResultsAvailable int    `json:"results_available"`
	ResultsReturned  string `json:"results_returned"`
	ResultsStart     int    `json:"results_start"`
	Shops            []Shop `json:"shop"`
}

type Shop struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	NameKana       string     `json:"name_kana"`
	LogoImage      string     `json:"logo_image"`
	Address        string     `json:"address"`
	StationName    string     `json:"station_name"`
	Lat            float64    `json:"lat"`
	Lng            float64    `json:"lng"`
	Genre          CodeName   `json:"genre"`
	SubGenre       CodeName   `json:"sub_genre"`
	Budget         Budget     `json:"budget"`
	Catch          string     `json:"catch"`
	Capacity       FlexInt    `json:"capacity"`
	Access         string     `json:"access"`
	MobileAccess   string     `json:"mobile_access"`
	URLs           URLs       `json:"urls"`
	Photo          Photo      `json:"photo"`
	Open           string     `json:"open"`
	Close          string     `json:"close"`
	PartyCapacity  FlexInt    `json:"party_capacity"`
	OtherMemo      string     `json:"other_memo"`
	ShopDetailMemo string     `json:"shop_detail_memo"`
	CouponURLs     CouponURLs `json:"coupon_urls"`

	// Area hierarchy
	LargeServiceArea CodeName `json:"large_service_area"`
	ServiceArea      CodeName `json:"service_area"`
	LargeArea        CodeName `json:"large_area"`
	MiddleArea       CodeName `json:"middle_area"`
	SmallArea        CodeName `json:"small_area"`

	// Amenity fields (string values like "あり", "なし")
	WiFi         string  `json:"wifi"`
	Wedding      string  `json:"wedding"`
	Course       string  `json:"course"`
	FreeDrink    string  `json:"free_drink"`
	FreeFood     string  `json:"free_food"`
	PrivateRoom  string  `json:"private_room"`
	Horigotatsu  string  `json:"horigotatsu"`
	Tatami       string  `json:"tatami"`
	Card         string  `json:"card"`
	NonSmoking   string  `json:"non_smoking"`
	Charter      string  `json:"charter"`
	Parking      string  `json:"parking"`
	BarrierFree  string  `json:"barrier_free"`
	Sommelier    string  `json:"sommelier"`
	OpenAir      string  `json:"open_air"`
	Show         string  `json:"show"`
	Equipment    string  `json:"equipment"`
	Karaoke      string  `json:"karaoke"`
	Band         string  `json:"band"`
	TV           string  `json:"tv"`
	English      string  `json:"english"`
	Pet          string  `json:"pet"`
	Child        string  `json:"child"`
	Lunch        string  `json:"lunch"`
	Midnight     string  `json:"midnight"`
	MidnightMeal string  `json:"midnight_meal"`
	Ktai         string  `json:"ktai"`
	KtaiCoupon   FlexInt `json:"ktai_coupon"`
	NightView    string  `json:"night_view"`
	Cocktail     string  `json:"cocktail"`
	Shochu       string  `json:"shochu"`
	Sake         string  `json:"sake"`
	Wine         string  `json:"wine"`
}

type CodeName struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Budget struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Average    string `json:"average"`
	BudgetMemo string `json:"budget_memo"`
}

type URLs struct {
	PC string `json:"pc"`
}

type Photo struct {
	PC     PhotoSizes `json:"pc"`
	Mobile PhotoSizes `json:"mobile"`
}

type PhotoSizes struct {
	L string `json:"l"`
	M string `json:"m"`
	S string `json:"s"`
}

type CouponURLs struct {
	PC string `json:"pc"`
	SP string `json:"sp"`
}

// ShopSearchResponse is the response from /shop/v1/
type ShopSearchResponse struct {
	Results ShopSearchResults `json:"results"`
}

type ShopSearchResults struct {
	APIVersion       string      `json:"api_version"`
	ResultsAvailable int         `json:"results_available"`
	ResultsReturned  string      `json:"results_returned"`
	ResultsStart     int         `json:"results_start"`
	Shops            []ShopBrief `json:"shop"`
}

type ShopBrief struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	NameKana string   `json:"name_kana"`
	Address  string   `json:"address"`
	Genre    CodeName `json:"genre"`
	URLs     URLs     `json:"urls"`
	Desc     string   `json:"desc"`
}

// Concrete response types for each master endpoint

type GenreResponse struct {
	Results GenreResults `json:"results"`
}
type GenreResults struct {
	APIVersion       string  `json:"api_version"`
	ResultsAvailable int     `json:"results_available"`
	ResultsReturned  string  `json:"results_returned"`
	ResultsStart     int     `json:"results_start"`
	Genres           []Genre `json:"genre"`
}
type Genre struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type BudgetResponse struct {
	Results BudgetResults `json:"results"`
}
type BudgetResults struct {
	APIVersion       string         `json:"api_version"`
	ResultsAvailable int            `json:"results_available"`
	ResultsReturned  string         `json:"results_returned"`
	ResultsStart     int            `json:"results_start"`
	Budgets          []BudgetMaster `json:"budget"`
}
type BudgetMaster struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LargeServiceAreaResponse struct {
	Results LargeServiceAreaResults `json:"results"`
}
type LargeServiceAreaResults struct {
	APIVersion        string             `json:"api_version"`
	ResultsAvailable  int                `json:"results_available"`
	ResultsReturned   string             `json:"results_returned"`
	ResultsStart      int                `json:"results_start"`
	LargeServiceAreas []LargeServiceArea `json:"large_service_area"`
}
type LargeServiceArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ServiceAreaResponse struct {
	Results ServiceAreaResults `json:"results"`
}
type ServiceAreaResults struct {
	APIVersion       string        `json:"api_version"`
	ResultsAvailable int           `json:"results_available"`
	ResultsReturned  string        `json:"results_returned"`
	ResultsStart     int           `json:"results_start"`
	ServiceAreas     []ServiceArea `json:"service_area"`
}
type ServiceArea struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	LargeServiceArea CodeName `json:"large_service_area"`
}

type LargeAreaResponse struct {
	Results LargeAreaResults `json:"results"`
}
type LargeAreaResults struct {
	APIVersion       string      `json:"api_version"`
	ResultsAvailable int         `json:"results_available"`
	ResultsReturned  string      `json:"results_returned"`
	ResultsStart     int         `json:"results_start"`
	LargeAreas       []LargeArea `json:"large_area"`
}
type LargeArea struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	ServiceArea      CodeName `json:"service_area"`
	LargeServiceArea CodeName `json:"large_service_area"`
}

type MiddleAreaResponse struct {
	Results MiddleAreaResults `json:"results"`
}
type MiddleAreaResults struct {
	APIVersion       string       `json:"api_version"`
	ResultsAvailable int          `json:"results_available"`
	ResultsReturned  string       `json:"results_returned"`
	ResultsStart     int          `json:"results_start"`
	MiddleAreas      []MiddleArea `json:"middle_area"`
}
type MiddleArea struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	LargeArea        CodeName `json:"large_area"`
	ServiceArea      CodeName `json:"service_area"`
	LargeServiceArea CodeName `json:"large_service_area"`
}

type SmallAreaResponse struct {
	Results SmallAreaResults `json:"results"`
}
type SmallAreaResults struct {
	APIVersion       string      `json:"api_version"`
	ResultsAvailable int         `json:"results_available"`
	ResultsReturned  string      `json:"results_returned"`
	ResultsStart     int         `json:"results_start"`
	SmallAreas       []SmallArea `json:"small_area"`
}
type SmallArea struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	MiddleArea       CodeName `json:"middle_area"`
	LargeArea        CodeName `json:"large_area"`
	ServiceArea      CodeName `json:"service_area"`
	LargeServiceArea CodeName `json:"large_service_area"`
}

type CreditCardResponse struct {
	Results CreditCardResults `json:"results"`
}
type CreditCardResults struct {
	APIVersion       string       `json:"api_version"`
	ResultsAvailable int          `json:"results_available"`
	ResultsReturned  string       `json:"results_returned"`
	ResultsStart     int          `json:"results_start"`
	CreditCards      []CreditCard `json:"credit_card"`
}
type CreditCard struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SpecialResponse struct {
	Results SpecialResults `json:"results"`
}
type SpecialResults struct {
	APIVersion       string    `json:"api_version"`
	ResultsAvailable int       `json:"results_available"`
	ResultsReturned  string    `json:"results_returned"`
	ResultsStart     int       `json:"results_start"`
	Specials         []Special `json:"special"`
}
type Special struct {
	Code            string   `json:"code"`
	Name            string   `json:"name"`
	SpecialCategory CodeName `json:"special_category"`
}

type SpecialCategoryResponse struct {
	Results SpecialCategoryResults `json:"results"`
}
type SpecialCategoryResults struct {
	APIVersion        string            `json:"api_version"`
	ResultsAvailable  int               `json:"results_available"`
	ResultsReturned   string            `json:"results_returned"`
	ResultsStart      int               `json:"results_start"`
	SpecialCategories []SpecialCategory `json:"special_category"`
}
type SpecialCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
